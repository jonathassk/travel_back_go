package user

import (
	"database/sql"
	"fmt"
	"github.com/jonathassk/travel_back_go/service/auth"
	"github.com/jonathassk/travel_back_go/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(u *types.RegistrationType) (*types.User, error) {
	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	user := &types.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  hashedPassword,
		City:      u.City,
		Country:   u.Country,
		Currency:  u.Currency,
		Language:  u.Language,
	}
	if err != nil {
		return nil, err
	}
	_, err = s.db.Exec("INSERT INTO users (first_name, last_name, email, password, city, country, currency, language) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		user.FirstName, user.LastName, user.Email, hashedPassword, user.City, user.Country, user.Currency, user.Language)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	var user types.User
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.City,
		&user.Country,
		&user.Currency,
		&user.Language,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
