package ent

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
)

func (c *Client) DB() *sql.DB {
	return c.driver.(*entsql.Driver).DB()
}
