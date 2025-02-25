package postgres

import (
	"github.com/ShudderStorm/go-github-tracker/internal/db/postgres/ssl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_Compile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		expect string

		host       string
		port       uint16
		user, pass string
		database   string
		ssl        ssl.Mode
	}{
		{
			name:   "example::ssl_disable",
			expect: "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable",

			host: "localhost", port: 5432,
			user: "postgres", pass: "secret",
			database: "postgres", ssl: ssl.Disable,
		},
		{
			name:   "example::ssl_prefer",
			expect: "postgres://postgres:secret@localhost:5432/postgres?sslmode=prefer",

			host: "localhost", port: 5432,
			user: "postgres", pass: "secret",
			database: "postgres", ssl: ssl.Prefer,
		},
		{
			name:   "example::ssl_require",
			expect: "postgres://postgres:secret@localhost:5432/postgres?sslmode=require",

			host: "localhost", port: 5432,
			user: "postgres", pass: "secret",
			database: "postgres", ssl: ssl.Require,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			provider := New(
				test.host, test.port, test.user, test.pass, test.database,
				WithSSL(test.ssl),
			).Compile()

			assert.Equal(t, test.expect, provider.dsn)
		})
	}
}
