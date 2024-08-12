/*
Package config consolidate all the configuration getter into one place
for centralize and easy adjustment in future neds
*/
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

// IsInDevelopmentStage will check the STAGE environment variable is in development stage
func IsInDevelopmentStage() bool {
	return strings.HasSuffix(os.Getenv("STAGE"), "dev")
}

// CurrentStage return current STAGE environment variable
func CurrentStage() string {
	return os.Getenv("STAGE")
}

// CurrentSchema return current DB_SCHEMA environment variable
func CurrentSchema() string {
	return os.Getenv("DB_SCHEMA")
}

// CurrentDBConnectionURL return cockroachdb.ConnectionURL based on:
// DB_NAME, DB_HOST, DB_USER, DB_PASS environment variables
func CurrentDBConnectionURL() cockroachdb.ConnectionURL {
	dbConnURL := cockroachdb.ConnectionURL{
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}

	return dbConnURL
}

func InitNotifTable(session db.Session) error {
	_, err := session.SQL().Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.push_notification (
		id VARCHAR NOT NULL,
		endpoint VARCHAR NOT NULL,
		keys_p256dh VARCHAR NOT NULL,
		keys_auth VARCHAR NOT NULL,
		created_at TIMESTAMP NULL,
		updated_at TIMESTAMP NULL,
		created_by VARCHAR NULL,
		updated_by VARCHAR NULL,
		CONSTRAINT push_notification_pk PRIMARY KEY (id ASC),
		UNIQUE INDEX push_notification_unique (endpoint ASC, keys_p256dh ASC, keys_auth ASC)
		);`, CurrentSchema()))
	if err != nil {
		return err
	}
	return nil
}

// GetTableNameOnCurrentSchema return the tableName passing parameter with the current schema
//
//	table := GetTableNameOnCurrentSchema("table1") => it will return [schema].table1
func GetTableNameOnCurrentSchema(tableName string) string {
	targetTableName := fmt.Sprintf("%s.%s", CurrentSchema(), tableName)

	return targetTableName
}

func CurrentDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
