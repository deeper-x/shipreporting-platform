package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDemoSelect(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	retRows := sqlmock.NewRows([]string{"user_id"}).AddRow("7")

	mock.ExpectQuery("SELECT user_id FROM authtoken_token WHERE key = ?").
		WithArgs("624df5a5abd1b86d7d51c238604489cd9c766a92").
		WillReturnRows(retRows)

	r := Repository{conn: db}

	rows := r.SelectUserID("624df5a5abd1b86d7d51c238604489cd9c766a92")

	assert.Equal(t, rows, []string{"7"})
}
