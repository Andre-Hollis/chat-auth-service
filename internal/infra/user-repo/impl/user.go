package user

import (
	"context"
	"database/sql"

	"github.com/Andre-Hollis/chat-auth-service/internal/domain/user"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func (r *UserRepositoryDB) Save(ctx context.Context, user *user.User) {
	// insert into db
}
