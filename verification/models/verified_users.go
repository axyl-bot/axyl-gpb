// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// VerifiedUser is an object representing the database table.
type VerifiedUser struct {
	GuildID    int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	UserID     int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	VerifiedAt time.Time `boil:"verified_at" json:"verified_at" toml:"verified_at" yaml:"verified_at"`
	IP         string    `boil:"ip" json:"ip" toml:"ip" yaml:"ip"`

	R *verifiedUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L verifiedUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VerifiedUserColumns = struct {
	GuildID    string
	UserID     string
	VerifiedAt string
	IP         string
}{
	GuildID:    "guild_id",
	UserID:     "user_id",
	VerifiedAt: "verified_at",
	IP:         "ip",
}

var VerifiedUserTableColumns = struct {
	GuildID    string
	UserID     string
	VerifiedAt string
	IP         string
}{
	GuildID:    "verified_users.guild_id",
	UserID:     "verified_users.user_id",
	VerifiedAt: "verified_users.verified_at",
	IP:         "verified_users.ip",
}

// Generated where

var VerifiedUserWhere = struct {
	GuildID    whereHelperint64
	UserID     whereHelperint64
	VerifiedAt whereHelpertime_Time
	IP         whereHelperstring
}{
	GuildID:    whereHelperint64{field: "\"verified_users\".\"guild_id\""},
	UserID:     whereHelperint64{field: "\"verified_users\".\"user_id\""},
	VerifiedAt: whereHelpertime_Time{field: "\"verified_users\".\"verified_at\""},
	IP:         whereHelperstring{field: "\"verified_users\".\"ip\""},
}

// VerifiedUserRels is where relationship names are stored.
var VerifiedUserRels = struct {
}{}

// verifiedUserR is where relationships are stored.
type verifiedUserR struct {
}

// NewStruct creates a new relationship struct
func (*verifiedUserR) NewStruct() *verifiedUserR {
	return &verifiedUserR{}
}

// verifiedUserL is where Load methods for each relationship are stored.
type verifiedUserL struct{}

var (
	verifiedUserAllColumns            = []string{"guild_id", "user_id", "verified_at", "ip"}
	verifiedUserColumnsWithoutDefault = []string{"guild_id", "user_id", "verified_at", "ip"}
	verifiedUserColumnsWithDefault    = []string{}
	verifiedUserPrimaryKeyColumns     = []string{"guild_id", "user_id"}
	verifiedUserGeneratedColumns      = []string{}
)

type (
	// VerifiedUserSlice is an alias for a slice of pointers to VerifiedUser.
	// This should almost always be used instead of []VerifiedUser.
	VerifiedUserSlice []*VerifiedUser

	verifiedUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	verifiedUserType                 = reflect.TypeOf(&VerifiedUser{})
	verifiedUserMapping              = queries.MakeStructMapping(verifiedUserType)
	verifiedUserPrimaryKeyMapping, _ = queries.BindMapping(verifiedUserType, verifiedUserMapping, verifiedUserPrimaryKeyColumns)
	verifiedUserInsertCacheMut       sync.RWMutex
	verifiedUserInsertCache          = make(map[string]insertCache)
	verifiedUserUpdateCacheMut       sync.RWMutex
	verifiedUserUpdateCache          = make(map[string]updateCache)
	verifiedUserUpsertCacheMut       sync.RWMutex
	verifiedUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single verifiedUser record from the query using the global executor.
func (q verifiedUserQuery) OneG(ctx context.Context) (*VerifiedUser, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single verifiedUser record from the query.
func (q verifiedUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VerifiedUser, error) {
	o := &VerifiedUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for verified_users")
	}

	return o, nil
}

// AllG returns all VerifiedUser records from the query using the global executor.
func (q verifiedUserQuery) AllG(ctx context.Context) (VerifiedUserSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all VerifiedUser records from the query.
func (q verifiedUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (VerifiedUserSlice, error) {
	var o []*VerifiedUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VerifiedUser slice")
	}

	return o, nil
}

// CountG returns the count of all VerifiedUser records in the query using the global executor
func (q verifiedUserQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all VerifiedUser records in the query.
func (q verifiedUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count verified_users rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q verifiedUserQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q verifiedUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if verified_users exists")
	}

	return count > 0, nil
}

// VerifiedUsers retrieves all the records using an executor.
func VerifiedUsers(mods ...qm.QueryMod) verifiedUserQuery {
	mods = append(mods, qm.From("\"verified_users\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"verified_users\".*"})
	}

	return verifiedUserQuery{q}
}

// FindVerifiedUserG retrieves a single record by ID.
func FindVerifiedUserG(ctx context.Context, guildID int64, userID int64, selectCols ...string) (*VerifiedUser, error) {
	return FindVerifiedUser(ctx, boil.GetContextDB(), guildID, userID, selectCols...)
}

// FindVerifiedUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVerifiedUser(ctx context.Context, exec boil.ContextExecutor, guildID int64, userID int64, selectCols ...string) (*VerifiedUser, error) {
	verifiedUserObj := &VerifiedUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"verified_users\" where \"guild_id\"=$1 AND \"user_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, userID)

	err := q.Bind(ctx, exec, verifiedUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from verified_users")
	}

	return verifiedUserObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VerifiedUser) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VerifiedUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no verified_users provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(verifiedUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	verifiedUserInsertCacheMut.RLock()
	cache, cached := verifiedUserInsertCache[key]
	verifiedUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			verifiedUserAllColumns,
			verifiedUserColumnsWithDefault,
			verifiedUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(verifiedUserType, verifiedUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(verifiedUserType, verifiedUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"verified_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"verified_users\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into verified_users")
	}

	if !cached {
		verifiedUserInsertCacheMut.Lock()
		verifiedUserInsertCache[key] = cache
		verifiedUserInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single VerifiedUser record using the global executor.
// See Update for more documentation.
func (o *VerifiedUser) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the VerifiedUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VerifiedUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	verifiedUserUpdateCacheMut.RLock()
	cache, cached := verifiedUserUpdateCache[key]
	verifiedUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			verifiedUserAllColumns,
			verifiedUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update verified_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"verified_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, verifiedUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(verifiedUserType, verifiedUserMapping, append(wl, verifiedUserPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update verified_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for verified_users")
	}

	if !cached {
		verifiedUserUpdateCacheMut.Lock()
		verifiedUserUpdateCache[key] = cache
		verifiedUserUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q verifiedUserQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q verifiedUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for verified_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for verified_users")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VerifiedUserSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VerifiedUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verifiedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"verified_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, verifiedUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in verifiedUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all verifiedUser")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *VerifiedUser) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VerifiedUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no verified_users provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(verifiedUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	verifiedUserUpsertCacheMut.RLock()
	cache, cached := verifiedUserUpsertCache[key]
	verifiedUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			verifiedUserAllColumns,
			verifiedUserColumnsWithDefault,
			verifiedUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			verifiedUserAllColumns,
			verifiedUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert verified_users, could not build update column list")
		}

		ret := strmangle.SetComplement(verifiedUserAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(verifiedUserPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert verified_users, could not build conflict column list")
			}

			conflict = make([]string, len(verifiedUserPrimaryKeyColumns))
			copy(conflict, verifiedUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"verified_users\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(verifiedUserType, verifiedUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(verifiedUserType, verifiedUserMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert verified_users")
	}

	if !cached {
		verifiedUserUpsertCacheMut.Lock()
		verifiedUserUpsertCache[key] = cache
		verifiedUserUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single VerifiedUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *VerifiedUser) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single VerifiedUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VerifiedUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no VerifiedUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), verifiedUserPrimaryKeyMapping)
	sql := "DELETE FROM \"verified_users\" WHERE \"guild_id\"=$1 AND \"user_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from verified_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for verified_users")
	}

	return rowsAff, nil
}

