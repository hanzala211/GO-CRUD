package models

import "time"

type Comment struct {
	ID        string    `json:"id,omitempty" pg:"id,pk"`
	Content   string    `json:"content" pg:"content"`
	CreatedAt time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	UserId    string    `json:"user_id" pg:"user_id"`
	User      *User     `json:"user,omitempty" pg:"rel:has-one,fk:user_id"`
	PostId    string    `json:"post_id" pg:"post_id"`
}
