package mysql_test

import (
	"testing"

	"github.com/ShudderStorm/go-github-tracker/internal/db/mysql"
	"github.com/stretchr/testify/assert"
)

func TestConfigurator(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		casename string

		user        string
		passwd      string
		addres      string
		dbname      string
		params      map[string]string
		expectedDSN string
	}{
		{
			casename: "EmptyPassword",

			user:        "testuser",
			passwd:      "",
			addres:      "localhost:3306",
			dbname:      "testdb",
			params:      map[string]string{},
			expectedDSN: "testuser@tcp(localhost:3306)/testdb",
		},
		{
			casename: "WithPassword",

			user:        "testuser",
			passwd:      "testpass",
			addres:      "localhost:3306",
			dbname:      "testdb",
			params:      map[string]string{},
			expectedDSN: "testuser:testpass@tcp(localhost:3306)/testdb",
		},
		{
			casename: "WithParams",

			user:        "testuser",
			passwd:      "testpass",
			addres:      "localhost:3306",
			dbname:      "testdb",
			params:      map[string]string{"charset": "utf8", "parseTime": "true"},
			expectedDSN: "testuser:testpass@tcp(localhost:3306)/testdb?charset=utf8&parseTime=true",
		},
	}

	for _, test := range testcases {
		test := test
		t.Run(test.casename, func(t *testing.T) {
			t.Parallel()

			resultDSN := mysql.New().
				SetUser(test.user).
				SetPassword(test.passwd).
				UseTCP().
				SetAddress(test.addres).
				SetDBName(test.dbname).
				AddParams(test.params).
				FormatDSN()

			assert.Equal(t, test.expectedDSN, resultDSN)
		})
	}
}
