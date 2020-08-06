package user

import (
	"database/sql"

	"github.com/AVieraUY/go-forum/internal/domain"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// MariaDBRepo mariadb repo
type MariaDBRepo struct {
	db *sql.DB
}

// NewMariaDBRepo creates a new mariadb repo
func NewMariaDBRepo(db *sql.DB) *MariaDBRepo {
	return &MariaDBRepo{db: db}
}

// Get an user
func (r *MariaDBRepo) Get(id entity.ID) (*User, error) {
	stmt, err := r.db.Prepare(`select * from user where id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	var u User
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.AvatarPath, &u.CreatedAt, &u.UpdatedAt)
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Search users
func (r *MariaDBRepo) Search(query string) ([]*User, error) {
	stmt, err := r.db.Prepare(`select * from user where username like ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(query)
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.AvatarPath, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, &u)
	}

	if len(users) <= 0 {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

// List users
func (r *MariaDBRepo) List() ([]*User, error) {
	stmt, err := r.db.Prepare("select * from user")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.AvatarPath, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, &u)
	}

	if len(users) <= 0 {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Create an user
func (r *MariaDBRepo) Create(e *User) (entity.ID, error) {
	stmt, err := r.db.Prepare(`insert into user (id, username, email, password, avatar_path, created_at, updated_at) values(?,?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}

	_, err = stmt.Exec(e.ID, e.Username, e.Email, e.Password, e.AvatarPath, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// Update an user
func (r *MariaDBRepo) Update(e *User) error {
	stmt, err := r.db.Prepare(`update user set username = ?, email = ?, password = ?, avatar_path = ?, updated_at = ? where id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Username, e.Email, e.Password, e.AvatarPath, e.UpdatedAt, e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete an user
func (r *MariaDBRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec(`delete from user where id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}
