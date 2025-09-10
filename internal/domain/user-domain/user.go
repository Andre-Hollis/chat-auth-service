package userdomain

import (
	"time"
)

type User struct {
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"passwordHash"`
	CreatedAt    time.Time `json:"createdAt"`
	LastLogin    time.Time `json:"lastLogin"`
}

func (u *User) SetId(_id string) {

}
