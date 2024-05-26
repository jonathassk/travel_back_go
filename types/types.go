package types

type RegistrationType struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Currency  string `json:"currency"`
	Language  string `json:"language"`
}

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Currency  string `json:"currency"`
	Language  string `json:"language"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(u *RegistrationType) (*User, error)
}