func (q verifiedUserQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q verifiedUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no verifiedUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from verified_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for verified_users")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o VerifiedUserSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VerifiedUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verifiedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"verified_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, verifiedUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from verifiedUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for verified_users")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *VerifiedUser) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no VerifiedUser provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VerifiedUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVerifiedUser(ctx, exec, o.GuildID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VerifiedUserSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty VerifiedUserSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VerifiedUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VerifiedUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verifiedUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"verified_users\".* FROM \"verified_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, verifiedUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VerifiedUserSlice")
	}

	*o = slice

	return nil
}

// VerifiedUserExistsG checks if the VerifiedUser row exists.
func VerifiedUserExistsG(ctx context.Context, guildID int64, userID int64) (bool, error) {
	return VerifiedUserExists(ctx, boil.GetContextDB(), guildID, userID)
}

// VerifiedUserExists checks if the VerifiedUser row exists.
func VerifiedUserExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, userID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"verified_users\" where \"guild_id\"=$1 AND \"user_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, guildID, userID)
	}
	row := exec.QueryRowContext(ctx, sql, guildID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if verified_users exists")
	}

	return exists, nil
}

// Exists checks if the VerifiedUser row exists.
func (o *VerifiedUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return VerifiedUserExists(ctx, exec, o.GuildID, o.UserID)
}
