package post

import (
	"database/sql"

	"github.com/AVieraUY/go-forum/internal/domain/entity/user"

	"github.com/AVieraUY/go-forum/internal/domain"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// MariaDBRepo mariadb repository
type MariaDBRepo struct {
	db *sql.DB
}

// NewMariaDBRepo creates a new mariadbrepo
func NewMariaDBRepo(db *sql.DB) *MariaDBRepo {
	return &MariaDBRepo{db: db}
}

// Get an post
func (r *MariaDBRepo) Get(id entity.ID) (*Post, error) {
	rows, err := r.db.Query("select * from post where id = ?", id)
	if err != nil {
		return nil, err
	}

	var p Post
	var authorID entity.ID
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &authorID, &p.CreatedAt, &p.UpdatedAt)
		p.Author = user.User{ID: authorID}
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Search posts
func (r *MariaDBRepo) Search(query string) ([]*Post, error) {
	rows, err := r.db.Query("select * from post where title like ?", query)
	if err != nil {
		return nil, err
	}

	posts, err := fillPosts(rows)

	if err != nil {
		return nil, err
	}

	if len(posts) <= 0 {
		return nil, domain.ErrNotFound
	}

	return posts, nil
}

// List posts
func (r *MariaDBRepo) List() ([]*Post, error) {
	rows, err := r.db.Query("select * from post")
	if err != nil {
		return nil, err
	}

	posts, err := fillPosts(rows)

	if err != nil {
		return nil, err
	}

	if len(posts) <= 0 {
		return nil, domain.ErrNotFound
	}

	return posts, nil
}

// Create an post
func (r *MariaDBRepo) Create(e *Post) (entity.ID, error) {
	_, err := r.db.Exec("insert into post (id, title, content, author_id, created_at, updated_at) values (?,?,?,?,?,?)", e.ID, e.Title, e.Content, e.Author.ID, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// Update an post
func (r *MariaDBRepo) Update(e *Post) error {
	_, err := r.db.Exec("update post set title = ?, content = ?, author_id = ?, updated_at = ? where id = ?", e.Title, e.Content, e.Author.ID, e.UpdatedAt, e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete an post
func (r *MariaDBRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from post where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func fillPosts(rows *sql.Rows) ([]*Post, error) {
	var err error
	var posts []*Post
	for rows.Next() {
		var p Post
		var authorID entity.ID
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &authorID, &p.CreatedAt, &p.UpdatedAt)
		p.Author = user.User{ID: authorID}
		posts = append(posts, &p)
	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}
