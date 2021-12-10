package database

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	_ "github.com/lib/pq"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn := viper.GetInt("db.max_connections")
	maxIdleConn, _ := strconv.Atoi(viper.GetString("db.max_idle_connections"))
	maxLifetimeConn, _ := strconv.Atoi(viper.GetString("db.max_lifetime_connections"))

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("postgres", dbServerUrl())
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

func dbServerUrl() string {
	var (
		host     = viper.GetString("db.host")
		port     = viper.GetString("db.port")
		username = viper.GetString("db.username")
		dbname   = viper.GetString("db.dbname")
		sslmode  = viper.GetString("db.sslmode")
		password = viper.GetString("db.password")
	)

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", host, port, username, dbname, sslmode, password)
}
