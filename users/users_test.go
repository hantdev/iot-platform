package users_test

import (
	"fmt"
	"testing"

	"github.com/hantdev/athena/errors"
	"github.com/hantdev/athena/users"
	"github.com/stretchr/testify/assert"
)

const (
	email    = "user@example.com"
	password = "password"
	metadata = `{"role":"manager"}`
)

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		user users.User
		err  error
	}{
		"validate user with valid data": {
			user: users.User{
				Email:    email,
				Password: password,
			},
			err: nil,
		},
		"validate user with empty email": {
			user: users.User{
				Email:    "",
				Password: password,
			},
			err: users.ErrMalformedEntity,
		},
		"validate user with empty password": {
			user: users.User{
				Email:    email,
				Password: "",
			},
			err: users.ErrMalformedEntity,
		},
		"validate user with invalid email": {
			user: users.User{
				Email:    "userexample.com",
				Password: password,
			},
			err: users.ErrMalformedEntity,
		},
	}

	for desc, tc := range cases {
		err := tc.user.Validate()
		assert.True(t, errors.Contains(err, tc.err), fmt.Sprintf("%s: expected %s got %s\n", desc, tc.err, err))
	}
}
