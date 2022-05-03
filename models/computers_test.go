// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

func testComputersUpsert(t *testing.T) {
	t.Parallel()
	if len(computerAllColumns) == len(computerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Computer{}
	if err = randomize.Struct(seed, &o, computerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Computer: %s", err)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, computerDBTypes, false, computerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Computer: %s", err)
	}

	count, err = Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testComputers(t *testing.T) {
	t.Parallel()

	query := Computers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testComputersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
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

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComputersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Computers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComputersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ComputerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComputersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ComputerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Computer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ComputerExists to return true, but got false.")
	}
}

func testComputersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	computerFound, err := FindComputer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if computerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testComputersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Computers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testComputersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Computers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testComputersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	computerOne := &Computer{}
	computerTwo := &Computer{}
	if err = randomize.Struct(seed, computerOne, computerDBTypes, false, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}
	if err = randomize.Struct(seed, computerTwo, computerDBTypes, false, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = computerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = computerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Computers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testComputersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	computerOne := &Computer{}
	computerTwo := &Computer{}
	if err = randomize.Struct(seed, computerOne, computerDBTypes, false, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}
	if err = randomize.Struct(seed, computerTwo, computerDBTypes, false, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = computerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = computerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func computerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func computerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Computer) error {
	*o = Computer{}
	return nil
}

func testComputersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Computer{}
	o := &Computer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, computerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Computer object: %s", err)
	}

	AddComputerHook(boil.BeforeInsertHook, computerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	computerBeforeInsertHooks = []ComputerHook{}

	AddComputerHook(boil.AfterInsertHook, computerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	computerAfterInsertHooks = []ComputerHook{}

	AddComputerHook(boil.AfterSelectHook, computerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	computerAfterSelectHooks = []ComputerHook{}

	AddComputerHook(boil.BeforeUpdateHook, computerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	computerBeforeUpdateHooks = []ComputerHook{}

	AddComputerHook(boil.AfterUpdateHook, computerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	computerAfterUpdateHooks = []ComputerHook{}

	AddComputerHook(boil.BeforeDeleteHook, computerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	computerBeforeDeleteHooks = []ComputerHook{}

	AddComputerHook(boil.AfterDeleteHook, computerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	computerAfterDeleteHooks = []ComputerHook{}

	AddComputerHook(boil.BeforeUpsertHook, computerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	computerBeforeUpsertHooks = []ComputerHook{}

	AddComputerHook(boil.AfterUpsertHook, computerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	computerAfterUpsertHooks = []ComputerHook{}
}

func testComputersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testComputersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(computerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testComputersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
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

func testComputersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ComputerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testComputersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Computers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	computerDBTypes = map[string]string{`ID`: `INTEGER`, `Name`: `TEXT`, `SSHUser`: `TEXT`, `SSHKey`: `TEXT`, `SSHPort`: `INTEGER`, `IPAddress`: `TEXT`, `MacAddress`: `TEXT`}
	_               = bytes.MinRead
)

func testComputersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(computerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(computerAllColumns) == len(computerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, computerDBTypes, true, computerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testComputersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(computerAllColumns) == len(computerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Computer{}
	if err = randomize.Struct(seed, o, computerDBTypes, true, computerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Computers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, computerDBTypes, true, computerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Computer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(computerAllColumns, computerPrimaryKeyColumns) {
		fields = computerAllColumns
	} else {
		fields = strmangle.SetComplement(
			computerAllColumns,
			computerPrimaryKeyColumns,
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

	slice := ComputerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
