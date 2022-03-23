package main

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindServicesSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"service_id", "title", "summary", "org_id", "version_count"}).
		AddRow("1", "Title 1", "Summary 1", "1", "1").
		AddRow("2", "Title 2", "Summary 2", "1", "1").
		AddRow("3", "Title 3", "Summary 3", "1", "1")

	mock.ExpectQuery("^SELECT (.+) FROM services s (.+) WHERE u.user_id = ?").
		WillReturnRows(rows)

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	services, err := data.FindServices(ctx)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 3, len(services))
	assert.Equal(t, Service{"1", "Title 1", "Summary 1", "1", 1}, services[0])
	assert.Equal(t, Service{"2", "Title 2", "Summary 2", "1", 1}, services[1])
	assert.Equal(t, Service{"3", "Title 3", "Summary 3", "1", 1}, services[2])
}

func TestFindServicesError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	mock.ExpectQuery("^SELECT (.+) FROM services s (.+) WHERE u.user_id = ?").
		WillReturnError(errors.New("sql error"))

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	services, err := data.FindServices(ctx)

	assert.Error(t, err, "sql error")
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Nil(t, services)
}

func TestFindVersionsForServiceSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"service_id", "version_id", "summary"}).
		AddRow("1", "1", "Summary 1").
		AddRow("1", "2", "Summary 2").
		AddRow("1", "3", "Summary 3")

	mock.ExpectQuery("^SELECT (.+) FROM versions").
		WillReturnRows(rows)

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	versions, err := data.FindVersionsForService(ctx, "1")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 3, len(versions))
	assert.Equal(t, Version{"1", "1", "Summary 1"}, versions[0])
	assert.Equal(t, Version{"1", "2", "Summary 2"}, versions[1])
	assert.Equal(t, Version{"1", "3", "Summary 3"}, versions[2])
}

func TestFindVersionsForServiceError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	mock.ExpectQuery("^SELECT (.+) FROM versions").
		WillReturnError(errors.New("sql error"))

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	services, err := data.FindVersionsForService(ctx, "1")

	assert.Error(t, err, "sql error")
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Nil(t, services)
}

func TestFindServiceByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"service_id", "title", "summary", "org_id"}).
		AddRow("1", "Title 1", "Summary 1", "1")

	mock.ExpectQuery("^SELECT (.+) FROM services s (.+)").
		WillReturnRows(rows)

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	s, err := data.FindServiceByID(ctx, "1")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NotNil(t, s)
	assert.Equal(t, "1", s.ServiceID)
	assert.Equal(t, "1", s.OrgID)
	assert.Equal(t, "Title 1", s.Title)
	assert.Equal(t, "Summary 1", s.Summary)
}

func TestFindServiceByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"service_id", "title", "summary", "org_id"})

	mock.ExpectQuery("^SELECT (.+) FROM services").
		WillReturnRows(rows)

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	s, err := data.FindServiceByID(ctx, "1")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Nil(t, s)
}

func TestFindServiceByIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	data := SQLDataService{
		db: db,
	}

	mock.ExpectQuery("^SELECT (.+) FROM services").
		WillReturnError(errors.New("sql error"))

	ctx := context.WithValue(context.TODO(), UserIDKey, "1")
	s, err := data.FindServiceByID(ctx, "1")

	assert.Error(t, err, "sql error")
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Nil(t, s)
}
