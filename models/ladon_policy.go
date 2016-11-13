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
)

// LadonPolicy is an object representing the database table.
type LadonPolicy struct {
	ID          string `db:"id" boil:"id" json:"id" xml:"id"`
	Description string `db:"description" boil:"description" json:"description" xml:"description"`
	Effect      string `db:"effect" boil:"effect" json:"effect" xml:"effect"`
	Conditions  string `db:"conditions" boil:"conditions" json:"conditions" xml:"conditions"`

	R *ladonPolicyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ladonPolicyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// ladonPolicyR is where relationships are stored.
type ladonPolicyR struct {
	PolicyBoilerplateRoles BoilerplateRoleSlice
}

// ladonPolicyL is where Load methods for each relationship are stored.
type ladonPolicyL struct{}

var (
	ladonPolicyColumns               = []string{"id", "description", "effect", "conditions"}
	ladonPolicyColumnsWithoutDefault = []string{"id", "description", "effect", "conditions"}
	ladonPolicyColumnsWithDefault    = []string{}
	ladonPolicyPrimaryKeyColumns     = []string{"id"}
)

type (
	// LadonPolicySlice is an alias for a slice of pointers to LadonPolicy.
	// This should generally be used opposed to []LadonPolicy.
	LadonPolicySlice []*LadonPolicy
	// LadonPolicyHook is the signature for custom LadonPolicy hook methods
	LadonPolicyHook func(boil.Executor, *LadonPolicy) error

	ladonPolicyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ladonPolicyType                 = reflect.TypeOf(&LadonPolicy{})
	ladonPolicyMapping              = queries.MakeStructMapping(ladonPolicyType)
	ladonPolicyPrimaryKeyMapping, _ = queries.BindMapping(ladonPolicyType, ladonPolicyMapping, ladonPolicyPrimaryKeyColumns)
	ladonPolicyInsertCacheMut       sync.RWMutex
	ladonPolicyInsertCache          = make(map[string]insertCache)
	ladonPolicyUpdateCacheMut       sync.RWMutex
	ladonPolicyUpdateCache          = make(map[string]updateCache)
	ladonPolicyUpsertCacheMut       sync.RWMutex
	ladonPolicyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var ladonPolicyBeforeInsertHooks []LadonPolicyHook
var ladonPolicyBeforeUpdateHooks []LadonPolicyHook
var ladonPolicyBeforeDeleteHooks []LadonPolicyHook
var ladonPolicyBeforeUpsertHooks []LadonPolicyHook

var ladonPolicyAfterInsertHooks []LadonPolicyHook
var ladonPolicyAfterSelectHooks []LadonPolicyHook
var ladonPolicyAfterUpdateHooks []LadonPolicyHook
var ladonPolicyAfterDeleteHooks []LadonPolicyHook
var ladonPolicyAfterUpsertHooks []LadonPolicyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LadonPolicy) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LadonPolicy) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LadonPolicy) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LadonPolicy) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LadonPolicy) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LadonPolicy) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LadonPolicy) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LadonPolicy) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LadonPolicy) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ladonPolicyAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLadonPolicyHook registers your hook function for all future operations.
func AddLadonPolicyHook(hookPoint boil.HookPoint, ladonPolicyHook LadonPolicyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ladonPolicyBeforeInsertHooks = append(ladonPolicyBeforeInsertHooks, ladonPolicyHook)
	case boil.BeforeUpdateHook:
		ladonPolicyBeforeUpdateHooks = append(ladonPolicyBeforeUpdateHooks, ladonPolicyHook)
	case boil.BeforeDeleteHook:
		ladonPolicyBeforeDeleteHooks = append(ladonPolicyBeforeDeleteHooks, ladonPolicyHook)
	case boil.BeforeUpsertHook:
		ladonPolicyBeforeUpsertHooks = append(ladonPolicyBeforeUpsertHooks, ladonPolicyHook)
	case boil.AfterInsertHook:
		ladonPolicyAfterInsertHooks = append(ladonPolicyAfterInsertHooks, ladonPolicyHook)
	case boil.AfterSelectHook:
		ladonPolicyAfterSelectHooks = append(ladonPolicyAfterSelectHooks, ladonPolicyHook)
	case boil.AfterUpdateHook:
		ladonPolicyAfterUpdateHooks = append(ladonPolicyAfterUpdateHooks, ladonPolicyHook)
	case boil.AfterDeleteHook:
		ladonPolicyAfterDeleteHooks = append(ladonPolicyAfterDeleteHooks, ladonPolicyHook)
	case boil.AfterUpsertHook:
		ladonPolicyAfterUpsertHooks = append(ladonPolicyAfterUpsertHooks, ladonPolicyHook)
	}
}

