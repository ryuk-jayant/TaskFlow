package user

import (
	"database/sql"
	"example/web-service-gin/types"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserBYEmail(email string) (*types.User, error) {
	query := `
		SELECT CONVERT(VARCHAR(36), id) AS id, name, email, password, created_at
		FROM Users
		WHERE email = @p1
	`

	row := s.db.QueryRow(query, email)

	u := new(types.User)
	err := row.Scan(
		&u.Id,
		// &u.USERID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) GetUserBYID(id string) (*types.User, error) {
	query := `
		SELECT  CONVERT(VARCHAR(36), id) AS id , name, email, created_at
		FROM Users
		WHERE id = @p1
	`

	row := s.db.QueryRow(query, id)

	u := new(types.User)
	err := row.Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}
func (s *Store) CreateUser(u *types.User) error {
	query := `
		INSERT INTO Users (id, name, email, password, created_at)
		VALUES (@p1, @p2, @p3, @p4, @p5)
	`

	_, err := s.db.Exec(
		query,
		u.Id,
		u.Name,
		u.Email,
		u.Password,
		u.Created_at,
	)

	return err
}
func ScanUserRow(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Password,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
