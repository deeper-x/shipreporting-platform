package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSelectUserID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	retRows := sqlmock.NewRows([]string{"user_id"}).AddRow("7")

	mock.ExpectQuery("SELECT user_id FROM authtoken_token WHERE key = ?").
		WithArgs("ac772f9ebdb24b3b3399c30d82a63f56c808d070").
		WillReturnRows(retRows)

	r := Repository{Conn: db}

	rows := r.SelectUserID("ac772f9ebdb24b3b3399c30d82a63f56c808d070")

	assert.Equal(t, rows, "7")
}

func TestSelectUserPortinformer(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	retRows := sqlmock.NewRows([]string{"fk_portinformer"}).AddRow("28")

	mock.ExpectQuery("SELECT fk_portinformer FROM portinformer_managers WHERE fk_user = ?").
		WithArgs("7").
		WillReturnRows(retRows)

	r := Repository{Conn: db}

	rows := r.SelectUserPortinformer("7")

	assert.Equal(t, rows, "28")

}
