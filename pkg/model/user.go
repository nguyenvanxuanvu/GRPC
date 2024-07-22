package model


type User struct {
	ID                  int64      `json:"id"`
	DisplayName         *string    `json:"display_name"`
	Username            *string    `json:"username"`
	Email               *string    `json:"email"`
}