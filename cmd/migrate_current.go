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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var currentCmd = &cobra.Command{
	Use: "current",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sqlx.Connect("pgx", viper.GetString("migrate.current.postgres"))
		if err != nil {
			return fmt.Errorf("unable to connect to database: %s", err)
		}
		defer func(db *sqlx.DB) error {
			if err := db.Close(); err != nil {
				return fmt.Errorf("error closing database connection: %s", err)
			}
			return nil
		}(db)

		if !MigrationsTableExists(db) {
			return errors.New("run `migrate init` first")
		}

		m := GetCurrentMigration(db)
		if m == nil {
			return errors.New("no migrations found")
		}

		switch viper.GetString("format") {
		case "json":
			b, err := json.Marshal(m)
			if err != nil {
				return fmt.Errorf("error marshalling json: %s", err)
			}

			var out bytes.Buffer
			json.Indent(&out, b, "", "\t")
			fmt.Println(out.String())
		default:
			return fmt.Errorf("unsupported output format: %s", viper.GetString("format"))
		}

		return nil
	},
}

func init() {
	viper.RegisterAlias("migrate.current.postgres", "migrate.postgres")

	currentCmd.Flags().String("format", "json", "output format")
	viper.BindPFlag("format", currentCmd.Flags().Lookup("format"))

	migrateCmd.AddCommand(currentCmd)
}
