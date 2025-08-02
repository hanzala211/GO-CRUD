package utils

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/hanzala211/CRUD/internal/api/models"
)

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Post)(nil),
		(*models.Comment)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}

	// Set up UUID defaults for both tables
	_, err := db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
		
		-- Set default for users table
		ALTER TABLE users 
		ALTER COLUMN id SET DEFAULT gen_random_uuid();
		
		-- Set default for posts table  
		ALTER TABLE posts 
		ALTER COLUMN id SET DEFAULT gen_random_uuid();

		-- Set default for comments table  
		ALTER TABLE comments 
		ALTER COLUMN id SET DEFAULT gen_random_uuid();
		
		-- Set timestamp defaults for posts table
		ALTER TABLE posts 
		ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP,
		ALTER COLUMN updated_at SET DEFAULT CURRENT_TIMESTAMP;
		
		-- Set timestamp defaults for comments table
		ALTER TABLE comments 
		ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP,
		ALTER COLUMN updated_at SET DEFAULT CURRENT_TIMESTAMP;
		
		-- Create trigger function for updating updated_at
		CREATE OR REPLACE FUNCTION update_updated_at_column()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = CURRENT_TIMESTAMP;
			RETURN NEW;
		END;
		$$ language 'plpgsql';
			
		-- Create triggers for posts table
		DROP TRIGGER IF EXISTS update_posts_updated_at ON posts;
		CREATE TRIGGER update_posts_updated_at
			BEFORE UPDATE ON posts
			FOR EACH ROW
			EXECUTE FUNCTION update_updated_at_column();

		-- Create triggers for comments table
		DROP TRIGGER IF EXISTS update_comments_updated_at ON comments;
		CREATE TRIGGER update_comments_updated_at
			BEFORE UPDATE ON comments
			FOR EACH ROW
			EXECUTE FUNCTION update_updated_at_column();
	`)

	return err
}
