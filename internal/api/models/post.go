package models

import "time"

type Post struct {
	ID        string     `json:"id,omitempty" pg:"id,pk"`
	Title     string     `json:"title" pg:"title"`
	Content   string     `json:"content" pg:"content"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at"`
	UserID    string     `json:"user_id" pg:"user_id"`
	User      *User      `json:"user,omitempty" pg:"rel:has-one,fk:user_id"`
	Comment   []*Comment `json:"comments" pg:"rel:has-many,fk:post_id"`
}
