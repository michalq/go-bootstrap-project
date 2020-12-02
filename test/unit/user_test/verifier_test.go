package user_test

import (
	"testing"

	"github.com/michalq/go-bootstrap-project/internal/modules/user"
	"github.com/stretchr/testify/assert"
)

var usersRepoMock = struct{ user.Users }{}

func TestVerifier(t *testing.T) {
	testTable := []struct {
		name        string
		user        *user.User
		expectedErr error
	}{
		{
			name:        "return nil if every condition passes",
			user:        &user.User{Name: "Some dummy name", Verified: false},
			expectedErr: nil,
		},
		{
			name:        "return error on empty name",
			user:        &user.User{Name: "", Verified: false},
			expectedErr: user.ErrUserWithoutName,
		},
		{
			name:        "return error if user already verified",
			user:        &user.User{Name: "Some dummy name", Verified: true},
			expectedErr: user.ErrUserAlreadyVerified,
		},
	}

	verifier := user.NewVerifier(usersRepoMock)
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedErr, verifier.Verify(tt.user))
		})
	}
}
