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

// AutomodRuleset is an object representing the database table.
type AutomodRuleset struct {
	ID      int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID int64  `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Enabled bool   `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`

	R *automodRulesetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L automodRulesetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AutomodRulesetColumns = struct {
	ID      string
	GuildID string
	Name    string
	Enabled string
}{
	ID:      "id",
	GuildID: "guild_id",
	Name:    "name",
	Enabled: "enabled",
}

var AutomodRulesetTableColumns = struct {
	ID      string
	GuildID string
	Name    string
	Enabled string
}{
	ID:      "automod_rulesets.id",
	GuildID: "automod_rulesets.guild_id",
	Name:    "automod_rulesets.name",
	Enabled: "automod_rulesets.enabled",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var AutomodRulesetWhere = struct {
	ID      whereHelperint64
	GuildID whereHelperint64
	Name    whereHelperstring
	Enabled whereHelperbool
}{
	ID:      whereHelperint64{field: "\"automod_rulesets\".\"id\""},
	GuildID: whereHelperint64{field: "\"automod_rulesets\".\"guild_id\""},
	Name:    whereHelperstring{field: "\"automod_rulesets\".\"name\""},
	Enabled: whereHelperbool{field: "\"automod_rulesets\".\"enabled\""},
}

// AutomodRulesetRels is where relationship names are stored.
var AutomodRulesetRels = struct {
	RulesetAutomodRules             string
	RulesetAutomodRulesetConditions string
}{
	RulesetAutomodRules:             "RulesetAutomodRules",
	RulesetAutomodRulesetConditions: "RulesetAutomodRulesetConditions",
}

// automodRulesetR is where relationships are stored.
type automodRulesetR struct {
	RulesetAutomodRules             AutomodRuleSlice             `boil:"RulesetAutomodRules" json:"RulesetAutomodRules" toml:"RulesetAutomodRules" yaml:"RulesetAutomodRules"`
	RulesetAutomodRulesetConditions AutomodRulesetConditionSlice `boil:"RulesetAutomodRulesetConditions" json:"RulesetAutomodRulesetConditions" toml:"RulesetAutomodRulesetConditions" yaml:"RulesetAutomodRulesetConditions"`
}

// NewStruct creates a new relationship struct
func (*automodRulesetR) NewStruct() *automodRulesetR {
	return &automodRulesetR{}
}

func (r *automodRulesetR) GetRulesetAutomodRules() AutomodRuleSlice {
	if r == nil {
		return nil
	}
	return r.RulesetAutomodRules
}

func (r *automodRulesetR) GetRulesetAutomodRulesetConditions() AutomodRulesetConditionSlice {
	if r == nil {
		return nil
	}
	return r.RulesetAutomodRulesetConditions
}

// automodRulesetL is where Load methods for each relationship are stored.
type automodRulesetL struct{}

var (
	automodRulesetAllColumns            = []string{"id", "guild_id", "name", "enabled"}
	automodRulesetColumnsWithoutDefault = []string{"guild_id", "name", "enabled"}
	automodRulesetColumnsWithDefault    = []string{"id"}
	automodRulesetPrimaryKeyColumns     = []string{"id"}
	automodRulesetGeneratedColumns      = []string{}
)

type (
	// AutomodRulesetSlice is an alias for a slice of pointers to AutomodRuleset.
	// This should almost always be used instead of []AutomodRuleset.
	AutomodRulesetSlice []*AutomodRuleset

	automodRulesetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	automodRulesetType                 = reflect.TypeOf(&AutomodRuleset{})
	automodRulesetMapping              = queries.MakeStructMapping(automodRulesetType)
	automodRulesetPrimaryKeyMapping, _ = queries.BindMapping(automodRulesetType, automodRulesetMapping, automodRulesetPrimaryKeyColumns)
	automodRulesetInsertCacheMut       sync.RWMutex
	automodRulesetInsertCache          = make(map[string]insertCache)
	automodRulesetUpdateCacheMut       sync.RWMutex
	automodRulesetUpdateCache          = make(map[string]updateCache)
	automodRulesetUpsertCacheMut       sync.RWMutex
	automodRulesetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single automodRuleset record from the query using the global executor.
func (q automodRulesetQuery) OneG(ctx context.Context) (*AutomodRuleset, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single automodRuleset record from the query.
func (q automodRulesetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AutomodRuleset, error) {
	o := &AutomodRuleset{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for automod_rulesets")
	}

	return o, nil
}

// AllG returns all AutomodRuleset records from the query using the global executor.
func (q automodRulesetQuery) AllG(ctx context.Context) (AutomodRulesetSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all AutomodRuleset records from the query.
func (q automodRulesetQuery) All(ctx context.Context, exec boil.ContextExecutor) (AutomodRulesetSlice, error) {
	var o []*AutomodRuleset

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AutomodRuleset slice")
	}

	return o, nil
}

// CountG returns the count of all AutomodRuleset records in the query using the global executor
func (q automodRulesetQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all AutomodRuleset records in the query.
func (q automodRulesetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count automod_rulesets rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q automodRulesetQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q automodRulesetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if automod_rulesets exists")
	}

	return count > 0, nil
}

// RulesetAutomodRules retrieves all the automod_rule's AutomodRules with an executor via ruleset_id column.
func (o *AutomodRuleset) RulesetAutomodRules(mods ...qm.QueryMod) automodRuleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"automod_rules\".\"ruleset_id\"=?", o.ID),
	)

	return AutomodRules(queryMods...)
}

// RulesetAutomodRulesetConditions retrieves all the automod_ruleset_condition's AutomodRulesetConditions with an executor via ruleset_id column.
func (o *AutomodRuleset) RulesetAutomodRulesetConditions(mods ...qm.QueryMod) automodRulesetConditionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"automod_ruleset_conditions\".\"ruleset_id\"=?", o.ID),
	)

	return AutomodRulesetConditions(queryMods...)
}

// LoadRulesetAutomodRules allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (automodRulesetL) LoadRulesetAutomodRules(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutomodRuleset interface{}, mods queries.Applicator) error {
	var slice []*AutomodRuleset
	var object *AutomodRuleset

	if singular {
		var ok bool
		object, ok = maybeAutomodRuleset.(*AutomodRuleset)
		if !ok {
			object = new(AutomodRuleset)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAutomodRuleset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAutomodRuleset))
			}
		}
	} else {
		s, ok := maybeAutomodRuleset.(*[]*AutomodRuleset)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAutomodRuleset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAutomodRuleset))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &automodRulesetR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &automodRulesetR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`automod_rules`),
		qm.WhereIn(`automod_rules.ruleset_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load automod_rules")
	}

	var resultSlice []*AutomodRule
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice automod_rules")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on automod_rules")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for automod_rules")
	}

	if singular {
		object.R.RulesetAutomodRules = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &automodRuleR{}
			}
			foreign.R.Ruleset = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RulesetID {
				local.R.RulesetAutomodRules = append(local.R.RulesetAutomodRules, foreign)
				if foreign.R == nil {
					foreign.R = &automodRuleR{}
				}
				foreign.R.Ruleset = local
				break
			}
		}
	}

	return nil
}

