package like

import (
	"database/sql"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// MariaDBRepo mariadb repository
type MariaDBRepo struct {
	db *sql.DB
}

// NewMariaDBRepo creates a new mariadb repository
func NewMariaDBRepo(db *sql.DB) *MariaDBRepo {
	return &MariaDBRepo{db: db}
}

// Get a like
func (s *MariaDBRepo) Get(id entity.ID) (*Like, error) {
	rows, err := s.db.Query("select * from like where id = ?", id)
	if err != nil {
		return nil, err
	}

	var l Like
	for rows.Next() {
		err = rows.Scan(&l.ID, &l.UserID, &l.PostID, &l.CreatedAt)
	}

	if err != nil {
		return nil, err
	}

	return &l, nil
}

// GetByPostID likes from a post
func (s *MariaDBRepo) GetByPostID(id entity.ID) ([]*Like, error) {
	rows, err := s.db.Query("select * from like where post_id = ?", id)
	if err != nil {
		return nil, err
	}

	var likes []*Like
	for rows.Next() {
		var l Like
		err = rows.Scan(&l.ID, &l.UserID, &l.PostID, &l.CreatedAt)
		likes = append(likes, &l)
	}

	if err != nil {
		return nil, err
	}

	return likes, nil
}

// Create a new like
func (s *MariaDBRepo) Create(e *Like) (entity.ID, error) {
	_, err := s.db.Exec("insert into like (id, user_id, post_id, created_at, updated_at) values (?,?,?,?,?)", e.ID, e.UserID, e.PostID, e.CreatedAt)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// Delete like
func (s *MariaDBRepo) Delete(id entity.ID) error {
	_, err := s.db.Exec("delete from like where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// DeletePostLikes deletes all likes from a post
func (s *MariaDBRepo) DeletePostLikes(id entity.ID) error {
	_, err := s.db.Exec("delete from like where post_id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserLikes deletes all likes from a user
func (s *MariaDBRepo) DeleteUserLikes(id entity.ID) error {
	_, err := s.db.Exec("delete from like where user_id = ?", id)
	if err != nil {
		return err
	}

	return err
}