// OneP returns a single ladonPolicy record from the query, and panics on error.
func (q ladonPolicyQuery) OneP() *LadonPolicy {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single ladonPolicy record from the query.
func (q ladonPolicyQuery) One() (*LadonPolicy, error) {
	o := &LadonPolicy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ladon_policy")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all LadonPolicy records from the query, and panics on error.
func (q ladonPolicyQuery) AllP() LadonPolicySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all LadonPolicy records from the query.
func (q ladonPolicyQuery) All() (LadonPolicySlice, error) {
	var o LadonPolicySlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LadonPolicy slice")
	}

	if len(ladonPolicyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all LadonPolicy records in the query, and panics on error.
func (q ladonPolicyQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all LadonPolicy records in the query.
func (q ladonPolicyQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ladon_policy rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q ladonPolicyQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q ladonPolicyQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ladon_policy exists")
	}

	return count > 0, nil
}

// PolicyBoilerplateRolesG retrieves all the boilerplate_role's boilerplate role via policy_id column.
func (o *LadonPolicy) PolicyBoilerplateRolesG(mods ...qm.QueryMod) boilerplateRoleQuery {
	return o.PolicyBoilerplateRoles(boil.GetDB(), mods...)
}

// PolicyBoilerplateRoles retrieves all the boilerplate_role's boilerplate role with an executor via policy_id column.
func (o *LadonPolicy) PolicyBoilerplateRoles(exec boil.Executor, mods ...qm.QueryMod) boilerplateRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"policy_id\"=$1", o.ID),
	)

	query := BoilerplateRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"boilerplate\".\"boilerplate_role\" as \"a\"")
	return query
}

// LoadPolicyBoilerplateRoles allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (ladonPolicyL) LoadPolicyBoilerplateRoles(e boil.Executor, singular bool, maybeLadonPolicy interface{}) error {
	var slice []*LadonPolicy
	var object *LadonPolicy

	count := 1
	if singular {
		object = maybeLadonPolicy.(*LadonPolicy)
	} else {
		slice = *maybeLadonPolicy.(*LadonPolicySlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &ladonPolicyR{}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			obj.R = &ladonPolicyR{}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"boilerplate\".\"boilerplate_role\" where \"policy_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load boilerplate_role")
	}
	defer results.Close()

	var resultSlice []*BoilerplateRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice boilerplate_role")
	}

	if len(boilerplateRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PolicyBoilerplateRoles = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PolicyID {
				local.R.PolicyBoilerplateRoles = append(local.R.PolicyBoilerplateRoles, foreign)
				break
			}
		}
	}

	return nil
}

// AddPolicyBoilerplateRoles adds the given related objects to the existing relationships
// of the ladon_policy, optionally inserting them as new records.
// Appends related to o.R.PolicyBoilerplateRoles.
// Sets related.R.Policy appropriately.
func (o *LadonPolicy) AddPolicyBoilerplateRoles(exec boil.Executor, insert bool, related ...*BoilerplateRole) error {
	var err error
	for _, rel := range related {
		rel.PolicyID = o.ID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "policy_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &ladonPolicyR{
			PolicyBoilerplateRoles: related,
		}
	} else {
		o.R.PolicyBoilerplateRoles = append(o.R.PolicyBoilerplateRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &boilerplateRoleR{
				Policy: o,
			}
		} else {
			rel.R.Policy = o
		}
	}
	return nil
}

// LadonPoliciesG retrieves all records.
func LadonPoliciesG(mods ...qm.QueryMod) ladonPolicyQuery {
	return LadonPolicies(boil.GetDB(), mods...)
}

// LadonPolicies retrieves all the records using an executor.
func LadonPolicies(exec boil.Executor, mods ...qm.QueryMod) ladonPolicyQuery {
	mods = append(mods, qm.From("\"boilerplate\".\"ladon_policy\""))
	return ladonPolicyQuery{NewQuery(exec, mods...)}
}

// FindLadonPolicyG retrieves a single record by ID.
func FindLadonPolicyG(id string, selectCols ...string) (*LadonPolicy, error) {
	return FindLadonPolicy(boil.GetDB(), id, selectCols...)
}

// FindLadonPolicyGP retrieves a single record by ID, and panics on error.
func FindLadonPolicyGP(id string, selectCols ...string) *LadonPolicy {
	retobj, err := FindLadonPolicy(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindLadonPolicy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLadonPolicy(exec boil.Executor, id string, selectCols ...string) (*LadonPolicy, error) {
	ladonPolicyObj := &LadonPolicy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"boilerplate\".\"ladon_policy\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(ladonPolicyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ladon_policy")
	}

	return ladonPolicyObj, nil
}