// LoadRulesetAutomodRulesetConditions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (automodRulesetL) LoadRulesetAutomodRulesetConditions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutomodRuleset interface{}, mods queries.Applicator) error {
	var slice []*AutomodRuleset
	var object *AutomodRuleset

	if singular {
		var ok bool
		object, ok = maybeAutomodRuleset.(*AutomodRuleset)
		if !ok {
			object = new(AutomodRuleset)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAutomodRuleset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAutomodRuleset))
			}
		}
	} else {
		s, ok := maybeAutomodRuleset.(*[]*AutomodRuleset)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAutomodRuleset)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAutomodRuleset))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &automodRulesetR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &automodRulesetR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`automod_ruleset_conditions`),
		qm.WhereIn(`automod_ruleset_conditions.ruleset_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load automod_ruleset_conditions")
	}

	var resultSlice []*AutomodRulesetCondition
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice automod_ruleset_conditions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on automod_ruleset_conditions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for automod_ruleset_conditions")
	}

	if singular {
		object.R.RulesetAutomodRulesetConditions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &automodRulesetConditionR{}
			}
			foreign.R.Ruleset = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RulesetID {
				local.R.RulesetAutomodRulesetConditions = append(local.R.RulesetAutomodRulesetConditions, foreign)
				if foreign.R == nil {
					foreign.R = &automodRulesetConditionR{}
				}
				foreign.R.Ruleset = local
				break
			}
		}
	}

	return nil
}

// AddRulesetAutomodRulesG adds the given related objects to the existing relationships
// of the automod_ruleset, optionally inserting them as new records.
// Appends related to o.R.RulesetAutomodRules.
// Sets related.R.Ruleset appropriately.
// Uses the global database handle.
func (o *AutomodRuleset) AddRulesetAutomodRulesG(ctx context.Context, insert bool, related ...*AutomodRule) error {
	return o.AddRulesetAutomodRules(ctx, boil.GetContextDB(), insert, related...)
}

// AddRulesetAutomodRules adds the given related objects to the existing relationships
// of the automod_ruleset, optionally inserting them as new records.
// Appends related to o.R.RulesetAutomodRules.
// Sets related.R.Ruleset appropriately.
func (o *AutomodRuleset) AddRulesetAutomodRules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AutomodRule) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RulesetID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"automod_rules\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"ruleset_id"}),
				strmangle.WhereClause("\"", "\"", 2, automodRulePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RulesetID = o.ID
		}
	}

	if o.R == nil {
		o.R = &automodRulesetR{
			RulesetAutomodRules: related,
		}
	} else {
		o.R.RulesetAutomodRules = append(o.R.RulesetAutomodRules, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &automodRuleR{
				Ruleset: o,
			}
		} else {
			rel.R.Ruleset = o
		}
	}
	return nil
}

// AddRulesetAutomodRulesetConditionsG adds the given related objects to the existing relationships
// of the automod_ruleset, optionally inserting them as new records.
// Appends related to o.R.RulesetAutomodRulesetConditions.
// Sets related.R.Ruleset appropriately.
// Uses the global database handle.
func (o *AutomodRuleset) AddRulesetAutomodRulesetConditionsG(ctx context.Context, insert bool, related ...*AutomodRulesetCondition) error {
	return o.AddRulesetAutomodRulesetConditions(ctx, boil.GetContextDB(), insert, related...)
}

// AddRulesetAutomodRulesetConditions adds the given related objects to the existing relationships
// of the automod_ruleset, optionally inserting them as new records.
// Appends related to o.R.RulesetAutomodRulesetConditions.
// Sets related.R.Ruleset appropriately.
func (o *AutomodRuleset) AddRulesetAutomodRulesetConditions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AutomodRulesetCondition) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RulesetID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"automod_ruleset_conditions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"ruleset_id"}),
				strmangle.WhereClause("\"", "\"", 2, automodRulesetConditionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RulesetID = o.ID
		}
	}

	if o.R == nil {
		o.R = &automodRulesetR{
			RulesetAutomodRulesetConditions: related,
		}
	} else {
		o.R.RulesetAutomodRulesetConditions = append(o.R.RulesetAutomodRulesetConditions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &automodRulesetConditionR{
				Ruleset: o,
			}
		} else {
			rel.R.Ruleset = o
		}
	}
	return nil
}

// AutomodRulesets retrieves all the records using an executor.
func AutomodRulesets(mods ...qm.QueryMod) automodRulesetQuery {
	mods = append(mods, qm.From("\"automod_rulesets\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"automod_rulesets\".*"})
	}

	return automodRulesetQuery{q}
}

// FindAutomodRulesetG retrieves a single record by ID.
func FindAutomodRulesetG(ctx context.Context, iD int64, selectCols ...string) (*AutomodRuleset, error) {
	return FindAutomodRuleset(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAutomodRuleset retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAutomodRuleset(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AutomodRuleset, error) {
	automodRulesetObj := &AutomodRuleset{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"automod_rulesets\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, automodRulesetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from automod_rulesets")
	}

	return automodRulesetObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AutomodRuleset) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AutomodRuleset) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no automod_rulesets provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(automodRulesetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	automodRulesetInsertCacheMut.RLock()
	cache, cached := automodRulesetInsertCache[key]
	automodRulesetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			automodRulesetAllColumns,
			automodRulesetColumnsWithDefault,
			automodRulesetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(automodRulesetType, automodRulesetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(automodRulesetType, automodRulesetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"automod_rulesets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"automod_rulesets\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into automod_rulesets")
	}

	if !cached {
		automodRulesetInsertCacheMut.Lock()
		automodRulesetInsertCache[key] = cache
		automodRulesetInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single AutomodRuleset record using the global executor.
// See Update for more documentation.
func (o *AutomodRuleset) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the AutomodRuleset.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AutomodRuleset) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	automodRulesetUpdateCacheMut.RLock()
	cache, cached := automodRulesetUpdateCache[key]
	automodRulesetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			automodRulesetAllColumns,
			automodRulesetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update automod_rulesets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"automod_rulesets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, automodRulesetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(automodRulesetType, automodRulesetMapping, append(wl, automodRulesetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update automod_rulesets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for automod_rulesets")
	}

	if !cached {
		automodRulesetUpdateCacheMut.Lock()
		automodRulesetUpdateCache[key] = cache
		automodRulesetUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q automodRulesetQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q automodRulesetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for automod_rulesets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for automod_rulesets")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AutomodRulesetSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AutomodRulesetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRulesetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"automod_rulesets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, automodRulesetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in automodRuleset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all automodRuleset")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AutomodRuleset) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AutomodRuleset) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no automod_rulesets provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(automodRulesetColumnsWithDefault, o)

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

	automodRulesetUpsertCacheMut.RLock()
	cache, cached := automodRulesetUpsertCache[key]
	automodRulesetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			automodRulesetAllColumns,
			automodRulesetColumnsWithDefault,
			automodRulesetColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			automodRulesetAllColumns,
			automodRulesetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert automod_rulesets, could not build update column list")
		}

		ret := strmangle.SetComplement(automodRulesetAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(automodRulesetPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert automod_rulesets, could not build conflict column list")
			}

			conflict = make([]string, len(automodRulesetPrimaryKeyColumns))
			copy(conflict, automodRulesetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"automod_rulesets\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(automodRulesetType, automodRulesetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(automodRulesetType, automodRulesetMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert automod_rulesets")
	}

	if !cached {
		automodRulesetUpsertCacheMut.Lock()
		automodRulesetUpsertCache[key] = cache
		automodRulesetUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single AutomodRuleset record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AutomodRuleset) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single AutomodRuleset record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AutomodRuleset) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AutomodRuleset provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), automodRulesetPrimaryKeyMapping)
	sql := "DELETE FROM \"automod_rulesets\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from automod_rulesets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for automod_rulesets")
	}

	return rowsAff, nil
}

func (q automodRulesetQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q automodRulesetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no automodRulesetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automod_rulesets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_rulesets")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AutomodRulesetSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AutomodRulesetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRulesetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"automod_rulesets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodRulesetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automodRuleset slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_rulesets")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AutomodRuleset) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no AutomodRuleset provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AutomodRuleset) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAutomodRuleset(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodRulesetSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty AutomodRulesetSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodRulesetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AutomodRulesetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRulesetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"automod_rulesets\".* FROM \"automod_rulesets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodRulesetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AutomodRulesetSlice")
	}

	*o = slice

	return nil
}

// AutomodRulesetExistsG checks if the AutomodRuleset row exists.
func AutomodRulesetExistsG(ctx context.Context, iD int64) (bool, error) {
	return AutomodRulesetExists(ctx, boil.GetContextDB(), iD)
}

// AutomodRulesetExists checks if the AutomodRuleset row exists.
func AutomodRulesetExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"automod_rulesets\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if automod_rulesets exists")
	}

	return exists, nil
}

// Exists checks if the AutomodRuleset row exists.
func (o *AutomodRuleset) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AutomodRulesetExists(ctx, exec, o.ID)
}
