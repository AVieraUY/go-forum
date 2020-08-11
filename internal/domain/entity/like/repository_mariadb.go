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
		err = rows.Scan(&l.ID, &l.UserID, &l.PostID, &l.CreatedAt, &l.UpdatedAt)
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
		err = rows.Scan(&l.ID, &l.UserID, &l.PostID, &l.CreatedAt, &l.UpdatedAt)
		likes = append(likes, &l)
	}

	if err != nil {
		return nil, err
	}

	return likes, nil
}

// Create a new like
func (s *MariaDBRepo) Create(e *Like) (entity.ID, error) {
	return e.ID, nil
}

// Delete like
func (s *MariaDBRepo) Delete(id entity.ID) error {
	return nil
}

// DeletePostLikes deletes all likes from a post
func (s *MariaDBRepo) DeletePostLikes(id entity.ID) error {
	return nil
}

// DeleteUserLikes deletes all likes from a user
func (s *MariaDBRepo) DeleteUserLikes(id entity.ID) error {
	return nil
}
