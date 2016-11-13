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

func testBoilerplateRoles(t *testing.T) {
	t.Parallel()

	query := BoilerplateRoles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBoilerplateRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = boilerplateRole.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoilerplateRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BoilerplateRoles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBoilerplateRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BoilerplateRoleSlice{boilerplateRole}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBoilerplateRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BoilerplateRoleExists(tx, boilerplateRole.ID)
	if err != nil {
		t.Errorf("Unable to check if BoilerplateRole exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BoilerplateRoleExistsG to return true, but got false.")
	}
}
func testBoilerplateRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	boilerplateRoleFound, err := FindBoilerplateRole(tx, boilerplateRole.ID)
	if err != nil {
		t.Error(err)
	}

	if boilerplateRoleFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBoilerplateRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BoilerplateRoles(tx).Bind(boilerplateRole); err != nil {
		t.Error(err)
	}
}

func testBoilerplateRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := BoilerplateRoles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBoilerplateRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRoleOne := &BoilerplateRole{}
	boilerplateRoleTwo := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRoleOne, boilerplateRoleDBTypes, false, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}
	if err = randomize.Struct(seed, boilerplateRoleTwo, boilerplateRoleDBTypes, false, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = boilerplateRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BoilerplateRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBoilerplateRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	boilerplateRoleOne := &BoilerplateRole{}
	boilerplateRoleTwo := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRoleOne, boilerplateRoleDBTypes, false, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}
	if err = randomize.Struct(seed, boilerplateRoleTwo, boilerplateRoleDBTypes, false, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRoleOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = boilerplateRoleTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func boilerplateRoleBeforeInsertHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleAfterInsertHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleAfterSelectHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleBeforeUpdateHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleAfterUpdateHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleBeforeDeleteHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleAfterDeleteHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleBeforeUpsertHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func boilerplateRoleAfterUpsertHook(e boil.Executor, o *BoilerplateRole) error {
	*o = BoilerplateRole{}
	return nil
}

func testBoilerplateRolesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &BoilerplateRole{}
	o := &BoilerplateRole{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, boilerplateRoleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole object: %s", err)
	}

	AddBoilerplateRoleHook(boil.BeforeInsertHook, boilerplateRoleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleBeforeInsertHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.AfterInsertHook, boilerplateRoleAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleAfterInsertHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.AfterSelectHook, boilerplateRoleAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleAfterSelectHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.BeforeUpdateHook, boilerplateRoleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleBeforeUpdateHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.AfterUpdateHook, boilerplateRoleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleAfterUpdateHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.BeforeDeleteHook, boilerplateRoleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleBeforeDeleteHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.AfterDeleteHook, boilerplateRoleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleAfterDeleteHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.BeforeUpsertHook, boilerplateRoleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleBeforeUpsertHooks = []BoilerplateRoleHook{}

	AddBoilerplateRoleHook(boil.AfterUpsertHook, boilerplateRoleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	boilerplateRoleAfterUpsertHooks = []BoilerplateRoleHook{}
}
func testBoilerplateRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBoilerplateRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx, boilerplateRoleColumns...); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBoilerplateRoleToManyRoleBoilerplateUsers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a BoilerplateRole
	var b, c BoilerplateUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, boilerplateUserDBTypes, false, boilerplateUserColumnsWithDefault...)
	randomize.Struct(seed, &c, boilerplateUserDBTypes, false, boilerplateUserColumnsWithDefault...)

	b.RoleID = a.ID
	c.RoleID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	boilerplateUser, err := a.RoleBoilerplateUsers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range boilerplateUser {
		if v.RoleID == b.RoleID {
			bFound = true
		}
		if v.RoleID == c.RoleID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BoilerplateRoleSlice{&a}
	if err = a.L.LoadRoleBoilerplateUsers(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleBoilerplateUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RoleBoilerplateUsers = nil
	if err = a.L.LoadRoleBoilerplateUsers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleBoilerplateUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", boilerplateUser)
	}
}

