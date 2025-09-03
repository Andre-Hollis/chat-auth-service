package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	Username     string     `json:"username"`
	PasswordHash string     `json:"passwordHash"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastLogin    *time.Time `json:"lastLogin"`
}

type UserCreate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) SetId(_id string) {

}
