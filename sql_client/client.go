package sql_client

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type SQLClient struct {
	*sqlx.DB
}

func (client *SQLClient) Get() *sqlx.DB {
	return client.DB
}

// NewSqlxDB type;
// For MySQL, posgreSQL
func NewSqlxDB(c *SQLConfig) *SQLClient {
	db, err := sqlx.Connect(c.Driver, c.DSN)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	if c.Lifetime < 60 {
		c.Lifetime = 5 * 60
	}
	db.SetConnMaxLifetime(time.Duration(c.Lifetime) * time.Second)

	return &SQLClient{db}
}
