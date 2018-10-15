// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRoleMenuOptions(t *testing.T) {
	t.Parallel()

	query := RoleMenuOptions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRoleMenuOptionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleMenuOptionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RoleMenuOptions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleMenuOptionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleMenuOptionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleMenuOptionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoleMenuOptionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RoleMenuOption exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoleMenuOptionExists to return true, but got false.")
	}
}

func testRoleMenuOptionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	roleMenuOptionFound, err := FindRoleMenuOption(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if roleMenuOptionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRoleMenuOptionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RoleMenuOptions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRoleMenuOptionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RoleMenuOptions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRoleMenuOptionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	roleMenuOptionOne := &RoleMenuOption{}
	roleMenuOptionTwo := &RoleMenuOption{}
	if err = randomize.Struct(seed, roleMenuOptionOne, roleMenuOptionDBTypes, false, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}
	if err = randomize.Struct(seed, roleMenuOptionTwo, roleMenuOptionDBTypes, false, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleMenuOptionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleMenuOptionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoleMenuOptions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRoleMenuOptionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	roleMenuOptionOne := &RoleMenuOption{}
	roleMenuOptionTwo := &RoleMenuOption{}
	if err = randomize.Struct(seed, roleMenuOptionOne, roleMenuOptionDBTypes, false, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}
	if err = randomize.Struct(seed, roleMenuOptionTwo, roleMenuOptionDBTypes, false, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleMenuOptionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleMenuOptionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testRoleMenuOptionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleMenuOptionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(roleMenuOptionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleMenuOptionToOneRoleCommandUsingRoleCommand(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoleMenuOption
	var foreign RoleCommand

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roleCommandDBTypes, false, roleCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleCommand struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.RoleCommandID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.RoleCommand().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RoleMenuOptionSlice{&local}
	if err = local.L.LoadRoleCommand(ctx, tx, false, (*[]*RoleMenuOption)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RoleCommand == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RoleCommand = nil
	if err = local.L.LoadRoleCommand(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RoleCommand == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoleMenuOptionToOneRoleMenuUsingRoleMenu(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoleMenuOption
	var foreign RoleMenu

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roleMenuOptionDBTypes, false, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roleMenuDBTypes, false, roleMenuColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenu struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RoleMenuID = foreign.MessageID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.RoleMenu().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.MessageID != foreign.MessageID {
		t.Errorf("want: %v, got %v", foreign.MessageID, check.MessageID)
	}

	slice := RoleMenuOptionSlice{&local}
	if err = local.L.LoadRoleMenu(ctx, tx, false, (*[]*RoleMenuOption)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RoleMenu == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RoleMenu = nil
	if err = local.L.LoadRoleMenu(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RoleMenu == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoleMenuOptionToOneSetOpRoleCommandUsingRoleCommand(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleMenuOption
	var b, c RoleCommand

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleMenuOptionDBTypes, false, strmangle.SetComplement(roleMenuOptionPrimaryKeyColumns, roleMenuOptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roleCommandDBTypes, false, strmangle.SetComplement(roleCommandPrimaryKeyColumns, roleCommandColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleCommandDBTypes, false, strmangle.SetComplement(roleCommandPrimaryKeyColumns, roleCommandColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*RoleCommand{&b, &c} {
		err = a.SetRoleCommand(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RoleCommand != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleMenuOptions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.RoleCommandID, x.ID) {
			t.Error("foreign key was wrong value", a.RoleCommandID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RoleCommandID))
		reflect.Indirect(reflect.ValueOf(&a.RoleCommandID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.RoleCommandID, x.ID) {
			t.Error("foreign key was wrong value", a.RoleCommandID, x.ID)
		}
	}
}

func testRoleMenuOptionToOneRemoveOpRoleCommandUsingRoleCommand(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleMenuOption
	var b RoleCommand

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleMenuOptionDBTypes, false, strmangle.SetComplement(roleMenuOptionPrimaryKeyColumns, roleMenuOptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roleCommandDBTypes, false, strmangle.SetComplement(roleCommandPrimaryKeyColumns, roleCommandColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetRoleCommand(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveRoleCommand(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.RoleCommand().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.RoleCommand != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.RoleCommandID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.RoleMenuOptions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRoleMenuOptionToOneSetOpRoleMenuUsingRoleMenu(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleMenuOption
	var b, c RoleMenu

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleMenuOptionDBTypes, false, strmangle.SetComplement(roleMenuOptionPrimaryKeyColumns, roleMenuOptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roleMenuDBTypes, false, strmangle.SetComplement(roleMenuPrimaryKeyColumns, roleMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleMenuDBTypes, false, strmangle.SetComplement(roleMenuPrimaryKeyColumns, roleMenuColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*RoleMenu{&b, &c} {
		err = a.SetRoleMenu(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RoleMenu != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleMenuOptions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoleMenuID != x.MessageID {
			t.Error("foreign key was wrong value", a.RoleMenuID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RoleMenuID))
		reflect.Indirect(reflect.ValueOf(&a.RoleMenuID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RoleMenuID != x.MessageID {
			t.Error("foreign key was wrong value", a.RoleMenuID, x.MessageID)
		}
	}
}

func testRoleMenuOptionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoleMenuOptionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleMenuOptionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoleMenuOptionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoleMenuOptions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	roleMenuOptionDBTypes = map[string]string{`EmojiAnimated`: `boolean`, `EmojiID`: `bigint`, `ID`: `bigint`, `RoleCommandID`: `bigint`, `RoleMenuID`: `bigint`, `UnicodeEmoji`: `text`}
	_                     = bytes.MinRead
)

func testRoleMenuOptionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(roleMenuOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(roleMenuOptionColumns) == len(roleMenuOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRoleMenuOptionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(roleMenuOptionColumns) == len(roleMenuOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoleMenuOption{}
	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleMenuOptionDBTypes, true, roleMenuOptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(roleMenuOptionColumns, roleMenuOptionPrimaryKeyColumns) {
		fields = roleMenuOptionColumns
	} else {
		fields = strmangle.SetComplement(
			roleMenuOptionColumns,
			roleMenuOptionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RoleMenuOptionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRoleMenuOptionsUpsert(t *testing.T) {
	t.Parallel()

	if len(roleMenuOptionColumns) == len(roleMenuOptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RoleMenuOption{}
	if err = randomize.Struct(seed, &o, roleMenuOptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoleMenuOption: %s", err)
	}

	count, err := RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, roleMenuOptionDBTypes, false, roleMenuOptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleMenuOption struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoleMenuOption: %s", err)
	}

	count, err = RoleMenuOptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
