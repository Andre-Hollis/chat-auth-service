package impl

import (
	"context"
	"encoding/json"

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

func (r *UserRepositoryDB) ReadUserByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	j, err := r.db.Get(ctx, email).Result()

	u := userdomain.User{}

	if err != nil {
		return &userdomain.User{}, err
	}

	err = json.Unmarshal([]byte(j), &u)

	if err != nil {
		return &userdomain.User{}, err
	}

	return &u, nil
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
