package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	spannerClient, err := spanner.NewClient(ctx, "projects/scarlet-spanner/instances/scarlet-instance/databases/scarlet-database")
	if err != nil {
		panic(err)
	}

	err = count(ctx, spannerClient)
	if err != nil {
		panic(err)
	}

	return
}

func count(ctx context.Context, client *spanner.Client) error {
	stmt := spanner.Statement{SQL: "SELECT UserID, Score FROM UserScore ORDER BY Score DESC LIMIT 10"}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		var playerID string
		var score int64
		if err := row.Columns(&playerID, &score); err != nil {
			return err
		}
		fmt.Printf("PlayerId: %s Score: %d  \n", playerID, score)
	}
}
