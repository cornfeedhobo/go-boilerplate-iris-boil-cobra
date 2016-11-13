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

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v5"
)

// BoilerplateRole is an object representing the database table.
type BoilerplateRole struct {
	ID        int       `db:"id" boil:"id" json:"id" xml:"id"`
	PolicyID  string    `db:"policy_id" boil:"policy_id" json:"policy_id" xml:"policy_id"`
	CreatedAt null.Time `db:"created_at" boil:"created_at" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt null.Time `db:"updated_at" boil:"updated_at" json:"updated_at,omitempty" xml:"updated_at,omitempty"`

	R *boilerplateRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L boilerplateRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// boilerplateRoleR is where relationships are stored.
type boilerplateRoleR struct {
	Policy               *LadonPolicy
	RoleBoilerplateUsers BoilerplateUserSlice
}

// boilerplateRoleL is where Load methods for each relationship are stored.
type boilerplateRoleL struct{}

var (
	boilerplateRoleColumns               = []string{"id", "policy_id", "created_at", "updated_at"}
	boilerplateRoleColumnsWithoutDefault = []string{"policy_id"}
	boilerplateRoleColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	boilerplateRolePrimaryKeyColumns     = []string{"id"}
)

type (
	// BoilerplateRoleSlice is an alias for a slice of pointers to BoilerplateRole.
	// This should generally be used opposed to []BoilerplateRole.
	BoilerplateRoleSlice []*BoilerplateRole
	// BoilerplateRoleHook is the signature for custom BoilerplateRole hook methods
	BoilerplateRoleHook func(boil.Executor, *BoilerplateRole) error

	boilerplateRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	boilerplateRoleType                 = reflect.TypeOf(&BoilerplateRole{})
	boilerplateRoleMapping              = queries.MakeStructMapping(boilerplateRoleType)
	boilerplateRolePrimaryKeyMapping, _ = queries.BindMapping(boilerplateRoleType, boilerplateRoleMapping, boilerplateRolePrimaryKeyColumns)
	boilerplateRoleInsertCacheMut       sync.RWMutex
	boilerplateRoleInsertCache          = make(map[string]insertCache)
	boilerplateRoleUpdateCacheMut       sync.RWMutex
	boilerplateRoleUpdateCache          = make(map[string]updateCache)
	boilerplateRoleUpsertCacheMut       sync.RWMutex
	boilerplateRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var boilerplateRoleBeforeInsertHooks []BoilerplateRoleHook
var boilerplateRoleBeforeUpdateHooks []BoilerplateRoleHook
var boilerplateRoleBeforeDeleteHooks []BoilerplateRoleHook
var boilerplateRoleBeforeUpsertHooks []BoilerplateRoleHook

var boilerplateRoleAfterInsertHooks []BoilerplateRoleHook
var boilerplateRoleAfterSelectHooks []BoilerplateRoleHook
var boilerplateRoleAfterUpdateHooks []BoilerplateRoleHook
var boilerplateRoleAfterDeleteHooks []BoilerplateRoleHook
var boilerplateRoleAfterUpsertHooks []BoilerplateRoleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BoilerplateRole) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BoilerplateRole) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BoilerplateRole) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BoilerplateRole) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BoilerplateRole) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BoilerplateRole) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BoilerplateRole) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BoilerplateRole) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BoilerplateRole) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateRoleAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBoilerplateRoleHook registers your hook function for all future operations.
func AddBoilerplateRoleHook(hookPoint boil.HookPoint, boilerplateRoleHook BoilerplateRoleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		boilerplateRoleBeforeInsertHooks = append(boilerplateRoleBeforeInsertHooks, boilerplateRoleHook)
	case boil.BeforeUpdateHook:
		boilerplateRoleBeforeUpdateHooks = append(boilerplateRoleBeforeUpdateHooks, boilerplateRoleHook)
	case boil.BeforeDeleteHook:
		boilerplateRoleBeforeDeleteHooks = append(boilerplateRoleBeforeDeleteHooks, boilerplateRoleHook)
	case boil.BeforeUpsertHook:
		boilerplateRoleBeforeUpsertHooks = append(boilerplateRoleBeforeUpsertHooks, boilerplateRoleHook)
	case boil.AfterInsertHook:
		boilerplateRoleAfterInsertHooks = append(boilerplateRoleAfterInsertHooks, boilerplateRoleHook)
	case boil.AfterSelectHook:
		boilerplateRoleAfterSelectHooks = append(boilerplateRoleAfterSelectHooks, boilerplateRoleHook)
	case boil.AfterUpdateHook:
		boilerplateRoleAfterUpdateHooks = append(boilerplateRoleAfterUpdateHooks, boilerplateRoleHook)
	case boil.AfterDeleteHook:
		boilerplateRoleAfterDeleteHooks = append(boilerplateRoleAfterDeleteHooks, boilerplateRoleHook)
	case boil.AfterUpsertHook:
		boilerplateRoleAfterUpsertHooks = append(boilerplateRoleAfterUpsertHooks, boilerplateRoleHook)
	}
}

// OneP returns a single boilerplateRole record from the query, and panics on error.
func (q boilerplateRoleQuery) OneP() *BoilerplateRole {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single boilerplateRole record from the query.
func (q boilerplateRoleQuery) One() (*BoilerplateRole, error) {
	o := &BoilerplateRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for boilerplate_role")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all BoilerplateRole records from the query, and panics on error.
func (q boilerplateRoleQuery) AllP() BoilerplateRoleSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all BoilerplateRole records from the query.
func (q boilerplateRoleQuery) All() (BoilerplateRoleSlice, error) {
	var o BoilerplateRoleSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BoilerplateRole slice")
	}

	if len(boilerplateRoleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all BoilerplateRole records in the query, and panics on error.
func (q boilerplateRoleQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all BoilerplateRole records in the query.
func (q boilerplateRoleQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count boilerplate_role rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q boilerplateRoleQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q boilerplateRoleQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if boilerplate_role exists")
	}

	return count > 0, nil
}

// PolicyG pointed to by the foreign key.
func (o *BoilerplateRole) PolicyG(mods ...qm.QueryMod) ladonPolicyQuery {
	return o.Policy(boil.GetDB(), mods...)
}

// Policy pointed to by the foreign key.
func (o *BoilerplateRole) Policy(exec boil.Executor, mods ...qm.QueryMod) ladonPolicyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=$1", o.PolicyID),
	}

	queryMods = append(queryMods, mods...)

	query := LadonPolicies(exec, queryMods...)
	queries.SetFrom(query.Query, "\"boilerplate\".\"ladon_policy\"")

	return query
}

// RoleBoilerplateUsersG retrieves all the boilerplate_user's boilerplate user via role_id column.
func (o *BoilerplateRole) RoleBoilerplateUsersG(mods ...qm.QueryMod) boilerplateUserQuery {
	return o.RoleBoilerplateUsers(boil.GetDB(), mods...)
}

// RoleBoilerplateUsers retrieves all the boilerplate_user's boilerplate user with an executor via role_id column.
func (o *BoilerplateRole) RoleBoilerplateUsers(exec boil.Executor, mods ...qm.QueryMod) boilerplateUserQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"role_id\"=$1", o.ID),
	)

	query := BoilerplateUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"boilerplate\".\"boilerplate_user\" as \"a\"")
	return query
}

// LoadPolicy allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (boilerplateRoleL) LoadPolicy(e boil.Executor, singular bool, maybeBoilerplateRole interface{}) error {
	var slice []*BoilerplateRole
	var object *BoilerplateRole

	count := 1
	if singular {
		object = maybeBoilerplateRole.(*BoilerplateRole)
	} else {
		slice = *maybeBoilerplateRole.(*BoilerplateRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &boilerplateRoleR{}
		args[0] = object.PolicyID
	} else {
		for i, obj := range slice {
			obj.R = &boilerplateRoleR{}
			args[i] = obj.PolicyID
		}
	}

	query := fmt.Sprintf(
		"select * from \"boilerplate\".\"ladon_policy\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load LadonPolicy")
	}
	defer results.Close()

	var resultSlice []*LadonPolicy
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice LadonPolicy")
	}

	if len(boilerplateRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Policy = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PolicyID == foreign.ID {
				local.R.Policy = foreign
				break
			}
		}
	}

	return nil
}

// LoadRoleBoilerplateUsers allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (boilerplateRoleL) LoadRoleBoilerplateUsers(e boil.Executor, singular bool, maybeBoilerplateRole interface{}) error {
	var slice []*BoilerplateRole
	var object *BoilerplateRole

	count := 1
	if singular {
		object = maybeBoilerplateRole.(*BoilerplateRole)
	} else {
		slice = *maybeBoilerplateRole.(*BoilerplateRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &boilerplateRoleR{}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			obj.R = &boilerplateRoleR{}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"boilerplate\".\"boilerplate_user\" where \"role_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load boilerplate_user")
	}
	defer results.Close()

	var resultSlice []*BoilerplateUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice boilerplate_user")
	}

	if len(boilerplateUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RoleBoilerplateUsers = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RoleID {
				local.R.RoleBoilerplateUsers = append(local.R.RoleBoilerplateUsers, foreign)
				break
			}
		}
	}

	return nil
}

// SetPolicy of the boilerplate_role to the related item.
// Sets o.R.Policy to related.
// Adds o to related.R.PolicyBoilerplateRoles.
func (o *BoilerplateRole) SetPolicy(exec boil.Executor, insert bool, related *LadonPolicy) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"boilerplate\".\"boilerplate_role\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"policy_id"}),
		strmangle.WhereClause("\"", "\"", 2, boilerplateRolePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PolicyID = related.ID

	if o.R == nil {
		o.R = &boilerplateRoleR{
			Policy: related,
		}
	} else {
		o.R.Policy = related
	}

	if related.R == nil {
		related.R = &ladonPolicyR{
			PolicyBoilerplateRoles: BoilerplateRoleSlice{o},
		}
	} else {
		related.R.PolicyBoilerplateRoles = append(related.R.PolicyBoilerplateRoles, o)
	}

	return nil
}

// AddRoleBoilerplateUsers adds the given related objects to the existing relationships
// of the boilerplate_role, optionally inserting them as new records.
// Appends related to o.R.RoleBoilerplateUsers.
// Sets related.R.Role appropriately.
func (o *BoilerplateRole) AddRoleBoilerplateUsers(exec boil.Executor, insert bool, related ...*BoilerplateUser) error {
	var err error
	for _, rel := range related {
		rel.RoleID = o.ID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "role_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &boilerplateRoleR{
			RoleBoilerplateUsers: related,
		}
	} else {
		o.R.RoleBoilerplateUsers = append(o.R.RoleBoilerplateUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &boilerplateUserR{
				Role: o,
			}
		} else {
			rel.R.Role = o
		}
	}
	return nil
}

// BoilerplateRolesG retrieves all records.
func BoilerplateRolesG(mods ...qm.QueryMod) boilerplateRoleQuery {
	return BoilerplateRoles(boil.GetDB(), mods...)
}

// BoilerplateRoles retrieves all the records using an executor.
func BoilerplateRoles(exec boil.Executor, mods ...qm.QueryMod) boilerplateRoleQuery {
	mods = append(mods, qm.From("\"boilerplate\".\"boilerplate_role\""))
	return boilerplateRoleQuery{NewQuery(exec, mods...)}
}

// FindBoilerplateRoleG retrieves a single record by ID.
func FindBoilerplateRoleG(id int, selectCols ...string) (*BoilerplateRole, error) {
	return FindBoilerplateRole(boil.GetDB(), id, selectCols...)
}

// FindBoilerplateRoleGP retrieves a single record by ID, and panics on error.
func FindBoilerplateRoleGP(id int, selectCols ...string) *BoilerplateRole {
	retobj, err := FindBoilerplateRole(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindBoilerplateRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBoilerplateRole(exec boil.Executor, id int, selectCols ...string) (*BoilerplateRole, error) {
	boilerplateRoleObj := &BoilerplateRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"boilerplate\".\"boilerplate_role\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(boilerplateRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from boilerplate_role")
	}

	return boilerplateRoleObj, nil
}

// FindBoilerplateRoleP retrieves a single record by ID with an executor, and panics on error.
func FindBoilerplateRoleP(exec boil.Executor, id int, selectCols ...string) *BoilerplateRole {
	retobj, err := FindBoilerplateRole(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BoilerplateRole) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *BoilerplateRole) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *BoilerplateRole) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *BoilerplateRole) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no boilerplate_role provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(boilerplateRoleColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	boilerplateRoleInsertCacheMut.RLock()
	cache, cached := boilerplateRoleInsertCache[key]
	boilerplateRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			boilerplateRoleColumns,
			boilerplateRoleColumnsWithDefault,
			boilerplateRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(boilerplateRoleType, boilerplateRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(boilerplateRoleType, boilerplateRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"boilerplate\".\"boilerplate_role\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into boilerplate_role")
	}

	if !cached {
		boilerplateRoleInsertCacheMut.Lock()
		boilerplateRoleInsertCache[key] = cache
		boilerplateRoleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single BoilerplateRole record. See Update for
// whitelist behavior description.
func (o *BoilerplateRole) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single BoilerplateRole record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *BoilerplateRole) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the BoilerplateRole, and panics on error.
// See Update for whitelist behavior description.
func (o *BoilerplateRole) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the BoilerplateRole.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *BoilerplateRole) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	boilerplateRoleUpdateCacheMut.RLock()
	cache, cached := boilerplateRoleUpdateCache[key]
	boilerplateRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(boilerplateRoleColumns, boilerplateRolePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update boilerplate_role, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"boilerplate\".\"boilerplate_role\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, boilerplateRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(boilerplateRoleType, boilerplateRoleMapping, append(wl, boilerplateRolePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update boilerplate_role row")
	}

	if !cached {
		boilerplateRoleUpdateCacheMut.Lock()
		boilerplateRoleUpdateCache[key] = cache
		boilerplateRoleUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q boilerplateRoleQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q boilerplateRoleQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for boilerplate_role")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BoilerplateRoleSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o BoilerplateRoleSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o BoilerplateRoleSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BoilerplateRoleSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boilerplateRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"boilerplate\".\"boilerplate_role\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(boilerplateRolePrimaryKeyColumns), len(colNames)+1, len(boilerplateRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in boilerplateRole slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BoilerplateRole) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *BoilerplateRole) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *BoilerplateRole) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *BoilerplateRole) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no boilerplate_role provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(boilerplateRoleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	boilerplateRoleUpsertCacheMut.RLock()
	cache, cached := boilerplateRoleUpsertCache[key]
	boilerplateRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			boilerplateRoleColumns,
			boilerplateRoleColumnsWithDefault,
			boilerplateRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			boilerplateRoleColumns,
			boilerplateRolePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert boilerplate_role, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(boilerplateRolePrimaryKeyColumns))
			copy(conflict, boilerplateRolePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"boilerplate\".\"boilerplate_role\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(boilerplateRoleType, boilerplateRoleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(boilerplateRoleType, boilerplateRoleMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for boilerplate_role")
	}

	if !cached {
		boilerplateRoleUpsertCacheMut.Lock()
		boilerplateRoleUpsertCache[key] = cache
		boilerplateRoleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single BoilerplateRole record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *BoilerplateRole) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single BoilerplateRole record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BoilerplateRole) DeleteG() error {
	if o == nil {
		return errors.New("models: no BoilerplateRole provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single BoilerplateRole record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *BoilerplateRole) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single BoilerplateRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BoilerplateRole) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no BoilerplateRole provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), boilerplateRolePrimaryKeyMapping)
	sql := "DELETE FROM \"boilerplate\".\"boilerplate_role\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from boilerplate_role")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q boilerplateRoleQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q boilerplateRoleQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no boilerplateRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from boilerplate_role")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o BoilerplateRoleSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o BoilerplateRoleSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no BoilerplateRole slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o BoilerplateRoleSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BoilerplateRoleSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no BoilerplateRole slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(boilerplateRoleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boilerplateRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"boilerplate\".\"boilerplate_role\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, boilerplateRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(boilerplateRolePrimaryKeyColumns), 1, len(boilerplateRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from boilerplateRole slice")
	}

	if len(boilerplateRoleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *BoilerplateRole) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *BoilerplateRole) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BoilerplateRole) ReloadG() error {
	if o == nil {
		return errors.New("models: no BoilerplateRole provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BoilerplateRole) Reload(exec boil.Executor) error {
	ret, err := FindBoilerplateRole(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BoilerplateRoleSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BoilerplateRoleSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BoilerplateRoleSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BoilerplateRoleSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BoilerplateRoleSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	boilerplateRoles := BoilerplateRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boilerplateRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"boilerplate\".\"boilerplate_role\".* FROM \"boilerplate\".\"boilerplate_role\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, boilerplateRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(boilerplateRolePrimaryKeyColumns), 1, len(boilerplateRolePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&boilerplateRoles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BoilerplateRoleSlice")
	}

	*o = boilerplateRoles

	return nil
}

// BoilerplateRoleExists checks if the BoilerplateRole row exists.
func BoilerplateRoleExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"boilerplate\".\"boilerplate_role\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if boilerplate_role exists")
	}

	return exists, nil
}

// BoilerplateRoleExistsG checks if the BoilerplateRole row exists.
func BoilerplateRoleExistsG(id int) (bool, error) {
	return BoilerplateRoleExists(boil.GetDB(), id)
}

// BoilerplateRoleExistsGP checks if the BoilerplateRole row exists. Panics on error.
func BoilerplateRoleExistsGP(id int) bool {
	e, err := BoilerplateRoleExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// BoilerplateRoleExistsP checks if the BoilerplateRole row exists. Panics on error.
func BoilerplateRoleExistsP(exec boil.Executor, id int) bool {
	e, err := BoilerplateRoleExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
