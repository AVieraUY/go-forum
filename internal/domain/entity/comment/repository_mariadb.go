package comment

import (
	"database/sql"

	"github.com/AVieraUY/go-forum/internal/domain"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
	"github.com/AVieraUY/go-forum/internal/domain/entity/post"
	"github.com/AVieraUY/go-forum/internal/domain/entity/user"
)

// MariaDBRepo mariadb repository
type MariaDBRepo struct {
	db *sql.DB
}

// NewMariaDBRepo creates a new mariadbrepo
func NewMariaDBRepo(db *sql.DB) *MariaDBRepo {
	return &MariaDBRepo{db: db}
}

// Get a comment
func (r *MariaDBRepo) Get(id entity.ID) (*Comment, error) {
	rows, err := r.db.Query("select * from comment where id = ?", id)
	if err != nil {
		return nil, err
	}

	var c Comment
	var userID entity.ID
	var postID entity.ID
	for rows.Next() {
		err = rows.Scan(&c.ID, &userID, &postID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
		c.User = user.User{ID: userID}
		c.Post = post.Post{ID: postID}
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Search comments
func (r *MariaDBRepo) Search(query string) ([]*Comment, error) {
	rows, err := r.db.Query("select * from comment where body like ?", query)
	if err != nil {
		return nil, err
	}

	comments, err := fillComments(rows)
	if err != nil {
		return nil, err
	}

	if len(comments) <= 0 {
		return nil, domain.ErrNotFound
	}

	return comments, nil
}

// List comments
func (r *MariaDBRepo) List() ([]*Comment, error) {
	rows, err := r.db.Query("select * from comment")
	if err != nil {
		return nil, err
	}

	comments, err := fillComments(rows)
	if err != nil {
		return nil, err
	}

	if len(comments) <= 0 {
		return nil, domain.ErrNotFound
	}

	return comments, nil
}

// Create a new comment
func (r *MariaDBRepo) Create(e *Comment) (entity.ID, error) {
	_, err := r.db.Exec("insert into comment (id, user_id, post_id, body, created_at, updated_at) values (?,?,?,?,?,?)", e.ID, e.User.ID, e.Post.ID, e.Body, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// Update an existing comment
func (r *MariaDBRepo) Update(e *Comment) error {
	_, err := r.db.Exec("update comment set body = ?, updated_at = ? where id = ?", e.Body, e.UpdatedAt, e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete an existing comment
func (r *MariaDBRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from comment where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func fillComments(rows *sql.Rows) ([]*Comment, error) {
	var comments []*Comment
	var err error
	for rows.Next() {
		var c Comment
		var userID entity.ID
		var postID entity.ID
		err = rows.Scan(&c.ID, &userID, &postID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
		c.User = user.User{ID: userID}
		c.Post = post.Post{ID: postID}
		comments = append(comments, &c)
	}

	if err != nil {
		return nil, err
	}

	return comments, nil
}
