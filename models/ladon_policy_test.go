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

func testLadonPolicies(t *testing.T) {
	t.Parallel()

	query := LadonPolicies(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testLadonPoliciesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ladonPolicy.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLadonPoliciesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = LadonPolicies(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLadonPoliciesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LadonPolicySlice{ladonPolicy}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testLadonPoliciesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := LadonPolicyExists(tx, ladonPolicy.ID)
	if err != nil {
		t.Errorf("Unable to check if LadonPolicy exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LadonPolicyExistsG to return true, but got false.")
	}
}
func testLadonPoliciesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	ladonPolicyFound, err := FindLadonPolicy(tx, ladonPolicy.ID)
	if err != nil {
		t.Error(err)
	}

	if ladonPolicyFound == nil {
		t.Error("want a record, got nil")
	}
}
func testLadonPoliciesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = LadonPolicies(tx).Bind(ladonPolicy); err != nil {
		t.Error(err)
	}
}

func testLadonPoliciesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := LadonPolicies(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLadonPoliciesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicyOne := &LadonPolicy{}
	ladonPolicyTwo := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicyOne, ladonPolicyDBTypes, false, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}
	if err = randomize.Struct(seed, ladonPolicyTwo, ladonPolicyDBTypes, false, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicyOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = ladonPolicyTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := LadonPolicies(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLadonPoliciesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ladonPolicyOne := &LadonPolicy{}
	ladonPolicyTwo := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicyOne, ladonPolicyDBTypes, false, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}
	if err = randomize.Struct(seed, ladonPolicyTwo, ladonPolicyDBTypes, false, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicyOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = ladonPolicyTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func ladonPolicyBeforeInsertHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyAfterInsertHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyAfterSelectHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyBeforeUpdateHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyAfterUpdateHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyBeforeDeleteHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyAfterDeleteHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyBeforeUpsertHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func ladonPolicyAfterUpsertHook(e boil.Executor, o *LadonPolicy) error {
	*o = LadonPolicy{}
	return nil
}

func testLadonPoliciesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &LadonPolicy{}
	o := &LadonPolicy{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ladonPolicyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LadonPolicy object: %s", err)
	}

	AddLadonPolicyHook(boil.BeforeInsertHook, ladonPolicyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ladonPolicyBeforeInsertHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.AfterInsertHook, ladonPolicyAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ladonPolicyAfterInsertHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.AfterSelectHook, ladonPolicyAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ladonPolicyAfterSelectHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.BeforeUpdateHook, ladonPolicyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ladonPolicyBeforeUpdateHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.AfterUpdateHook, ladonPolicyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ladonPolicyAfterUpdateHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.BeforeDeleteHook, ladonPolicyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ladonPolicyBeforeDeleteHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.AfterDeleteHook, ladonPolicyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ladonPolicyAfterDeleteHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.BeforeUpsertHook, ladonPolicyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ladonPolicyBeforeUpsertHooks = []LadonPolicyHook{}

	AddLadonPolicyHook(boil.AfterUpsertHook, ladonPolicyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ladonPolicyAfterUpsertHooks = []LadonPolicyHook{}
}
func testLadonPoliciesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLadonPoliciesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx, ladonPolicyColumns...); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLadonPolicyToManyPolicyBoilerplateRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a LadonPolicy
	var b, c BoilerplateRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, boilerplateRoleDBTypes, false, boilerplateRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, boilerplateRoleDBTypes, false, boilerplateRoleColumnsWithDefault...)

	b.PolicyID = a.ID
	c.PolicyID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	boilerplateRole, err := a.PolicyBoilerplateRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range boilerplateRole {
		if v.PolicyID == b.PolicyID {
			bFound = true
		}
		if v.PolicyID == c.PolicyID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LadonPolicySlice{&a}
	if err = a.L.LoadPolicyBoilerplateRoles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PolicyBoilerplateRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PolicyBoilerplateRoles = nil
	if err = a.L.LoadPolicyBoilerplateRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PolicyBoilerplateRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", boilerplateRole)
	}
}

func testLadonPolicyToManyAddOpPolicyBoilerplateRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a LadonPolicy
	var b, c, d, e BoilerplateRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ladonPolicyDBTypes, false, strmangle.SetComplement(ladonPolicyPrimaryKeyColumns, ladonPolicyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BoilerplateRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, boilerplateRoleDBTypes, false, strmangle.SetComplement(boilerplateRolePrimaryKeyColumns, boilerplateRoleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BoilerplateRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPolicyBoilerplateRoles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PolicyID {
			t.Error("foreign key was wrong value", a.ID, first.PolicyID)
		}
		if a.ID != second.PolicyID {
			t.Error("foreign key was wrong value", a.ID, second.PolicyID)
		}

		if first.R.Policy != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Policy != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PolicyBoilerplateRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PolicyBoilerplateRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PolicyBoilerplateRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testLadonPoliciesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ladonPolicy.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testLadonPoliciesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LadonPolicySlice{ladonPolicy}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testLadonPoliciesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := LadonPolicies(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ladonPolicyDBTypes = map[string]string{"Conditions": "text", "Description": "text", "Effect": "text", "ID": "character varying"}
	_                  = bytes.MinRead
)

func testLadonPoliciesUpdate(t *testing.T) {
	t.Parallel()

	if len(ladonPolicyColumns) == len(ladonPolicyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	if err = ladonPolicy.Update(tx); err != nil {
		t.Error(err)
	}
}

func testLadonPoliciesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ladonPolicyColumns) == len(ladonPolicyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ladonPolicy := &LadonPolicy{}
	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ladonPolicy, ladonPolicyDBTypes, true, ladonPolicyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ladonPolicyColumns, ladonPolicyPrimaryKeyColumns) {
		fields = ladonPolicyColumns
	} else {
		fields = strmangle.SetComplement(
			ladonPolicyColumns,
			ladonPolicyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(ladonPolicy))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := LadonPolicySlice{ladonPolicy}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testLadonPoliciesUpsert(t *testing.T) {
	t.Parallel()

	if len(ladonPolicyColumns) == len(ladonPolicyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	ladonPolicy := LadonPolicy{}
	if err = randomize.Struct(seed, &ladonPolicy, ladonPolicyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ladonPolicy.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert LadonPolicy: %s", err)
	}

	count, err := LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &ladonPolicy, ladonPolicyDBTypes, false, ladonPolicyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LadonPolicy struct: %s", err)
	}

	if err = ladonPolicy.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert LadonPolicy: %s", err)
	}

	count, err = LadonPolicies(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
