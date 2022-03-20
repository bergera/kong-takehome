package main

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindServicesForUserSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"service_id", "title", "summary", "org_id"}).
		AddRow("1", "Title 1", "Summary 1", "1").
		AddRow("2", "Title 2", "Summary 2", "1").
		AddRow("3", "Title 3", "Summary 3", "1")

	mock.ExpectQuery("SELECT (.+) FROM services s JOIN users u ON s.org_id = u.org_id WHERE u.user_id = ?").
		WillReturnRows(rows)

	services, err := data.FindServicesForUser(context.TODO(), 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 3, len(services))
	assert.Equal(t, Service{1, "Title 1", "Summary 1", 1}, services[0])
	assert.Equal(t, Service{2, "Title 2", "Summary 2", 1}, services[1])
	assert.Equal(t, Service{3, "Title 3", "Summary 3", 1}, services[2])
}

func TestFindServicesForUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	mock.ExpectQuery("SELECT (.+) FROM services s JOIN users u ON s.org_id = u.org_id WHERE u.user_id = ?").
		WillReturnError(errors.New("sql error"))

	services, err := data.FindServicesForUser(context.TODO(), 1)

	assert.Error(t, err, "sql error")
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Nil(t, services)
}
