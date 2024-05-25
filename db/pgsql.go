package db

import (
	"context"
	"database/sql"
	"fmt"
	"perqara_testing/ent"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func InitToYmal() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("perqara")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func SetupDatabasePgSQLConnection(ctx context.Context) (*bun.DB, error) {
	err := InitToYmal()
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, err
	}

	user := viper.GetString("sqldb.pgsql.user")
	password := viper.GetString("sqldb.pgsql.password")
	host := viper.GetString("sqldb.pgsql.host")
	port := viper.GetString("sqldb.pgsql.port")
	database := viper.GetString("sqldb.pgsql.database")
	maxOpenConn := viper.GetInt("sqldb.pgsql.maxopenconn")
	maxIdleConn := viper.GetInt("sqldb.pgsql.maxidleconn")

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
	orm.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	orm.NewCreateTable().Model((*ent.VendingMachines)(nil)).Exec(ctx)

	return orm, nil
}
