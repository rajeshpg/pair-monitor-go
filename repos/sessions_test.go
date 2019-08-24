package repos

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	. "github.com/rajeshpg/pair-monitor-go/models"
	"regexp"
	"testing"
)

func setupMockDb() (*gorm.DB, sqlmock.Sqlmock) {
	db, mockSql, err := sqlmock.New()
	mockDb, _ := gorm.Open("sqlite3", db)
	if err != nil {
		fmt.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return mockDb, mockSql
}

func TestDevPairDao_SaveSession(t *testing.T) {
	mockDb, mockSql := setupMockDb()
	defer mockDb.Close()
	t.Run("save given session", func(t *testing.T) {
		dao := &DevPairDao{Db: mockDb}
		devPair := &DevPair{Dev1: "superman", Dev2: "batman"}

		mockSql.ExpectBegin()
		mockSql.ExpectExec(regexp.QuoteMeta(`INSERT  INTO "dev_pairs" ("dev1","dev2") VALUES (?,?)`)).
			WithArgs("superman", "batman").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mockSql.ExpectCommit()

		got, _ := dao.SaveSession(devPair)
		want := uint(1)

		if got != want {
			t.Errorf("got %v, want %v ", got , want)
		}
	})
}

func TestDevPairDao_AllSessions(t *testing.T) {
	mockDb, mockSql := setupMockDb()
	defer mockDb.Close()
	t.Run("retreive all sessions", func(t *testing.T) {
		dao := &DevPairDao{Db: mockDb}

		rows := sqlmock.NewRows([]string{"id", "dev1", "dev2"}).
			AddRow(1, "superman", "batman")
		mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "dev_pairs"`)).
			WillReturnRows(rows)
		got, _ := dao.AllSessions()
		want := DevPair{ID: 1, Dev1: "superman", Dev2: "batman"}
		if got[0] != want {
			t.Errorf("got %v, want %v ", got , want)
		}
		if err := mockSql.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
