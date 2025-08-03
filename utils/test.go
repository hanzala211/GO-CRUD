package utils

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/repo"
)

func Test(n int, db *pg.DB) {
	var wg sync.WaitGroup
	repo := repo.NewCommentRepo(db)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(index int) {
			comment := &models.Comment{
				UserId:  "465706ef-61f7-4bc4-9fd5-52a10cec25b0",
				PostId:  "e096c433-700f-457a-b8b9-f1ae76d1b63a",
				Content: "test123$",
			}
			err := repo.TestFunc(context.Background(), comment)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Successfully created comment %d\n", index)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