func testBoilerplateRoleToManyAddOpRoleBoilerplateUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a BoilerplateRole
	var b, c, d, e BoilerplateUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boilerplateRoleDBTypes, false, strmangle.SetComplement(boilerplateRolePrimaryKeyColumns, boilerplateRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BoilerplateUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, boilerplateUserDBTypes, false, strmangle.SetComplement(boilerplateUserPrimaryKeyColumns, boilerplateUserColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BoilerplateUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRoleBoilerplateUsers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RoleID {
			t.Error("foreign key was wrong value", a.ID, first.RoleID)
		}
		if a.ID != second.RoleID {
			t.Error("foreign key was wrong value", a.ID, second.RoleID)
		}

		if first.R.Role != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Role != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RoleBoilerplateUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RoleBoilerplateUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RoleBoilerplateUsers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBoilerplateRoleToOneLadonPolicyUsingPolicy(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local BoilerplateRole
	var foreign LadonPolicy

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PolicyID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Policy(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BoilerplateRoleSlice{&local}
	if err = local.L.LoadPolicy(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Policy == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Policy = nil
	if err = local.L.LoadPolicy(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Policy == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBoilerplateRoleToOneSetOpLadonPolicyUsingPolicy(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a BoilerplateRole
	var b, c LadonPolicy

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, boilerplateRoleDBTypes, false, strmangle.SetComplement(boilerplateRolePrimaryKeyColumns, boilerplateRoleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ladonPolicyDBTypes, false, strmangle.SetComplement(ladonPolicyPrimaryKeyColumns, ladonPolicyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, ladonPolicyDBTypes, false, strmangle.SetComplement(ladonPolicyPrimaryKeyColumns, ladonPolicyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*LadonPolicy{&b, &c} {
		err = a.SetPolicy(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Policy != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PolicyBoilerplateRoles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PolicyID != x.ID {
			t.Error("foreign key was wrong value", a.PolicyID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PolicyID))
		reflect.Indirect(reflect.ValueOf(&a.PolicyID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PolicyID != x.ID {
			t.Error("foreign key was wrong value", a.PolicyID, x.ID)
		}
	}
}
func testBoilerplateRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = boilerplateRole.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBoilerplateRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BoilerplateRoleSlice{boilerplateRole}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBoilerplateRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BoilerplateRoles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	boilerplateRoleDBTypes = map[string]string{"CreatedAt": "timestamp with time zone", "ID": "integer", "PolicyID": "character varying", "UpdatedAt": "timestamp with time zone"}
	_                      = bytes.MinRead
)

func testBoilerplateRolesUpdate(t *testing.T) {
	t.Parallel()

	if len(boilerplateRoleColumns) == len(boilerplateRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRoleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	if err = boilerplateRole.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBoilerplateRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(boilerplateRoleColumns) == len(boilerplateRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	boilerplateRole := &BoilerplateRole{}
	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, boilerplateRole, boilerplateRoleDBTypes, true, boilerplateRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(boilerplateRoleColumns, boilerplateRolePrimaryKeyColumns) {
		fields = boilerplateRoleColumns
	} else {
		fields = strmangle.SetComplement(
			boilerplateRoleColumns,
			boilerplateRolePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(boilerplateRole))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BoilerplateRoleSlice{boilerplateRole}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBoilerplateRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(boilerplateRoleColumns) == len(boilerplateRolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	boilerplateRole := BoilerplateRole{}
	if err = randomize.Struct(seed, &boilerplateRole, boilerplateRoleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = boilerplateRole.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert BoilerplateRole: %s", err)
	}

	count, err := BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &boilerplateRole, boilerplateRoleDBTypes, false, boilerplateRolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BoilerplateRole struct: %s", err)
	}

	if err = boilerplateRole.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert BoilerplateRole: %s", err)
	}

	count, err = BoilerplateRoles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
