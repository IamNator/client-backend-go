package user

import (
	"context"
	"database/sql"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ivorscott/devpie-client-backend-go/internal/platform/database"
	"github.com/pkg/errors"
)

// The User package shouldn't know anything about http
// While it may identify common know errors, how to respond is left to the handlers
var (
	ErrNotFound       = errors.New("user not found")
	ErrInvalidID      = errors.New("id provided was not a valid UUID")
	ErrInvalidAuth0ID = errors.New("id provided was not a valid Auth0ID")
)

// Retrieve finds the User identified by a given Auth0ID.
func RetrieveMe(ctx context.Context, repo *database.Repository, sub string) (*User, error) {
	var u User

	s := strings.Split(sub, "|")
	auth0Id := s[1]

	if auth0Id == "" {
		return nil, ErrInvalidAuth0ID
	}

	stmt := repo.SQ.Select(
		"id",
		"auth0id",
		"email",
		"firstname",
		"lastname",
		"emailverified",
		"locale",
		"picture",
		"created",
	).From(
		"users",
	).Where(sq.Eq{"auth0id": "?"})

	q, args, err := stmt.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "building query: %v", args)
	}

	if err := repo.DB.GetContext(ctx, &u, q, auth0Id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &u, nil
}

// Create adds a new User
func Create(ctx context.Context, repo *database.Repository, nu NewUser, now time.Time) (*User, error) {

	s := strings.Split(nu.Auth0ID, "|")
	auth0Id := s[1]

	u := User{
		ID:            uuid.New().String(),
		Auth0ID:       auth0Id,
		Email:         nu.Email,
		EmailVerified: nu.EmailVerified,
		FirstName:     nu.FirstName,
		LastName:      nu.LastName,
		Picture:       nu.Picture,
		Locale:        nu.Locale,
		Created:       now.UTC(),
	}

	stmt := repo.SQ.Insert(
		"users",
	).SetMap(map[string]interface{}{
		"id":            u.ID,
		"auth0id":       u.Auth0ID,
		"email":         u.Email,
		"emailverified": u.EmailVerified,
		"firstname":     u.FirstName,
		"lastname":      u.LastName,
		"picture":       u.Picture,
		"locale":        u.Locale,
		"created":       u.Created,
	})

	if _, err := stmt.ExecContext(ctx); err != nil {
		return nil, errors.Wrapf(err, "inserting user: %v", nu)
	}

	return &u, nil
}
