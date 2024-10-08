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

// ReputationUser is an object representing the database table.
type ReputationUser struct {
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Points    int64     `boil:"points" json:"points" toml:"points" yaml:"points"`

	R *reputationUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reputationUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReputationUserColumns = struct {
	UserID    string
	GuildID   string
	CreatedAt string
	Points    string
}{
	UserID:    "user_id",
	GuildID:   "guild_id",
	CreatedAt: "created_at",
	Points:    "points",
}

var ReputationUserTableColumns = struct {
	UserID    string
	GuildID   string
	CreatedAt string
	Points    string
}{
	UserID:    "reputation_users.user_id",
	GuildID:   "reputation_users.guild_id",
	CreatedAt: "reputation_users.created_at",
	Points:    "reputation_users.points",
}

// Generated where

var ReputationUserWhere = struct {
	UserID    whereHelperint64
	GuildID   whereHelperint64
	CreatedAt whereHelpertime_Time
	Points    whereHelperint64
}{
	UserID:    whereHelperint64{field: "\"reputation_users\".\"user_id\""},
	GuildID:   whereHelperint64{field: "\"reputation_users\".\"guild_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"reputation_users\".\"created_at\""},
	Points:    whereHelperint64{field: "\"reputation_users\".\"points\""},
}

// ReputationUserRels is where relationship names are stored.
var ReputationUserRels = struct {
}{}

// reputationUserR is where relationships are stored.
type reputationUserR struct {
}

// NewStruct creates a new relationship struct
func (*reputationUserR) NewStruct() *reputationUserR {
	return &reputationUserR{}
}

// reputationUserL is where Load methods for each relationship are stored.
type reputationUserL struct{}

var (
	reputationUserAllColumns            = []string{"user_id", "guild_id", "created_at", "points"}
	reputationUserColumnsWithoutDefault = []string{"user_id", "guild_id", "created_at", "points"}
	reputationUserColumnsWithDefault    = []string{}
	reputationUserPrimaryKeyColumns     = []string{"guild_id", "user_id"}
	reputationUserGeneratedColumns      = []string{}
)

type (
	// ReputationUserSlice is an alias for a slice of pointers to ReputationUser.
	// This should almost always be used instead of []ReputationUser.
	ReputationUserSlice []*ReputationUser

	reputationUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reputationUserType                 = reflect.TypeOf(&ReputationUser{})
	reputationUserMapping              = queries.MakeStructMapping(reputationUserType)
	reputationUserPrimaryKeyMapping, _ = queries.BindMapping(reputationUserType, reputationUserMapping, reputationUserPrimaryKeyColumns)
	reputationUserInsertCacheMut       sync.RWMutex
	reputationUserInsertCache          = make(map[string]insertCache)
	reputationUserUpdateCacheMut       sync.RWMutex
	reputationUserUpdateCache          = make(map[string]updateCache)
	reputationUserUpsertCacheMut       sync.RWMutex
	reputationUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single reputationUser record from the query using the global executor.
func (q reputationUserQuery) OneG(ctx context.Context) (*ReputationUser, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single reputationUser record from the query.
func (q reputationUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ReputationUser, error) {
	o := &ReputationUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for reputation_users")
	}

	return o, nil
}

// AllG returns all ReputationUser records from the query using the global executor.
func (q reputationUserQuery) AllG(ctx context.Context) (ReputationUserSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ReputationUser records from the query.
func (q reputationUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReputationUserSlice, error) {
	var o []*ReputationUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ReputationUser slice")
	}

	return o, nil
}

// CountG returns the count of all ReputationUser records in the query using the global executor
func (q reputationUserQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ReputationUser records in the query.
func (q reputationUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count reputation_users rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q reputationUserQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q reputationUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if reputation_users exists")
	}

	return count > 0, nil
}

// ReputationUsers retrieves all the records using an executor.
func ReputationUsers(mods ...qm.QueryMod) reputationUserQuery {
	mods = append(mods, qm.From("\"reputation_users\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"reputation_users\".*"})
	}

	return reputationUserQuery{q}
}

// FindReputationUserG retrieves a single record by ID.
func FindReputationUserG(ctx context.Context, guildID int64, userID int64, selectCols ...string) (*ReputationUser, error) {
	return FindReputationUser(ctx, boil.GetContextDB(), guildID, userID, selectCols...)
}

// FindReputationUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReputationUser(ctx context.Context, exec boil.ContextExecutor, guildID int64, userID int64, selectCols ...string) (*ReputationUser, error) {
	reputationUserObj := &ReputationUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reputation_users\" where \"guild_id\"=$1 AND \"user_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, userID)

	err := q.Bind(ctx, exec, reputationUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from reputation_users")
	}

	return reputationUserObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ReputationUser) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ReputationUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reputation_users provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(reputationUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	reputationUserInsertCacheMut.RLock()
	cache, cached := reputationUserInsertCache[key]
	reputationUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			reputationUserAllColumns,
			reputationUserColumnsWithDefault,
			reputationUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(reputationUserType, reputationUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reputationUserType, reputationUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"reputation_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"reputation_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into reputation_users")
	}

	if !cached {
		reputationUserInsertCacheMut.Lock()
		reputationUserInsertCache[key] = cache
		reputationUserInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ReputationUser record using the global executor.
// See Update for more documentation.
func (o *ReputationUser) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ReputationUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ReputationUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	reputationUserUpdateCacheMut.RLock()
	cache, cached := reputationUserUpdateCache[key]
	reputationUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			reputationUserAllColumns,
			reputationUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update reputation_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reputation_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, reputationUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reputationUserType, reputationUserMapping, append(wl, reputationUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update reputation_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for reputation_users")
	}

	if !cached {
		reputationUserUpdateCacheMut.Lock()
		reputationUserUpdateCache[key] = cache
		reputationUserUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q reputationUserQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q reputationUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for reputation_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for reputation_users")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ReputationUserSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReputationUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"reputation_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, reputationUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in reputationUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all reputationUser")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ReputationUser) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ReputationUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no reputation_users provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(reputationUserColumnsWithDefault, o)

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

	reputationUserUpsertCacheMut.RLock()
	cache, cached := reputationUserUpsertCache[key]
	reputationUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			reputationUserAllColumns,
			reputationUserColumnsWithDefault,
			reputationUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			reputationUserAllColumns,
			reputationUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert reputation_users, could not build update column list")
		}

		ret := strmangle.SetComplement(reputationUserAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(reputationUserPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert reputation_users, could not build conflict column list")
			}

			conflict = make([]string, len(reputationUserPrimaryKeyColumns))
			copy(conflict, reputationUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"reputation_users\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(reputationUserType, reputationUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reputationUserType, reputationUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert reputation_users")
	}

	if !cached {
		reputationUserUpsertCacheMut.Lock()
		reputationUserUpsertCache[key] = cache
		reputationUserUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single ReputationUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ReputationUser) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ReputationUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReputationUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ReputationUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reputationUserPrimaryKeyMapping)
	sql := "DELETE FROM \"reputation_users\" WHERE \"guild_id\"=$1 AND \"user_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from reputation_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for reputation_users")
	}

	return rowsAff, nil
}

func (q reputationUserQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q reputationUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no reputationUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reputation_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reputation_users")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ReputationUserSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReputationUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"reputation_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reputationUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reputationUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reputation_users")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ReputationUser) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ReputationUser provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ReputationUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReputationUser(ctx, exec, o.GuildID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationUserSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ReputationUserSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReputationUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"reputation_users\".* FROM \"reputation_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reputationUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReputationUserSlice")
	}

	*o = slice

	return nil
}

// ReputationUserExistsG checks if the ReputationUser row exists.
func ReputationUserExistsG(ctx context.Context, guildID int64, userID int64) (bool, error) {
	return ReputationUserExists(ctx, boil.GetContextDB(), guildID, userID)
}

// ReputationUserExists checks if the ReputationUser row exists.
func ReputationUserExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, userID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"reputation_users\" where \"guild_id\"=$1 AND \"user_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, guildID, userID)
	}
	row := exec.QueryRowContext(ctx, sql, guildID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if reputation_users exists")
	}

	return exists, nil
}

// Exists checks if the ReputationUser row exists.
func (o *ReputationUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ReputationUserExists(ctx, exec, o.GuildID, o.UserID)
}
