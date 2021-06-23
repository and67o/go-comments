package configuration

import (
	"github.com/pkg/errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type test struct {
	path string
	err  error
}

func TestConfiguration(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		for _, tst := range [...]test{
			{
				path: "",
				err:  errors.New("path empty"),
			},
		} {
			_, err := New(tst.path)
			if err != nil {
				require.Equal(t, tst.err.Error(), err.Error())
			}
		}
	})

	t.Run("wrong path", func(t *testing.T) {
		for _, tst := range [...]test{
			{
				path: "./wrong_path/config_test.toml",
				err:  errors.New("open file: open ./wrong_path/config_test.toml: no such file or directory"),
			},
			{
				path: "wrong_path",
				err:  errors.New("open file: open wrong_path: no such file or directory"),
			},
		} {
			_, err := New(tst.path)
			if err != nil {
				require.Equal(t, tst.err.Error(), err.Error())
			}
		}
	})

	t.Run("pass result", func(t *testing.T) {
		for _, tst := range [...]test{
			{
				path: "./testdata/config_test.yaml",
			},
		} {
			c, err := New(tst.path)
			require.Nil(t, err)

			require.Equal(t, c.Server.Host, "lh")
			require.Equal(t, c.Server.Port, "6767")
			require.Equal(t, c.Server.Timeout.Read, time.Duration(15))
		}
	})
}
