package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"app/db"
	"app/templates"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	ctx := context.Background()

	client, err := db.NewClient(ctx)
	if err != nil {
		log.Fatalf("db client: %v", err)
	}
	if err := db.Bootstrap(ctx, client); err != nil {
		log.Fatalf("db bootstrap: %v", err)
	}
	log.Println("Database ready")

	// remove
	if err := testDB(ctx, client); err != nil {
		log.Fatalf("db test: %v", err)
	}
	// remove

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Index().Render(r.Context(), w)
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from the updated server!")
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO: remove — temporary smoke test for DynamoDB read/write
func testDB(ctx context.Context, client *dynamodb.Client) error {
	user := db.User{
		ID:        "test-user-1",
		Email:     "test@example.com",
		CreatedAt: time.Now().Unix(),
	}

	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.TableUsers),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("put: %w", err)
	}
	log.Printf("wrote user: %+v", user)

	out, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.TableUsers),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: user.ID},
		},
	})
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}

	var got db.User
	if err := attributevalue.UnmarshalMap(out.Item, &got); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	log.Printf("read back user: %+v", got)
	return nil
}
