// Copyright © 2016 cornfeedhobo <cornfeedhobo@vfemail.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"time"

	"github.com/cornfeedhobo/go-boilerplate-iris-boil-cobra/util"
	"github.com/jmoiron/sqlx"
	"github.com/ory-am/ladon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tablePrefix = "boilerplate"

var migrateCmd = &cobra.Command{Use: "migrate"}

func init() {
	viper.RegisterAlias("migrate.postgres", "postgres")
	rootCmd.AddCommand(migrateCmd)
}

type Migration struct {
	ID          uint64               `db:"id" json:"id"`
	Description string               `db:"description" json:"description"`
	CreatedAt   time.Time            `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time            `db:"updated_at" json:"updated_at"`
	MigrateUp   func(*sqlx.DB) error `db:"-" json:"-"`
	MigrateDown func(*sqlx.DB) error `db:"-" json:"-"`
}

var Migrations []*Migration = []*Migration{
	{
		ID:          1,
		Description: "ladon tables",
		MigrateUp: func(db *sqlx.DB) error {
			return ladon.NewSQLManager(db, nil).CreateSchemas()
		},
		MigrateDown: func(db *sqlx.DB) error {
			var tables []string
			if err := db.Select(&tables, util.HereDoc(`
				SELECT table_name
				FROM information_schema.tables
				WHERE table_name LIKE 'ladon_%';
			`)); err != nil {
				return err
			}
			for _, t := range tables {
				if _, err := db.Exec(`DROP IF EXISTS ` + t + ` CASCADE;`); err != nil {
					return err
				}
			}
			return nil
		},
	},
	{
		ID:          2,
		Description: "boilerplate tables",
		MigrateUp: func(db *sqlx.DB) error {
			tx, err := db.Begin()
			if err != nil {
				return err
			}
			if _, err = tx.Exec(util.HereDocf(`
				CREATE TABLE IF NOT EXISTS %[1]s_role (
					id SERIAL PRIMARY KEY,
					policy_id VARCHAR(255) NOT NULL,
					created_at TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc'),
					updated_at TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc')
				);
				ALTER TABLE %[1]s_role ADD CONSTRAINT fk_%[1]s_role_policy_id FOREIGN KEY (policy_id) REFERENCES ladon_policy(id);
				CREATE TABLE IF NOT EXISTS %[1]s_user (
					id SERIAL PRIMARY KEY,
					username TEXT NOT NULL,
					password TEXT NOT NULL,
					role_id INTEGER NOT NULL,
					created_at TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc'),
					updated_at TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc')
				);
				ALTER TABLE %[1]s_user ADD CONSTRAINT fk_%[1]s_user_role_id FOREIGN KEY (role_id) REFERENCES %[1]s_role(id);
			`, tablePrefix)); err != nil {
				return err
			}
			if err := tx.Commit(); err != nil {
				return err
			}
			return nil
		},
		MigrateDown: func(db *sqlx.DB) error {
			tx, err := db.Begin()
			if err != nil {
				return err
			}
			if _, err = tx.Exec(util.HereDocf(`
				DROP TABLE IF EXISTS %s_user CASCADE;
				DROP TABLE IF EXISTS %s_role CASCADE;
			`, tablePrefix)); err != nil {
				return err
			}
			if err = tx.Commit(); err != nil {
				return err
			}
			return nil
		},
	},
}

func MigrationsTableExists(db *sqlx.DB) bool {
	var id int
	db.Get(&id, fmt.Sprintf(`SELECT count(*) FROM information_schema.tables WHERE table_name = '%[1]s_migration';`, tablePrefix))
	if id > 0 {
		return true
	}
	return false
}

func GetCurrentMigration(db *sqlx.DB) *Migration {
	if !MigrationsTableExists(db) {
		return nil
	}
	m := &Migration{}
	db.Get(m, fmt.Sprintf(`SELECT * FROM %[1]s_migration ORDER BY id DESC LIMIT 1;`, tablePrefix))
	return m
}

func CreateMigrationsTable(db *sqlx.DB) error {
	_, err := db.Exec(util.HereDocf(`
		CREATE TABLE IF NOT EXISTS %[1]s_migration (
			id SERIAL PRIMARY KEY,
			description TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc'),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE 'utc')
		);
	`, tablePrefix))
	return err
}

func CreateAuditTriggers(db *sqlx.DB) error {
	var tables []string
	db.Select(&tables, util.HereDoc(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema NOT IN ('pg_catalog', 'information_schema')
			AND table_schema NOT LIKE 'pg_toast%';
	`))
	for _, table := range tables {
		tx, err := db.Beginx()
		if err != nil {
			return err
		}
		tx.Exec(util.HereDocf(`
			DROP TRIGGER IF EXISTS history_trigger ON %[1]s RESTRICT;
			CREATE TRIGGER history_trigger BEFORE INSERT OR UPDATE OR DELETE ON %[1]s
			FOR EACH ROW EXECUTE PROCEDURE table_history_trigger();
		`, table))
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}
