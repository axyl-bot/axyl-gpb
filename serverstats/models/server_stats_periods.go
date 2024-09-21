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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ServerStatsPeriod is an object representing the database table.
type ServerStatsPeriod struct {
	ID        int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Started   null.Time  `boil:"started" json:"started,omitempty" toml:"started" yaml:"started,omitempty"`
	Duration  null.Int64 `boil:"duration" json:"duration,omitempty" toml:"duration" yaml:"duration,omitempty"`
	GuildID   null.Int64 `boil:"guild_id" json:"guild_id,omitempty" toml:"guild_id" yaml:"guild_id,omitempty"`
	UserID    null.Int64 `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	ChannelID null.Int64 `boil:"channel_id" json:"channel_id,omitempty" toml:"channel_id" yaml:"channel_id,omitempty"`
	Count     null.Int64 `boil:"count" json:"count,omitempty" toml:"count" yaml:"count,omitempty"`

	R *serverStatsPeriodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serverStatsPeriodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServerStatsPeriodColumns = struct {
	ID        string
	Started   string
	Duration  string
	GuildID   string
	UserID    string
	ChannelID string
	Count     string
}{
	ID:        "id",
	Started:   "started",
	Duration:  "duration",
	GuildID:   "guild_id",
	UserID:    "user_id",
	ChannelID: "channel_id",
	Count:     "count",
}

var ServerStatsPeriodTableColumns = struct {
	ID        string
	Started   string
	Duration  string
	GuildID   string
	UserID    string
	ChannelID string
	Count     string
}{
	ID:        "server_stats_periods.id",
	Started:   "server_stats_periods.started",
	Duration:  "server_stats_periods.duration",
	GuildID:   "server_stats_periods.guild_id",
	UserID:    "server_stats_periods.user_id",
	ChannelID: "server_stats_periods.channel_id",
	Count:     "server_stats_periods.count",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ServerStatsPeriodWhere = struct {
	ID        whereHelperint64
	Started   whereHelpernull_Time
	Duration  whereHelpernull_Int64
	GuildID   whereHelpernull_Int64
	UserID    whereHelpernull_Int64
	ChannelID whereHelpernull_Int64
	Count     whereHelpernull_Int64
}{
	ID:        whereHelperint64{field: "\"server_stats_periods\".\"id\""},
	Started:   whereHelpernull_Time{field: "\"server_stats_periods\".\"started\""},
	Duration:  whereHelpernull_Int64{field: "\"server_stats_periods\".\"duration\""},
	GuildID:   whereHelpernull_Int64{field: "\"server_stats_periods\".\"guild_id\""},
	UserID:    whereHelpernull_Int64{field: "\"server_stats_periods\".\"user_id\""},
	ChannelID: whereHelpernull_Int64{field: "\"server_stats_periods\".\"channel_id\""},
	Count:     whereHelpernull_Int64{field: "\"server_stats_periods\".\"count\""},
}

// ServerStatsPeriodRels is where relationship names are stored.
var ServerStatsPeriodRels = struct {
}{}

// serverStatsPeriodR is where relationships are stored.
type serverStatsPeriodR struct {
}

// NewStruct creates a new relationship struct
func (*serverStatsPeriodR) NewStruct() *serverStatsPeriodR {
	return &serverStatsPeriodR{}
}

// serverStatsPeriodL is where Load methods for each relationship are stored.
type serverStatsPeriodL struct{}

var (
	serverStatsPeriodAllColumns            = []string{"id", "started", "duration", "guild_id", "user_id", "channel_id", "count"}
	serverStatsPeriodColumnsWithoutDefault = []string{}
	serverStatsPeriodColumnsWithDefault    = []string{"id", "started", "duration", "guild_id", "user_id", "channel_id", "count"}
	serverStatsPeriodPrimaryKeyColumns     = []string{"id"}
	serverStatsPeriodGeneratedColumns      = []string{}
)

type (
	// ServerStatsPeriodSlice is an alias for a slice of pointers to ServerStatsPeriod.
	// This should almost always be used instead of []ServerStatsPeriod.
	ServerStatsPeriodSlice []*ServerStatsPeriod

	serverStatsPeriodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serverStatsPeriodType                 = reflect.TypeOf(&ServerStatsPeriod{})
	serverStatsPeriodMapping              = queries.MakeStructMapping(serverStatsPeriodType)
	serverStatsPeriodPrimaryKeyMapping, _ = queries.BindMapping(serverStatsPeriodType, serverStatsPeriodMapping, serverStatsPeriodPrimaryKeyColumns)
	serverStatsPeriodInsertCacheMut       sync.RWMutex
	serverStatsPeriodInsertCache          = make(map[string]insertCache)
	serverStatsPeriodUpdateCacheMut       sync.RWMutex
	serverStatsPeriodUpdateCache          = make(map[string]updateCache)
	serverStatsPeriodUpsertCacheMut       sync.RWMutex
	serverStatsPeriodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single serverStatsPeriod record from the query using the global executor.
func (q serverStatsPeriodQuery) OneG(ctx context.Context) (*ServerStatsPeriod, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single serverStatsPeriod record from the query.
func (q serverStatsPeriodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServerStatsPeriod, error) {
	o := &ServerStatsPeriod{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for server_stats_periods")
	}

	return o, nil
}

// AllG returns all ServerStatsPeriod records from the query using the global executor.
func (q serverStatsPeriodQuery) AllG(ctx context.Context) (ServerStatsPeriodSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ServerStatsPeriod records from the query.
func (q serverStatsPeriodQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServerStatsPeriodSlice, error) {
	var o []*ServerStatsPeriod

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServerStatsPeriod slice")
	}

	return o, nil
}

// CountG returns the count of all ServerStatsPeriod records in the query using the global executor
func (q serverStatsPeriodQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ServerStatsPeriod records in the query.
func (q serverStatsPeriodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count server_stats_periods rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q serverStatsPeriodQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q serverStatsPeriodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if server_stats_periods exists")
	}

	return count > 0, nil
}

// ServerStatsPeriods retrieves all the records using an executor.
func ServerStatsPeriods(mods ...qm.QueryMod) serverStatsPeriodQuery {
	mods = append(mods, qm.From("\"server_stats_periods\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"server_stats_periods\".*"})
	}

	return serverStatsPeriodQuery{q}
}

// FindServerStatsPeriodG retrieves a single record by ID.
func FindServerStatsPeriodG(ctx context.Context, iD int64, selectCols ...string) (*ServerStatsPeriod, error) {
	return FindServerStatsPeriod(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindServerStatsPeriod retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServerStatsPeriod(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ServerStatsPeriod, error) {
	serverStatsPeriodObj := &ServerStatsPeriod{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"server_stats_periods\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serverStatsPeriodObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from server_stats_periods")
	}

	return serverStatsPeriodObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ServerStatsPeriod) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ServerStatsPeriod) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no server_stats_periods provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(serverStatsPeriodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serverStatsPeriodInsertCacheMut.RLock()
	cache, cached := serverStatsPeriodInsertCache[key]
	serverStatsPeriodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serverStatsPeriodAllColumns,
			serverStatsPeriodColumnsWithDefault,
			serverStatsPeriodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serverStatsPeriodType, serverStatsPeriodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serverStatsPeriodType, serverStatsPeriodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"server_stats_periods\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"server_stats_periods\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into server_stats_periods")
	}

	if !cached {
		serverStatsPeriodInsertCacheMut.Lock()
		serverStatsPeriodInsertCache[key] = cache
		serverStatsPeriodInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ServerStatsPeriod record using the global executor.
// See Update for more documentation.
func (o *ServerStatsPeriod) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ServerStatsPeriod.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ServerStatsPeriod) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	serverStatsPeriodUpdateCacheMut.RLock()
	cache, cached := serverStatsPeriodUpdateCache[key]
	serverStatsPeriodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serverStatsPeriodAllColumns,
			serverStatsPeriodPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update server_stats_periods, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"server_stats_periods\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serverStatsPeriodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serverStatsPeriodType, serverStatsPeriodMapping, append(wl, serverStatsPeriodPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update server_stats_periods row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for server_stats_periods")
	}

	if !cached {
		serverStatsPeriodUpdateCacheMut.Lock()
		serverStatsPeriodUpdateCache[key] = cache
		serverStatsPeriodUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q serverStatsPeriodQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q serverStatsPeriodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for server_stats_periods")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for server_stats_periods")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ServerStatsPeriodSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServerStatsPeriodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsPeriodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"server_stats_periods\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serverStatsPeriodPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serverStatsPeriod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serverStatsPeriod")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ServerStatsPeriod) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServerStatsPeriod) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no server_stats_periods provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(serverStatsPeriodColumnsWithDefault, o)

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

	serverStatsPeriodUpsertCacheMut.RLock()
	cache, cached := serverStatsPeriodUpsertCache[key]
	serverStatsPeriodUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			serverStatsPeriodAllColumns,
			serverStatsPeriodColumnsWithDefault,
			serverStatsPeriodColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			serverStatsPeriodAllColumns,
			serverStatsPeriodPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert server_stats_periods, could not build update column list")
		}

		ret := strmangle.SetComplement(serverStatsPeriodAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(serverStatsPeriodPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert server_stats_periods, could not build conflict column list")
			}

			conflict = make([]string, len(serverStatsPeriodPrimaryKeyColumns))
			copy(conflict, serverStatsPeriodPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"server_stats_periods\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(serverStatsPeriodType, serverStatsPeriodMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serverStatsPeriodType, serverStatsPeriodMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert server_stats_periods")
	}

	if !cached {
		serverStatsPeriodUpsertCacheMut.Lock()
		serverStatsPeriodUpsertCache[key] = cache
		serverStatsPeriodUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single ServerStatsPeriod record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ServerStatsPeriod) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ServerStatsPeriod record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ServerStatsPeriod) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServerStatsPeriod provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serverStatsPeriodPrimaryKeyMapping)
	sql := "DELETE FROM \"server_stats_periods\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from server_stats_periods")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for server_stats_periods")
	}

	return rowsAff, nil
}

func (q serverStatsPeriodQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q serverStatsPeriodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serverStatsPeriodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from server_stats_periods")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for server_stats_periods")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ServerStatsPeriodSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServerStatsPeriodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsPeriodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"server_stats_periods\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serverStatsPeriodPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serverStatsPeriod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for server_stats_periods")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ServerStatsPeriod) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ServerStatsPeriod provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ServerStatsPeriod) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServerStatsPeriod(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServerStatsPeriodSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ServerStatsPeriodSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServerStatsPeriodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServerStatsPeriodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serverStatsPeriodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"server_stats_periods\".* FROM \"server_stats_periods\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serverStatsPeriodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServerStatsPeriodSlice")
	}

	*o = slice

	return nil
}

// ServerStatsPeriodExistsG checks if the ServerStatsPeriod row exists.
func ServerStatsPeriodExistsG(ctx context.Context, iD int64) (bool, error) {
	return ServerStatsPeriodExists(ctx, boil.GetContextDB(), iD)
}

// ServerStatsPeriodExists checks if the ServerStatsPeriod row exists.
func ServerStatsPeriodExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"server_stats_periods\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if server_stats_periods exists")
	}

	return exists, nil
}

// Exists checks if the ServerStatsPeriod row exists.
func (o *ServerStatsPeriod) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ServerStatsPeriodExists(ctx, exec, o.ID)
}
