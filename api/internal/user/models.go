package user

import (
	"time"
)

// Users represent user data from the database
type User struct {
	ID            string    `db:"id" json:"id" `
	Auth0ID       string    `db:"auth0id" json:"auth0Id" `
	Email         string    `db:"email" json:"email"`
	EmailVerified bool      `db:"emailverified" json:"emailVerified"`
	FirstName     *string   `db:"firstname" json:"firstName"`
	LastName      *string   `db:"lastname" json:"lastName"`
	Picture       *string   `db:"picture" json:"picture"`
	Locale        *string   `db:"locale" json:"locale"`
	Created       time.Time `db:"created" json:"created"`
}

type NewUser struct {
	Auth0ID       string  `json:"auth0Id" `
	Email         string  `json:"email"`
	EmailVerified bool    `json:"emailVerified"`
	FirstName     *string `json:"firstName"`
	LastName      *string `json:"lastName"`
	Picture       *string `json:"picture"`
	Locale        *string `json:"locale"`
}

type UpdateUser struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Picture   *string `json:"picture"`
	Locale    *string `json:"locale"`
}
