package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"cloud.google.com/go/spanner"
)

func main() {
	ctx := context.Background()
	spannerClient, err := spanner.NewClient(ctx, "projects/local-project/instances/test-instance/databases/local-database")
	if err != nil {
		panic(err)
	}

	err = insertUsers(ctx, spannerClient)
	if err != nil {
		panic(err)
	}

	return
}

func insertUsers(ctx context.Context, client *spanner.Client) error {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9000000000
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmts := []spanner.Statement{}
		for i := 1; i <= 100; i++ {
			r := rand.Intn(max-min) + min
			userID := fmt.Sprintf("User %d", r)
			stmts = append(stmts, spanner.Statement{
				SQL: "INSERT INTO UserScore (UserID, Score, UpdatedAt) VALUES (@userID, @score, @updatedAt)",
				Params: map[string]interface{}{
					"userID":    userID,
					"score":     r,
					"updatedAt": time.Date(2020, 3, 10, 23, 59, 59, 999, time.UTC),
				},
			})
		}
		_, err := txn.BatchUpdate(ctx, stmts)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