// FindLadonPolicyP retrieves a single record by ID with an executor, and panics on error.
func FindLadonPolicyP(exec boil.Executor, id string, selectCols ...string) *LadonPolicy {
	retobj, err := FindLadonPolicy(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *LadonPolicy) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *LadonPolicy) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *LadonPolicy) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *LadonPolicy) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no ladon_policy provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ladonPolicyColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	ladonPolicyInsertCacheMut.RLock()
	cache, cached := ladonPolicyInsertCache[key]
	ladonPolicyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			ladonPolicyColumns,
			ladonPolicyColumnsWithDefault,
			ladonPolicyColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(ladonPolicyType, ladonPolicyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ladonPolicyType, ladonPolicyMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"boilerplate\".\"ladon_policy\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into ladon_policy")
	}

	if !cached {
		ladonPolicyInsertCacheMut.Lock()
		ladonPolicyInsertCache[key] = cache
		ladonPolicyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single LadonPolicy record. See Update for
// whitelist behavior description.
func (o *LadonPolicy) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single LadonPolicy record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *LadonPolicy) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the LadonPolicy, and panics on error.
// See Update for whitelist behavior description.
func (o *LadonPolicy) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the LadonPolicy.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *LadonPolicy) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	ladonPolicyUpdateCacheMut.RLock()
	cache, cached := ladonPolicyUpdateCache[key]
	ladonPolicyUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(ladonPolicyColumns, ladonPolicyPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update ladon_policy, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"boilerplate\".\"ladon_policy\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ladonPolicyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ladonPolicyType, ladonPolicyMapping, append(wl, ladonPolicyPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update ladon_policy row")
	}

	if !cached {
		ladonPolicyUpdateCacheMut.Lock()
		ladonPolicyUpdateCache[key] = cache
		ladonPolicyUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q ladonPolicyQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q ladonPolicyQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for ladon_policy")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o LadonPolicySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o LadonPolicySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o LadonPolicySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LadonPolicySlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ladonPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"boilerplate\".\"ladon_policy\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(ladonPolicyPrimaryKeyColumns), len(colNames)+1, len(ladonPolicyPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in ladonPolicy slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *LadonPolicy) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *LadonPolicy) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *LadonPolicy) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *LadonPolicy) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no ladon_policy provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ladonPolicyColumnsWithDefault, o)

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

	ladonPolicyUpsertCacheMut.RLock()
	cache, cached := ladonPolicyUpsertCache[key]
	ladonPolicyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			ladonPolicyColumns,
			ladonPolicyColumnsWithDefault,
			ladonPolicyColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			ladonPolicyColumns,
			ladonPolicyPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert ladon_policy, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ladonPolicyPrimaryKeyColumns))
			copy(conflict, ladonPolicyPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"boilerplate\".\"ladon_policy\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(ladonPolicyType, ladonPolicyMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ladonPolicyType, ladonPolicyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for ladon_policy")
	}

	if !cached {
		ladonPolicyUpsertCacheMut.Lock()
		ladonPolicyUpsertCache[key] = cache
		ladonPolicyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single LadonPolicy record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *LadonPolicy) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single LadonPolicy record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *LadonPolicy) DeleteG() error {
	if o == nil {
		return errors.New("models: no LadonPolicy provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single LadonPolicy record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *LadonPolicy) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single LadonPolicy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LadonPolicy) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no LadonPolicy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ladonPolicyPrimaryKeyMapping)
	sql := "DELETE FROM \"boilerplate\".\"ladon_policy\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from ladon_policy")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q ladonPolicyQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q ladonPolicyQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no ladonPolicyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ladon_policy")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o LadonPolicySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o LadonPolicySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no LadonPolicy slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o LadonPolicySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LadonPolicySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no LadonPolicy slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(ladonPolicyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ladonPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"boilerplate\".\"ladon_policy\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ladonPolicyPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(ladonPolicyPrimaryKeyColumns), 1, len(ladonPolicyPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ladonPolicy slice")
	}

	if len(ladonPolicyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *LadonPolicy) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *LadonPolicy) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *LadonPolicy) ReloadG() error {
	if o == nil {
		return errors.New("models: no LadonPolicy provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LadonPolicy) Reload(exec boil.Executor) error {
	ret, err := FindLadonPolicy(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LadonPolicySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LadonPolicySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LadonPolicySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty LadonPolicySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LadonPolicySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	ladonPolicies := LadonPolicySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ladonPolicyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"boilerplate\".\"ladon_policy\".* FROM \"boilerplate\".\"ladon_policy\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ladonPolicyPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(ladonPolicyPrimaryKeyColumns), 1, len(ladonPolicyPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&ladonPolicies)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LadonPolicySlice")
	}

	*o = ladonPolicies

	return nil
}

// LadonPolicyExists checks if the LadonPolicy row exists.
func LadonPolicyExists(exec boil.Executor, id string) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"boilerplate\".\"ladon_policy\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ladon_policy exists")
	}

	return exists, nil
}

// LadonPolicyExistsG checks if the LadonPolicy row exists.
func LadonPolicyExistsG(id string) (bool, error) {
	return LadonPolicyExists(boil.GetDB(), id)
}

// LadonPolicyExistsGP checks if the LadonPolicy row exists. Panics on error.
func LadonPolicyExistsGP(id string) bool {
	e, err := LadonPolicyExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// LadonPolicyExistsP checks if the LadonPolicy row exists. Panics on error.
func LadonPolicyExistsP(exec boil.Executor, id string) bool {
	e, err := LadonPolicyExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
