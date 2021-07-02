package token

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetClaims(t *testing.T) {
	t.Run("good variants", func(t *testing.T) {
		for _, tst := range [...]struct {
			token  string
			key    string
			error  error
			userId int
		}{
			{
				token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjI1MjM5NjU0fQ.rNyW5DyXeRXLrY_ex6c7wy8oHgX7c8SjfS87I20hrMg",
				key:    "secret",
				error:  nil,
				userId: 1,
			},
		} {
			claims, err := GetClaims(tst.token, tst.key)
			require.Equal(t, err, tst.error)
			require.Equal(t, claims.UserId, tst.userId)
		}
	})

	t.Run("bad variants", func(t *testing.T) {
		for _, tst := range [...]struct {
			token  string
			key    string
			error  string
			userId int
		}{
			{
				token: "",
				key:   "",
				error: "token contains an invalid number of segments",
			},
			{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjI1MjM5NjU0fQ.rNyW5DyXeRXLrY_ex6c7wy8oHgX7c8SjfS87I20hrMg",
				key:   "secret",
				error: "signature is invalid",
			},
			{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjI1MjM5NjU0fQ.rNyW5DyXeRXLrY_ex6c7wy8oHgX7c8SjfS87I20hrMg",
				key:   "not_secret",
				error: "signature is invalid",
			},
			{
				token:  "bad_token",
				key:    "secret",
				error:  "token contains an invalid number of segments",
				userId: 1,
			},
		} {
			token, _ := CreateToken(tst.userId, tst.key)
			_, err := GetClaims(token, tst.key)
			if err != nil {
				require.Contains(t, err.Error(), tst.error)
			}
		}
	})
}

func TestCreateToken(t *testing.T) {
	t.Run("good variants", func(t *testing.T) {
		for _, tst := range [...]struct {
			userId int
			key    string
			error  error
		}{
			{
				userId: 1,
				key:    "secret",
				error:  nil,
			},
			{
				userId: 12,
				key:    "",
				error:  nil,
			},
			{
				userId: 0,
				key:    "",
				error:  nil,
			},
		} {
			token, err := CreateToken(tst.userId, tst.key)
			require.Equal(t, err, tst.error)

			claims, err := GetClaims(token, tst.key)
			require.Equal(t, err, tst.error)
			require.Equal(t, claims.UserId, tst.userId)
		}
	})
}
