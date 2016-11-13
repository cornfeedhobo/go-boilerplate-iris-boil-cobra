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
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testBoilerplateUsers(t *testing.T) {
	t.Parallel()

	query := BoilerplateUsers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBoilerplateUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = boilerplateUser.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoilerplateUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BoilerplateUsers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoilerplateUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BoilerplateUserSlice{boilerplateUser}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBoilerplateUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BoilerplateUserExists(tx, boilerplateUser.ID)
	if err != nil {
		t.Errorf("Unable to check if BoilerplateUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BoilerplateUserExistsG to return true, but got false.")
	}
}
func testBoilerplateUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	boilerplateUserFound, err := FindBoilerplateUser(tx, boilerplateUser.ID)
	if err != nil {
		t.Error(err)
	}

	if boilerplateUserFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBoilerplateUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BoilerplateUsers(tx).Bind(boilerplateUser); err != nil {
		t.Error(err)
	}
}

func testBoilerplateUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := BoilerplateUsers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBoilerplateUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUserOne := &BoilerplateUser{}
	boilerplateUserTwo := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUserOne, boilerplateUserDBTypes, false, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}
	if err = randomize.Struct(seed, boilerplateUserTwo, boilerplateUserDBTypes, false, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = boilerplateUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BoilerplateUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBoilerplateUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	boilerplateUserOne := &BoilerplateUser{}
	boilerplateUserTwo := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUserOne, boilerplateUserDBTypes, false, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}
	if err = randomize.Struct(seed, boilerplateUserTwo, boilerplateUserDBTypes, false, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = boilerplateUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func boilerplateUserBeforeInsertHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserAfterInsertHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserAfterSelectHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserBeforeUpdateHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserAfterUpdateHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserBeforeDeleteHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserAfterDeleteHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserBeforeUpsertHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func boilerplateUserAfterUpsertHook(e boil.Executor, o *BoilerplateUser) error {
	*o = BoilerplateUser{}
	return nil
}

func testBoilerplateUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &BoilerplateUser{}
	o := &BoilerplateUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, boilerplateUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser object: %s", err)
	}

	AddBoilerplateUserHook(boil.BeforeInsertHook, boilerplateUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	boilerplateUserBeforeInsertHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.AfterInsertHook, boilerplateUserAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	boilerplateUserAfterInsertHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.AfterSelectHook, boilerplateUserAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	boilerplateUserAfterSelectHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.BeforeUpdateHook, boilerplateUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	boilerplateUserBeforeUpdateHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.AfterUpdateHook, boilerplateUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	boilerplateUserAfterUpdateHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.BeforeDeleteHook, boilerplateUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	boilerplateUserBeforeDeleteHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.AfterDeleteHook, boilerplateUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	boilerplateUserAfterDeleteHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.BeforeUpsertHook, boilerplateUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	boilerplateUserBeforeUpsertHooks = []BoilerplateUserHook{}

	AddBoilerplateUserHook(boil.AfterUpsertHook, boilerplateUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	boilerplateUserAfterUpsertHooks = []BoilerplateUserHook{}
}
func testBoilerplateUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBoilerplateUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx, boilerplateUserColumns...); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBoilerplateUserToOneBoilerplateRoleUsingRole(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local BoilerplateUser
	var foreign BoilerplateRole

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.RoleID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Role(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BoilerplateUserSlice{&local}
	if err = local.L.LoadRole(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Role = nil
	if err = local.L.LoadRole(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBoilerplateUserToOneSetOpBoilerplateRoleUsingRole(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a BoilerplateUser
	var b, c BoilerplateRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boilerplateUserDBTypes, false, strmangle.SetComplement(boilerplateUserPrimaryKeyColumns, boilerplateUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, boilerplateRoleDBTypes, false, strmangle.SetComplement(boilerplateRolePrimaryKeyColumns, boilerplateRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, boilerplateRoleDBTypes, false, strmangle.SetComplement(boilerplateRolePrimaryKeyColumns, boilerplateRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BoilerplateRole{&b, &c} {
		err = a.SetRole(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Role != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleBoilerplateUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RoleID))
		reflect.Indirect(reflect.ValueOf(&a.RoleID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID, x.ID)
		}
	}
}
func testBoilerplateUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = boilerplateUser.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBoilerplateUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BoilerplateUserSlice{boilerplateUser}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBoilerplateUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BoilerplateUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	boilerplateUserDBTypes = map[string]string{"CreatedAt": "timestamp with time zone", "ID": "integer", "Password": "text", "RoleID": "integer", "UpdatedAt": "timestamp with time zone", "Username": "text"}
	_                      = bytes.MinRead
)

func testBoilerplateUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(boilerplateUserColumns) == len(boilerplateUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	if err = boilerplateUser.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBoilerplateUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(boilerplateUserColumns) == len(boilerplateUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	boilerplateUser := &BoilerplateUser{}
	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, boilerplateUser, boilerplateUserDBTypes, true, boilerplateUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(boilerplateUserColumns, boilerplateUserPrimaryKeyColumns) {
		fields = boilerplateUserColumns
	} else {
		fields = strmangle.SetComplement(
			boilerplateUserColumns,
			boilerplateUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(boilerplateUser))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BoilerplateUserSlice{boilerplateUser}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBoilerplateUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(boilerplateUserColumns) == len(boilerplateUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	boilerplateUser := BoilerplateUser{}
	if err = randomize.Struct(seed, &boilerplateUser, boilerplateUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateUser.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert BoilerplateUser: %s", err)
	}

	count, err := BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &boilerplateUser, boilerplateUserDBTypes, false, boilerplateUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BoilerplateUser struct: %s", err)
	}

	if err = boilerplateUser.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert BoilerplateUser: %s", err)
	}

	count, err = BoilerplateUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
