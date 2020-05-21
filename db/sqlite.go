package db

import (
	"database/sql"
	"errors"
	"fmt"
	"mylab/mylabdiscoveries/config"
	"os"
	"strconv"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	logger "github.com/sirupsen/logrus"
)

const (
	dbDriverSQL      = "sqlite3"
	migrationPathSQL = "./migrationsSQL"
)

func InitSQL() (s Storer, err error) {
	uri := config.ReadEnvString("DB_URI_SQLITE")

	conn, err := sqlx.Connect(dbDriverSQL, uri)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot initialize database")
		return
	}

	logger.WithField("uri", uri).Info("Connected to sqlite database")
	return &pgStore{conn}, nil
}

func RunMigrationsSQLite() (err error) {
	uri := config.ReadEnvString("DB_URI_SQLITE")

	db, err := sql.Open(dbDriverSQL, uri)
	if err != nil {
		return
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(getMigrationPathSQLite(), dbDriverSQL, driver)
	if err != nil {
		return
	}

	err = m.Up()
	if err == migrate.ErrNoChange || err == nil {
		err = nil
		return
	}

	return
}
func CreateMigrationFileSQLite(filename string) (err error) {
	if len(filename) == 0 {
		err = errors.New("filename is not provided")
		return
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migrationPathSQL, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migrationPathSQL, timeStamp, filename)

	err = createFile(upMigrationFilePath)
	if err != nil {
		return
	}

	err = createFile(downMigrationFilePath)
	if err != nil {
		os.Remove(upMigrationFilePath)
		return
	}

	logger.WithFields(logger.Fields{
		"up":   upMigrationFilePath,
		"down": downMigrationFilePath,
	}).Info("Created migration files")

	return
}

func RollbackMigrationsSQLite(s string) (err error) {
	uri := config.ReadEnvString("DB_URI_SQLITE")

	steps, err := strconv.Atoi(s)
	if err != nil {
		return
	}

	m, err := migrate.New(getMigrationPathSQLite(), uri)
	if err != nil {
		return
	}

	err = m.Steps(-1 * steps)
	if err == migrate.ErrNoChange || err == nil {
		err = nil
		return
	}

	return
}

func getMigrationPathSQLite() string {
	return fmt.Sprintf("file://%s", migrationPathSQL)
}
