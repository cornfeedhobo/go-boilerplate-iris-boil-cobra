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

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("LadonPolicies", testLadonPolicies)
	t.Run("BoilerplateRoles", testBoilerplateRoles)
	t.Run("BoilerplateUsers", testBoilerplateUsers)
}

func TestDelete(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesDelete)
	t.Run("BoilerplateRoles", testBoilerplateRolesDelete)
	t.Run("BoilerplateUsers", testBoilerplateUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesQueryDeleteAll)
	t.Run("BoilerplateRoles", testBoilerplateRolesQueryDeleteAll)
	t.Run("BoilerplateUsers", testBoilerplateUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesSliceDeleteAll)
	t.Run("BoilerplateRoles", testBoilerplateRolesSliceDeleteAll)
	t.Run("BoilerplateUsers", testBoilerplateUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesExists)
	t.Run("BoilerplateRoles", testBoilerplateRolesExists)
	t.Run("BoilerplateUsers", testBoilerplateUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesFind)
	t.Run("BoilerplateRoles", testBoilerplateRolesFind)
	t.Run("BoilerplateUsers", testBoilerplateUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesBind)
	t.Run("BoilerplateRoles", testBoilerplateRolesBind)
	t.Run("BoilerplateUsers", testBoilerplateUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesOne)
	t.Run("BoilerplateRoles", testBoilerplateRolesOne)
	t.Run("BoilerplateUsers", testBoilerplateUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesAll)
	t.Run("BoilerplateRoles", testBoilerplateRolesAll)
	t.Run("BoilerplateUsers", testBoilerplateUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesCount)
	t.Run("BoilerplateRoles", testBoilerplateRolesCount)
	t.Run("BoilerplateUsers", testBoilerplateUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesHooks)
	t.Run("BoilerplateRoles", testBoilerplateRolesHooks)
	t.Run("BoilerplateUsers", testBoilerplateUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesInsert)
	t.Run("LadonPolicies", testLadonPoliciesInsertWhitelist)
	t.Run("BoilerplateRoles", testBoilerplateRolesInsert)
	t.Run("BoilerplateRoles", testBoilerplateRolesInsertWhitelist)
	t.Run("BoilerplateUsers", testBoilerplateUsersInsert)
	t.Run("BoilerplateUsers", testBoilerplateUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BoilerplateRoleToLadonPolicyUsingPolicy", testBoilerplateRoleToOneLadonPolicyUsingPolicy)
	t.Run("BoilerplateUserToBoilerplateRoleUsingRole", testBoilerplateUserToOneBoilerplateRoleUsingRole)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("LadonPolicyToPolicyBoilerplateRoles", testLadonPolicyToManyPolicyBoilerplateRoles)
	t.Run("BoilerplateRoleToRoleBoilerplateUsers", testBoilerplateRoleToManyRoleBoilerplateUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BoilerplateRoleToLadonPolicyUsingPolicy", testBoilerplateRoleToOneSetOpLadonPolicyUsingPolicy)
	t.Run("BoilerplateUserToBoilerplateRoleUsingRole", testBoilerplateUserToOneSetOpBoilerplateRoleUsingRole)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("LadonPolicyToPolicyBoilerplateRoles", testLadonPolicyToManyAddOpPolicyBoilerplateRoles)
	t.Run("BoilerplateRoleToRoleBoilerplateUsers", testBoilerplateRoleToManyAddOpRoleBoilerplateUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesReload)
	t.Run("BoilerplateRoles", testBoilerplateRolesReload)
	t.Run("BoilerplateUsers", testBoilerplateUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesReloadAll)
	t.Run("BoilerplateRoles", testBoilerplateRolesReloadAll)
	t.Run("BoilerplateUsers", testBoilerplateUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesSelect)
	t.Run("BoilerplateRoles", testBoilerplateRolesSelect)
	t.Run("BoilerplateUsers", testBoilerplateUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesUpdate)
	t.Run("BoilerplateRoles", testBoilerplateRolesUpdate)
	t.Run("BoilerplateUsers", testBoilerplateUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesSliceUpdateAll)
	t.Run("BoilerplateRoles", testBoilerplateRolesSliceUpdateAll)
	t.Run("BoilerplateUsers", testBoilerplateUsersSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("LadonPolicies", testLadonPoliciesUpsert)
	t.Run("BoilerplateRoles", testBoilerplateRolesUpsert)
	t.Run("BoilerplateUsers", testBoilerplateUsersUpsert)
}
