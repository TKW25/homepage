package db

// User represents an authenticated user stored in the Users table.
type User struct {
	ID        string `dynamodbav:"id"`
	Email     string `dynamodbav:"email"`
	CreatedAt int64  `dynamodbav:"created_at"` // Unix timestamp (seconds)
}

// Tile represents a single link tile in the Tiles table.
// UserID is the partition key; ID is the sort key.
type Tile struct {
	UserID   string `dynamodbav:"user_id"`
	ID       string `dynamodbav:"id"`
	Label    string `dynamodbav:"label"`
	URL      string `dynamodbav:"url"`
	IconURL  string `dynamodbav:"icon_url,omitempty"`
	Position int    `dynamodbav:"position"`
}
