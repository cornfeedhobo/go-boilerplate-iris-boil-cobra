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

// BoilerplateUser is an object representing the database table.
type BoilerplateUser struct {
	ID        int       `db:"id" boil:"id" json:"id" xml:"id"`
	Username  string    `db:"username" boil:"username" json:"username" xml:"username"`
	Password  string    `db:"password" boil:"password" json:"password" xml:"password"`
	RoleID    int       `db:"role_id" boil:"role_id" json:"role_id" xml:"role_id"`
	CreatedAt null.Time `db:"created_at" boil:"created_at" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt null.Time `db:"updated_at" boil:"updated_at" json:"updated_at,omitempty" xml:"updated_at,omitempty"`

	R *boilerplateUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L boilerplateUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// boilerplateUserR is where relationships are stored.
type boilerplateUserR struct {
	Role *BoilerplateRole
}

// boilerplateUserL is where Load methods for each relationship are stored.
type boilerplateUserL struct{}

var (
	boilerplateUserColumns               = []string{"id", "username", "password", "role_id", "created_at", "updated_at"}
	boilerplateUserColumnsWithoutDefault = []string{"username", "password", "role_id"}
	boilerplateUserColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	boilerplateUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// BoilerplateUserSlice is an alias for a slice of pointers to BoilerplateUser.
	// This should generally be used opposed to []BoilerplateUser.
	BoilerplateUserSlice []*BoilerplateUser
	// BoilerplateUserHook is the signature for custom BoilerplateUser hook methods
	BoilerplateUserHook func(boil.Executor, *BoilerplateUser) error

	boilerplateUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	boilerplateUserType                 = reflect.TypeOf(&BoilerplateUser{})
	boilerplateUserMapping              = queries.MakeStructMapping(boilerplateUserType)
	boilerplateUserPrimaryKeyMapping, _ = queries.BindMapping(boilerplateUserType, boilerplateUserMapping, boilerplateUserPrimaryKeyColumns)
	boilerplateUserInsertCacheMut       sync.RWMutex
	boilerplateUserInsertCache          = make(map[string]insertCache)
	boilerplateUserUpdateCacheMut       sync.RWMutex
	boilerplateUserUpdateCache          = make(map[string]updateCache)
	boilerplateUserUpsertCacheMut       sync.RWMutex
	boilerplateUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var boilerplateUserBeforeInsertHooks []BoilerplateUserHook
var boilerplateUserBeforeUpdateHooks []BoilerplateUserHook
var boilerplateUserBeforeDeleteHooks []BoilerplateUserHook
var boilerplateUserBeforeUpsertHooks []BoilerplateUserHook

var boilerplateUserAfterInsertHooks []BoilerplateUserHook
var boilerplateUserAfterSelectHooks []BoilerplateUserHook
var boilerplateUserAfterUpdateHooks []BoilerplateUserHook
var boilerplateUserAfterDeleteHooks []BoilerplateUserHook
var boilerplateUserAfterUpsertHooks []BoilerplateUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BoilerplateUser) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BoilerplateUser) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BoilerplateUser) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BoilerplateUser) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BoilerplateUser) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BoilerplateUser) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BoilerplateUser) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BoilerplateUser) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BoilerplateUser) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range boilerplateUserAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBoilerplateUserHook registers your hook function for all future operations.
func AddBoilerplateUserHook(hookPoint boil.HookPoint, boilerplateUserHook BoilerplateUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		boilerplateUserBeforeInsertHooks = append(boilerplateUserBeforeInsertHooks, boilerplateUserHook)
	case boil.BeforeUpdateHook:
		boilerplateUserBeforeUpdateHooks = append(boilerplateUserBeforeUpdateHooks, boilerplateUserHook)
	case boil.BeforeDeleteHook:
		boilerplateUserBeforeDeleteHooks = append(boilerplateUserBeforeDeleteHooks, boilerplateUserHook)
	case boil.BeforeUpsertHook:
		boilerplateUserBeforeUpsertHooks = append(boilerplateUserBeforeUpsertHooks, boilerplateUserHook)
	case boil.AfterInsertHook:
		boilerplateUserAfterInsertHooks = append(boilerplateUserAfterInsertHooks, boilerplateUserHook)
	case boil.AfterSelectHook:
		boilerplateUserAfterSelectHooks = append(boilerplateUserAfterSelectHooks, boilerplateUserHook)
	case boil.AfterUpdateHook:
		boilerplateUserAfterUpdateHooks = append(boilerplateUserAfterUpdateHooks, boilerplateUserHook)
	case boil.AfterDeleteHook:
		boilerplateUserAfterDeleteHooks = append(boilerplateUserAfterDeleteHooks, boilerplateUserHook)
	case boil.AfterUpsertHook:
		boilerplateUserAfterUpsertHooks = append(boilerplateUserAfterUpsertHooks, boilerplateUserHook)
	}
}

// OneP returns a single boilerplateUser record from the query, and panics on error.
func (q boilerplateUserQuery) OneP() *BoilerplateUser {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single boilerplateUser record from the query.
func (q boilerplateUserQuery) One() (*BoilerplateUser, error) {
	o := &BoilerplateUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for boilerplate_user")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all BoilerplateUser records from the query, and panics on error.
func (q boilerplateUserQuery) AllP() BoilerplateUserSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all BoilerplateUser records from the query.
func (q boilerplateUserQuery) All() (BoilerplateUserSlice, error) {
	var o BoilerplateUserSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BoilerplateUser slice")
	}

	if len(boilerplateUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all BoilerplateUser records in the query, and panics on error.
func (q boilerplateUserQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all BoilerplateUser records in the query.
func (q boilerplateUserQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count boilerplate_user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q boilerplateUserQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q boilerplateUserQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if boilerplate_user exists")
	}

	return count > 0, nil
}

// RoleG pointed to by the foreign key.
func (o *BoilerplateUser) RoleG(mods ...qm.QueryMod) boilerplateRoleQuery {
	return o.Role(boil.GetDB(), mods...)
}

// Role pointed to by the foreign key.
func (o *BoilerplateUser) Role(exec boil.Executor, mods ...qm.QueryMod) boilerplateRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=$1", o.RoleID),
	}

	queryMods = append(queryMods, mods...)

	query := BoilerplateRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"boilerplate\".\"boilerplate_role\"")

	return query
}

// LoadRole allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (boilerplateUserL) LoadRole(e boil.Executor, singular bool, maybeBoilerplateUser interface{}) error {
	var slice []*BoilerplateUser
	var object *BoilerplateUser

	count := 1
	if singular {
		object = maybeBoilerplateUser.(*BoilerplateUser)
	} else {
		slice = *maybeBoilerplateUser.(*BoilerplateUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &boilerplateUserR{}
		args[0] = object.RoleID
	} else {
		for i, obj := range slice {
			obj.R = &boilerplateUserR{}
			args[i] = obj.RoleID
		}
	}

	query := fmt.Sprintf(
		"select * from \"boilerplate\".\"boilerplate_role\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BoilerplateRole")
	}
	defer results.Close()

	var resultSlice []*BoilerplateRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BoilerplateRole")
	}

	if len(boilerplateUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Role = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.RoleID == foreign.ID {
				local.R.Role = foreign
				break
			}
		}
	}

	return nil
}

// SetRole of the boilerplate_user to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleBoilerplateUsers.
func (o *BoilerplateUser) SetRole(exec boil.Executor, insert bool, related *BoilerplateRole) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"boilerplate\".\"boilerplate_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
		strmangle.WhereClause("\"", "\"", 2, boilerplateUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RoleID = related.ID

	if o.R == nil {
		o.R = &boilerplateUserR{
			Role: related,
		}
	} else {
		o.R.Role = related
	}

	if related.R == nil {
		related.R = &boilerplateRoleR{
			RoleBoilerplateUsers: BoilerplateUserSlice{o},
		}
	} else {
		related.R.RoleBoilerplateUsers = append(related.R.RoleBoilerplateUsers, o)
	}

	return nil
}

// BoilerplateUsersG retrieves all records.
func BoilerplateUsersG(mods ...qm.QueryMod) boilerplateUserQuery {
	return BoilerplateUsers(boil.GetDB(), mods...)
}

// BoilerplateUsers retrieves all the records using an executor.
func BoilerplateUsers(exec boil.Executor, mods ...qm.QueryMod) boilerplateUserQuery {
	mods = append(mods, qm.From("\"boilerplate\".\"boilerplate_user\""))
	return boilerplateUserQuery{NewQuery(exec, mods...)}
}

// FindBoilerplateUserG retrieves a single record by ID.
func FindBoilerplateUserG(id int, selectCols ...string) (*BoilerplateUser, error) {
	return FindBoilerplateUser(boil.GetDB(), id, selectCols...)
}

// FindBoilerplateUserGP retrieves a single record by ID, and panics on error.
func FindBoilerplateUserGP(id int, selectCols ...string) *BoilerplateUser {
	retobj, err := FindBoilerplateUser(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindBoilerplateUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBoilerplateUser(exec boil.Executor, id int, selectCols ...string) (*BoilerplateUser, error) {
	boilerplateUserObj := &BoilerplateUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"boilerplate\".\"boilerplate_user\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(boilerplateUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from boilerplate_user")
	}

	return boilerplateUserObj, nil
}

// FindBoilerplateUserP retrieves a single record by ID with an executor, and panics on error.
func FindBoilerplateUserP(exec boil.Executor, id int, selectCols ...string) *BoilerplateUser {
	retobj, err := FindBoilerplateUser(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BoilerplateUser) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *BoilerplateUser) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *BoilerplateUser) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *BoilerplateUser) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no boilerplate_user provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(boilerplateUserColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	boilerplateUserInsertCacheMut.RLock()
	cache, cached := boilerplateUserInsertCache[key]
	boilerplateUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			boilerplateUserColumns,
			boilerplateUserColumnsWithDefault,
			boilerplateUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(boilerplateUserType, boilerplateUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(boilerplateUserType, boilerplateUserMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"boilerplate\".\"boilerplate_user\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into boilerplate_user")
	}

	if !cached {
		boilerplateUserInsertCacheMut.Lock()
		boilerplateUserInsertCache[key] = cache
		boilerplateUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single BoilerplateUser record. See Update for
// whitelist behavior description.
func (o *BoilerplateUser) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single BoilerplateUser record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *BoilerplateUser) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the BoilerplateUser, and panics on error.
// See Update for whitelist behavior description.
func (o *BoilerplateUser) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the BoilerplateUser.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *BoilerplateUser) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	boilerplateUserUpdateCacheMut.RLock()
	cache, cached := boilerplateUserUpdateCache[key]
	boilerplateUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(boilerplateUserColumns, boilerplateUserPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update boilerplate_user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"boilerplate\".\"boilerplate_user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, boilerplateUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(boilerplateUserType, boilerplateUserMapping, append(wl, boilerplateUserPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update boilerplate_user row")
	}

	if !cached {
		boilerplateUserUpdateCacheMut.Lock()
		boilerplateUserUpdateCache[key] = cache
		boilerplateUserUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q boilerplateUserQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q boilerplateUserQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for boilerplate_user")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BoilerplateUserSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o BoilerplateUserSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o BoilerplateUserSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BoilerplateUserSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boilerplateUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"boilerplate\".\"boilerplate_user\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(boilerplateUserPrimaryKeyColumns), len(colNames)+1, len(boilerplateUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in boilerplateUser slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BoilerplateUser) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *BoilerplateUser) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *BoilerplateUser) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *BoilerplateUser) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no boilerplate_user provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(boilerplateUserColumnsWithDefault, o)

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

	boilerplateUserUpsertCacheMut.RLock()
	cache, cached := boilerplateUserUpsertCache[key]
	boilerplateUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			boilerplateUserColumns,
			boilerplateUserColumnsWithDefault,
			boilerplateUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			boilerplateUserColumns,
			boilerplateUserPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert boilerplate_user, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(boilerplateUserPrimaryKeyColumns))
			copy(conflict, boilerplateUserPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"boilerplate\".\"boilerplate_user\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(boilerplateUserType, boilerplateUserMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(boilerplateUserType, boilerplateUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for boilerplate_user")
	}

	if !cached {
		boilerplateUserUpsertCacheMut.Lock()
		boilerplateUserUpsertCache[key] = cache
		boilerplateUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single BoilerplateUser record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *BoilerplateUser) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single BoilerplateUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BoilerplateUser) DeleteG() error {
	if o == nil {
		return errors.New("models: no BoilerplateUser provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single BoilerplateUser record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *BoilerplateUser) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single BoilerplateUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BoilerplateUser) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no BoilerplateUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), boilerplateUserPrimaryKeyMapping)
	sql := "DELETE FROM \"boilerplate\".\"boilerplate_user\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from boilerplate_user")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q boilerplateUserQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q boilerplateUserQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no boilerplateUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from boilerplate_user")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o BoilerplateUserSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o BoilerplateUserSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no BoilerplateUser slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o BoilerplateUserSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BoilerplateUserSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no BoilerplateUser slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(boilerplateUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boilerplateUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"boilerplate\".\"boilerplate_user\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, boilerplateUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(boilerplateUserPrimaryKeyColumns), 1, len(boilerplateUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from boilerplateUser slice")
	}

	if len(boilerplateUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *BoilerplateUser) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *BoilerplateUser) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BoilerplateUser) ReloadG() error {
	if o == nil {
		return errors.New("models: no BoilerplateUser provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BoilerplateUser) Reload(exec boil.Executor) error {
	ret, err := FindBoilerplateUser(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BoilerplateUserSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BoilerplateUserSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BoilerplateUserSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BoilerplateUserSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BoilerplateUserSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	boilerplateUsers := BoilerplateUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), boilerplateUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"boilerplate\".\"boilerplate_user\".* FROM \"boilerplate\".\"boilerplate_user\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, boilerplateUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(boilerplateUserPrimaryKeyColumns), 1, len(boilerplateUserPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&boilerplateUsers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BoilerplateUserSlice")
	}

	*o = boilerplateUsers

	return nil
}

// BoilerplateUserExists checks if the BoilerplateUser row exists.
func BoilerplateUserExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"boilerplate\".\"boilerplate_user\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if boilerplate_user exists")
	}

	return exists, nil
}

// BoilerplateUserExistsG checks if the BoilerplateUser row exists.
func BoilerplateUserExistsG(id int) (bool, error) {
	return BoilerplateUserExists(boil.GetDB(), id)
}

// BoilerplateUserExistsGP checks if the BoilerplateUser row exists. Panics on error.
func BoilerplateUserExistsGP(id int) bool {
	e, err := BoilerplateUserExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// BoilerplateUserExistsP checks if the BoilerplateUser row exists. Panics on error.
func BoilerplateUserExistsP(exec boil.Executor, id int) bool {
	e, err := BoilerplateUserExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
