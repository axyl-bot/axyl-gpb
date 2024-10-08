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

// Ticket is an object representing the database table.
type Ticket struct {
	GuildID               int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	LocalID               int64     `boil:"local_id" json:"local_id" toml:"local_id" yaml:"local_id"`
	ChannelID             int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Title                 string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	CreatedAt             time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ClosedAt              null.Time `boil:"closed_at" json:"closed_at,omitempty" toml:"closed_at" yaml:"closed_at,omitempty"`
	LogsID                int64     `boil:"logs_id" json:"logs_id" toml:"logs_id" yaml:"logs_id"`
	AuthorID              int64     `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	AuthorUsernameDiscrim string    `boil:"author_username_discrim" json:"author_username_discrim" toml:"author_username_discrim" yaml:"author_username_discrim"`

	R *ticketR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ticketL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TicketColumns = struct {
	GuildID               string
	LocalID               string
	ChannelID             string
	Title                 string
	CreatedAt             string
	ClosedAt              string
	LogsID                string
	AuthorID              string
	AuthorUsernameDiscrim string
}{
	GuildID:               "guild_id",
	LocalID:               "local_id",
	ChannelID:             "channel_id",
	Title:                 "title",
	CreatedAt:             "created_at",
	ClosedAt:              "closed_at",
	LogsID:                "logs_id",
	AuthorID:              "author_id",
	AuthorUsernameDiscrim: "author_username_discrim",
}

var TicketTableColumns = struct {
	GuildID               string
	LocalID               string
	ChannelID             string
	Title                 string
	CreatedAt             string
	ClosedAt              string
	LogsID                string
	AuthorID              string
	AuthorUsernameDiscrim string
}{
	GuildID:               "tickets.guild_id",
	LocalID:               "tickets.local_id",
	ChannelID:             "tickets.channel_id",
	Title:                 "tickets.title",
	CreatedAt:             "tickets.created_at",
	ClosedAt:              "tickets.closed_at",
	LogsID:                "tickets.logs_id",
	AuthorID:              "tickets.author_id",
	AuthorUsernameDiscrim: "tickets.author_username_discrim",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var TicketWhere = struct {
	GuildID               whereHelperint64
	LocalID               whereHelperint64
	ChannelID             whereHelperint64
	Title                 whereHelperstring
	CreatedAt             whereHelpertime_Time
	ClosedAt              whereHelpernull_Time
	LogsID                whereHelperint64
	AuthorID              whereHelperint64
	AuthorUsernameDiscrim whereHelperstring
}{
	GuildID:               whereHelperint64{field: "\"tickets\".\"guild_id\""},
	LocalID:               whereHelperint64{field: "\"tickets\".\"local_id\""},
	ChannelID:             whereHelperint64{field: "\"tickets\".\"channel_id\""},
	Title:                 whereHelperstring{field: "\"tickets\".\"title\""},
	CreatedAt:             whereHelpertime_Time{field: "\"tickets\".\"created_at\""},
	ClosedAt:              whereHelpernull_Time{field: "\"tickets\".\"closed_at\""},
	LogsID:                whereHelperint64{field: "\"tickets\".\"logs_id\""},
	AuthorID:              whereHelperint64{field: "\"tickets\".\"author_id\""},
	AuthorUsernameDiscrim: whereHelperstring{field: "\"tickets\".\"author_username_discrim\""},
}

// TicketRels is where relationship names are stored.
var TicketRels = struct {
}{}

// ticketR is where relationships are stored.
type ticketR struct {
}

// NewStruct creates a new relationship struct
func (*ticketR) NewStruct() *ticketR {
	return &ticketR{}
}

// ticketL is where Load methods for each relationship are stored.
type ticketL struct{}

var (
	ticketAllColumns            = []string{"guild_id", "local_id", "channel_id", "title", "created_at", "closed_at", "logs_id", "author_id", "author_username_discrim"}
	ticketColumnsWithoutDefault = []string{"guild_id", "local_id", "channel_id", "title", "created_at", "logs_id", "author_id", "author_username_discrim"}
	ticketColumnsWithDefault    = []string{"closed_at"}
	ticketPrimaryKeyColumns     = []string{"guild_id", "local_id"}
	ticketGeneratedColumns      = []string{}
)

type (
	// TicketSlice is an alias for a slice of pointers to Ticket.
	// This should almost always be used instead of []Ticket.
	TicketSlice []*Ticket

	ticketQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ticketType                 = reflect.TypeOf(&Ticket{})
	ticketMapping              = queries.MakeStructMapping(ticketType)
	ticketPrimaryKeyMapping, _ = queries.BindMapping(ticketType, ticketMapping, ticketPrimaryKeyColumns)
	ticketInsertCacheMut       sync.RWMutex
	ticketInsertCache          = make(map[string]insertCache)
	ticketUpdateCacheMut       sync.RWMutex
	ticketUpdateCache          = make(map[string]updateCache)
	ticketUpsertCacheMut       sync.RWMutex
	ticketUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single ticket record from the query using the global executor.
func (q ticketQuery) OneG(ctx context.Context) (*Ticket, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single ticket record from the query.
func (q ticketQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Ticket, error) {
	o := &Ticket{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tickets")
	}

	return o, nil
}

// AllG returns all Ticket records from the query using the global executor.
func (q ticketQuery) AllG(ctx context.Context) (TicketSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Ticket records from the query.
func (q ticketQuery) All(ctx context.Context, exec boil.ContextExecutor) (TicketSlice, error) {
	var o []*Ticket

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Ticket slice")
	}

	return o, nil
}

// CountG returns the count of all Ticket records in the query using the global executor
func (q ticketQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Ticket records in the query.
func (q ticketQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tickets rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q ticketQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q ticketQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tickets exists")
	}

	return count > 0, nil
}

// Tickets retrieves all the records using an executor.
func Tickets(mods ...qm.QueryMod) ticketQuery {
	mods = append(mods, qm.From("\"tickets\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"tickets\".*"})
	}

	return ticketQuery{q}
}

// FindTicketG retrieves a single record by ID.
func FindTicketG(ctx context.Context, guildID int64, localID int64, selectCols ...string) (*Ticket, error) {
	return FindTicket(ctx, boil.GetContextDB(), guildID, localID, selectCols...)
}

// FindTicket retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTicket(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64, selectCols ...string) (*Ticket, error) {
	ticketObj := &Ticket{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tickets\" where \"guild_id\"=$1 AND \"local_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, localID)

	err := q.Bind(ctx, exec, ticketObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tickets")
	}

	return ticketObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Ticket) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Ticket) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tickets provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ticketInsertCacheMut.RLock()
	cache, cached := ticketInsertCache[key]
	ticketInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ticketAllColumns,
			ticketColumnsWithDefault,
			ticketColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ticketType, ticketMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ticketType, ticketMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tickets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tickets\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into tickets")
	}

	if !cached {
		ticketInsertCacheMut.Lock()
		ticketInsertCache[key] = cache
		ticketInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Ticket record using the global executor.
// See Update for more documentation.
func (o *Ticket) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Ticket.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Ticket) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	ticketUpdateCacheMut.RLock()
	cache, cached := ticketUpdateCache[key]
	ticketUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ticketAllColumns,
			ticketPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tickets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tickets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ticketPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ticketType, ticketMapping, append(wl, ticketPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tickets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tickets")
	}

	if !cached {
		ticketUpdateCacheMut.Lock()
		ticketUpdateCache[key] = cache
		ticketUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q ticketQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q ticketQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tickets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tickets")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TicketSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TicketSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tickets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ticketPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ticket slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ticket")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Ticket) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Ticket) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no tickets provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketColumnsWithDefault, o)

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

	ticketUpsertCacheMut.RLock()
	cache, cached := ticketUpsertCache[key]
	ticketUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			ticketAllColumns,
			ticketColumnsWithDefault,
			ticketColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			ticketAllColumns,
			ticketPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tickets, could not build update column list")
		}

		ret := strmangle.SetComplement(ticketAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(ticketPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert tickets, could not build conflict column list")
			}

			conflict = make([]string, len(ticketPrimaryKeyColumns))
			copy(conflict, ticketPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tickets\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(ticketType, ticketMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ticketType, ticketMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert tickets")
	}

	if !cached {
		ticketUpsertCacheMut.Lock()
		ticketUpsertCache[key] = cache
		ticketUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Ticket record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Ticket) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Ticket record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Ticket) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Ticket provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ticketPrimaryKeyMapping)
	sql := "DELETE FROM \"tickets\" WHERE \"guild_id\"=$1 AND \"local_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tickets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tickets")
	}

	return rowsAff, nil
}

func (q ticketQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q ticketQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ticketQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tickets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tickets")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TicketSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TicketSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tickets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ticket slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tickets")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Ticket) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Ticket provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Ticket) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTicket(ctx, exec, o.GuildID, o.LocalID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TicketSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TicketSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tickets\".* FROM \"tickets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TicketSlice")
	}

	*o = slice

	return nil
}

// TicketExistsG checks if the Ticket row exists.
func TicketExistsG(ctx context.Context, guildID int64, localID int64) (bool, error) {
	return TicketExists(ctx, boil.GetContextDB(), guildID, localID)
}

// TicketExists checks if the Ticket row exists.
func TicketExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tickets\" where \"guild_id\"=$1 AND \"local_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, guildID, localID)
	}
	row := exec.QueryRowContext(ctx, sql, guildID, localID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tickets exists")
	}

	return exists, nil
}

// Exists checks if the Ticket row exists.
func (o *Ticket) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TicketExists(ctx, exec, o.GuildID, o.LocalID)
}
