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
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cornfeedhobo/go-boilerplate-iris-boil-cobra/util"
	"github.com/jmoiron/sqlx"
	"github.com/ory-am/ladon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var roleCmd = &cobra.Command{
	Use: "role",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		name = strings.TrimSpace(name)
		if name == "" {
			return errors.New("name required")
		}

		policy, _ := cmd.Flags().GetString("policy")
		policy = strings.TrimSpace(policy)
		if policy == "" {
			return errors.New("policy required")
		}

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

		warden := ladon.Ladon{Manager: ladon.NewSQLManager(db, nil)}

		var ladonPolicy ladon.DefaultPolicy
		if err := json.Unmarshal([]byte(policy), &ladonPolicy); err != nil {
			return fmt.Errorf("unable to parse json: %s", err)
		}
		ladonPolicy.ID = name

		if err := warden.Manager.Create(ladon.Policy(&ladonPolicy)); err != nil {
			return fmt.Errorf("unable to create role: %s", err)
		}

		if _, err := db.NamedExec(util.HereDocf(`
			INSERT INTO %[1]s_role (policy_id)
			VALUES (:policy_id);
		`, tablePrefix), map[string]interface{}{"policy_id": ladonPolicy.ID}); err != nil {
			if err2 := warden.Manager.Delete(name); err2 != nil {
				fmt.Printf("unable to delete role: %s", err2)
			}
			return fmt.Errorf("unable to create role: %s", err)
		}

		fmt.Printf("created role '%s'\n", name)
		return nil
	},
}

func init() {
	viper.BindPFlag("create.role.postgres", roleCmd.PersistentFlags().Lookup("create.postgres"))

	roleCmd.Flags().String("name", "admin", "role name")
	roleCmd.Flags().String("policy", util.HereDoc(`
		{
			"description": "admin",
			"subjects": ["<.*>"],
			"resources": ["<.*>"],
			"actions": ["<.*>"],
			"effect": "allow"
		}
	`), "policy json document")

	createCmd.AddCommand(roleCmd)
}
