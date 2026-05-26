package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TableUsers = "Users"
	TableTiles = "Tiles"
)

func NewClient(ctx context.Context) (*dynamodb.Client, error) {
	endpoint := os.Getenv("DYNAMO_ENDPOINT")

	var cfg aws.Config
	var err error

	if endpoint != "" {
		// Local dev: point at DynamoDB Local with dummy credentials
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithRegion("us-east-1"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("local", "local", "")),
		)
		if err != nil {
			return nil, fmt.Errorf("load config: %w", err)
		}
		return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		}), nil
	}

	// Production: use ambient credentials (IAM role, env vars, etc.)
	cfg, err = config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}
	return dynamodb.NewFromConfig(cfg), nil
}

// Bootstrap creates tables if they do not already exist.
func Bootstrap(ctx context.Context, client *dynamodb.Client) error {
	if err := ensureTable(ctx, client, &dynamodb.CreateTableInput{
		TableName: aws.String(TableUsers),
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("id"), KeyType: types.KeyTypeHash},
		},
		BillingMode: types.BillingModePayPerRequest,
	}); err != nil {
		return err
	}

	return ensureTable(ctx, client, &dynamodb.CreateTableInput{
		TableName: aws.String(TableTiles),
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("user_id"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("user_id"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("id"), KeyType: types.KeyTypeRange},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
}

func ensureTable(ctx context.Context, client *dynamodb.Client, input *dynamodb.CreateTableInput) error {
	_, err := client.CreateTable(ctx, input)
	if err == nil {
		return nil
	}
	// ResourceInUseException means the table already exists — that's fine
	var inUse *types.ResourceInUseException
	if errors.As(err, &inUse) {
		return nil
	}
	return fmt.Errorf("create table %s: %w", *input.TableName, err)
}
