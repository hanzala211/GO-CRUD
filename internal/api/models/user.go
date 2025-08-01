package models

type User struct {
	ID    string `json:"id,omitempty" pg:"id,pk"`
	Name  string `json:"name" pg:"name"`
	Email string `json:"email" pg:"email"`
}
