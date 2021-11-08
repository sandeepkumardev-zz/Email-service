package models

type User struct {
	ID       BinaryUUID `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}
