package repositories

import (
	"database/sql"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(user *models.User) (*models.User, error) {

	createQuery := `
		insert into users (username, email, password, created_at)
		values ($1, $2, $3, $4) returning id;
	`

	if err := u.db.QueryRow(
		createQuery,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
	).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetEmailOrUsername(req string) (*models.User, error) {

	var user models.User

	getEmailOrUsernameQuery := `
		select id, username, email, created_at, updated_at from users 
		where email = $1 or username = $1;
	`

	if err := u.db.QueryRow(getEmailOrUsernameQuery, req).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
		return nil, err
	}

	return &user, nil
}
