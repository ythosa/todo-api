package repository

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"github.com/Inexpediency/todo-rest-api/pkg/models"
)

type AuthPostgres struct {
	db    *sqlx.DB
	cache *redis.Client
}

func NewAuthPostgres(db *sqlx.DB, cache *redis.Client) *AuthPostgres {
	return &AuthPostgres{
		db:    db,
		cache: cache,
	}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id",
		usersTable,
	)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)

	return user, err
}

func (r *AuthPostgres) SaveRefreshToken(userId int, token string) error {
	return r.cache.Set(redisCtx, strconv.Itoa(userId), token, redisTTL).Err()
}

func (r *AuthPostgres) GetRefreshToken(userId int) (string, error) {
	return r.cache.Get(redisCtx, strconv.Itoa(userId)).Result()
}
