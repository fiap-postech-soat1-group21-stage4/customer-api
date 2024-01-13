package repository_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/repository"
// 	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// const (
// 	createCustomer = `
// 	INSERT INTO "customer" (.+)
// 	VALUES (.+)
// 	ON CONFLICT DO NOTHING RETURNING "id"
// `

// 	fetchNotification = `SELECT (.+) FROM "customer"  WHERE (.+)`
// )

// var (
// 	defaultID = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")

// 	ctx = context.Background()

// 	customer = &entity.Customer{
// 		//	Name:  "João",
// 		CPF: "12345678933",
// 		//Email: "joao@email.com",
// 	}
// )

// func TestCreateCustomer(t *testing.T) {
// 	t.Run("when everything goes ok, should create a customer register", func(t *testing.T) {
// 		db, mock, _ := sqlmock.New()
// 		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
// 		defer db.Close()

// 		mock.ExpectBegin()
// 		mock.
// 			ExpectQuery(createCustomer).
// 			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(defaultID))
// 		mock.ExpectCommit()

// 		r := repository.New(dbGorm)
// 		_, err := r.Customer.CreateCustomer(ctx, customer)

// 		assert.NoError(t, err)
// 		//assert.Equal(t, result.ID, defaultID)
// 		//assert.NoError(t, mock.ExpectationsWereMet())
// 	})

// 	t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
// 		db, mock, _ := sqlmock.New()
// 		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
// 		defer db.Close()
// 		wantErr := errors.New("iamanerror")

// 		mock.ExpectBegin()
// 		mock.ExpectQuery(createCustomer).WillReturnError(wantErr)
// 		mock.ExpectRollback()

// 		r := repository.New(dbGorm)
// 		result, err := r.Customer.CreateCustomer(ctx, customer)
// 		assert.ErrorIs(t, err, wantErr)
// 		assert.Nil(t, result)
// 	})
// }

// func TestFetchNotify(t *testing.T) {
// 	t.Run("when everything goes ok, should return a notification list register", func(t *testing.T) {
// 		db, mock, _ := sqlmock.New()
// 		dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
// 		defer db.Close()

// 		wantRows := &entity.Customer{
// 			ID:    uuid.MustParse("24ecd959-96fb-4394-86e9-eb8f51878bf5"),
// 			Name:  "bia",
// 			CPF:   "222222222",
// 			Email: "b@email.com",
// 		}

// 		rows := sqlmock.NewRows([]string{"id", "name", "cpf", "email"}).
// 			AddRow(
// 				wantRows.ID,
// 				wantRows.Name,
// 				wantRows.CPF,
// 				wantRows.Email,
// 			)

// 		mock.
// 			ExpectQuery(fetchNotification).WillReturnRows(rows)

// 		r := repository.New(dbGorm)
// 		got, err := r.Customer.RetrieveCustomer(ctx, customer)

// 		assert.NoError(t, err)
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 		assert.Equal(t, wantRows, got)
// 	})

// t.Run("when db returns unmapped error, should propagate the error", func(t *testing.T) {
// 	db, mock, _ := sqlmock.New()
// 	dbGorm, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
// 	defer db.Close()
// 	wantErr := errors.New("iamanerror")

// 	mock.
// 		ExpectQuery(fetchNotification).WillReturnError(wantErr)

// 	r := repository.New(dbGorm)
// 	got, err := r.Notify.FetchNotify(ctx)

// 	assert.ErrorIs(t, err, wantErr)
// 	assert.Nil(t, got)
// })
//}
