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
	"errors"
	"fmt"
	"strings"
	"syscall"

	"github.com/cornfeedhobo/go-boilerplate-iris-boil-cobra/util"
	"github.com/elithrar/simple-scrypt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var userCmd = &cobra.Command{
	Use:       "user <username>",
	ValidArgs: []string{"username"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("invalid number of arguments %d", len(args))
		}

		username := strings.TrimSpace(args[0])
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			fmt.Print("Password: ")
			bytePassword, err := terminal.ReadPassword(syscall.Stdin)
			if err != nil {
				return err
			}
			fmt.Println()
			password = strings.TrimSpace(string(bytePassword))
		}

		role, _ := cmd.Flags().GetString("role")

		db, err := sqlx.Connect("pgx", viper.GetString("migrate.up.postgres"))
		if err != nil {
			return fmt.Errorf("unable to connect to database: %s", err)
		}
		defer func(db *sqlx.DB) error {
			if err := db.Close(); err != nil {
				return fmt.Errorf("error closing database connection: %s", err)
			}
			return nil
		}(db)

		var role_id int
		if db.Get(&role_id, util.HereDocf(`
			SELECT id FROM %[1]s_role WHERE policy_id = $1;
		`, tablePrefix), role); role_id == 0 {
			return errors.New("invalid role")
		}

		passwordHash, err := scrypt.GenerateFromPassword([]byte(password), scrypt.DefaultParams)
		if err != nil {
			return fmt.Errorf("error generating hash: %s", err)
		}

		if _, err := db.NamedExec(util.HereDocf(`
			INSERT INTO %[1]s_user (username, password, role_id)
			VALUES (:username, :password, :role_id)
		`, tablePrefix), map[string]interface{}{
			"username": username,
			"password": string(passwordHash),
			"role_id":  role_id,
		}); err != nil {
			return fmt.Errorf("unable to create user: %s", err)
		}

		fmt.Printf("created user '%s'\n", username)
		return nil
	},
}

func init() {
	viper.BindPFlag("create.user.postgres", userCmd.PersistentFlags().Lookup("create.postgres"))

	userCmd.Flags().String("password", "", "user password (leave blank to be prompted)")
	userCmd.Flags().String("role", "admin", "user role")
	createCmd.AddCommand(userCmd)
}
