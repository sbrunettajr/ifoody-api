package test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/sbrunettajr/ifoody-api/infra/db"
)

var d *sql.DB // Rename

func TestMain(m *testing.M) {
	var err error
	d, err = db.NewDB()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
