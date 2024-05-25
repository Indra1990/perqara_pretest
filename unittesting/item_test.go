package unittesting

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"perqara_testing/api"
	"perqara_testing/api/vendingmachine/usecase"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func SetupDatabasePgSQLConnection(ctx context.Context) (*bun.DB, error) {
	user := "postgres"
	password := "postgres"
	host := "localhost"
	port := "5432"
	database := "transactions"
	maxOpenConn := 30
	maxIdleConn := 1000

	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Jakarta", user, password, host, port, database)
	pgsldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(psqlconn)))
	pgsldb.SetMaxOpenConns(maxOpenConn)
	pgsldb.SetMaxIdleConns(maxIdleConn)

	log.Info("Ping PostgreSQL database ...")

	if err := pgsldb.Ping(); err != nil {
		log.Error(err.Error())
		panic(err)
	}

	log.Info("Ping PostgreSQL database OK ")

	orm := bun.NewDB(pgsldb, pgdialect.New())

	return orm, nil
}

func TestListItem(t *testing.T) {
	ctx := context.Background()

	conn, dbErr := SetupDatabasePgSQLConnection(ctx)
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)
	}

	req, err := http.NewRequest(http.MethodGet, "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	api.Routes(conn).ServeHTTP(newRecorder, req)
	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}

	t.Log(newRecorder)

}
func TestDetailItem(t *testing.T) {
	ctx := context.Background()

	conn, dbErr := SetupDatabasePgSQLConnection(ctx)
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)
	}

	url := "/item/5"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	api.Routes(conn).ServeHTTP(newRecorder, req)
	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf(" test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}

	t.Log(newRecorder)
}

func TestDeleteItem(t *testing.T) {
	ctx := context.Background()

	conn, dbErr := SetupDatabasePgSQLConnection(ctx)
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)
	}

	url := "/item/3"
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	api.Routes(conn).ServeHTTP(newRecorder, req)
	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}

	t.Log(newRecorder)
}

func TestCreateItem(t *testing.T) {
	ctx := context.Background()

	conn, dbErr := SetupDatabasePgSQLConnection(ctx)
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)
	}

	paramBody := usecase.VendingMachineRequestCommand{
		Item:  "testing item",
		Price: "500000",
	}

	requestBody, err := json.Marshal(paramBody)
	if err != nil {
		t.Fatal(err)

	}

	url := "/item/create"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	api.Routes(conn).ServeHTTP(newRecorder, req)
	statusCode := 201
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}

	t.Log(newRecorder)
}

func TestUpdateItem(t *testing.T) {
	ctx := context.Background()

	conn, dbErr := SetupDatabasePgSQLConnection(ctx)
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)
	}

	paramBody := usecase.VendingMachineRequestCommand{
		Item:  "test update",
		Price: "12000",
	}

	requestBody, err := json.Marshal(paramBody)
	if err != nil {
		t.Fatal(err)

	}

	url := "/item/7"
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	api.Routes(conn).ServeHTTP(newRecorder, req)
	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}

	t.Log(newRecorder)
}
