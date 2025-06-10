package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/models"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (p *PostRepo) CreatePost(post *models.Post) (*models.Post, error) {
	query := `
        insert into posts (user_id, content, image_url, created_at, updated_at)
        values ($1, $2, $3, $4, $5)
        returning id, created_at, updated_at
    `

	now := time.Now()

	err := p.db.QueryRow(
		query,
		post.UserID,
		post.Content,
		post.ImageURL,
		now,
		now,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostRepo) DeletePost(postId uint) error {
	query := `delete from posts where id = $1`

	result, err := p.db.Exec(query, postId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}

func (p *PostRepo) FindPostsByUserId(userId uint) ([]*models.Post, error) {
	query := `
        select 
            p.id, p.user_id, p.content, p.image_url, p.created_at, p.updated_at,
            u.id, u.username, u.email
        from 
            posts p
        join 
            users u ON p.user_id = u.id
        where 
            p.user_id = $1
        order by
            p.created_at DESC
    `

	rows, err := p.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {

		var post models.Post
		var user models.User

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Content,
			&post.ImageURL,
			&post.CreatedAt,
			&post.UpdatedAt,
			&user.ID,
			&user.Username,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}

		post.User = &user
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostRepo) FindPostByPostId(postId uint) (*models.Post, error) {
	query := `
		select 
			 p.id, p.user_id, p.content, p.image_url, p.created_at, p.updated_at,
            u.id, u.username, u.email
		from 
			posts as p
		join 
			users as u on p.user_id = u.id
		where
			p.id = $1
		limit 1
	`

	var user models.User
	var post models.Post

	err := p.db.QueryRow(query, postId).Scan(
		&post.ID,
		&post.UserID,
		&post.Content,
		&post.ImageURL,
		&post.CreatedAt,
		&post.UpdatedAt,
		&user.ID,
		&user.Username,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	post.User = &user
	return &post, nil
}

func (p *PostRepo) UpdatePost(post *models.Post) (*models.Post, error) {
	query := `
        update posts
        set 
            content = $1,
            image_url = $2,
            updated_at = $3
        where 
            id = $4 AND 
            user_id = $5
        returning updated_at
    `

	updatedAt := time.Now()
	err := p.db.QueryRow(
		query,
		post.Content,
		post.ImageURL,
		updatedAt,
		post.ID,
		post.UserID,
	).Scan(&post.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found or user not authorized")
		}
		return nil, err
	}

	post.UpdatedAt = updatedAt
	return post, nil
}
