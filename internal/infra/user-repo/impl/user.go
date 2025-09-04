package user

import (
	"context"

	userdomain "github.com/Andre-Hollis/chat-auth-service/internal/domain/user-domain"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserRepositoryDB struct {
	db *redis.Client
}

func NewUserRedisRepo() *UserRepositoryDB {
	return &UserRepositoryDB{
		db: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (r *UserRepositoryDB) Save(ctx context.Context, user *userdomain.User) (string, error) {
	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	err := r.db.Set(ctx, user.ID, user, 0).Err()

	if err != nil {
		return "", err
	}

	return user.ID, err
}
