// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package commentsb

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Comments", testComments)
	t.Run("OutboxEvents", testOutboxEvents)
	t.Run("Replies", testReplies)
	t.Run("SchemaMigrations", testSchemaMigrations)
}

func TestDelete(t *testing.T) {
	t.Run("Comments", testCommentsDelete)
	t.Run("OutboxEvents", testOutboxEventsDelete)
	t.Run("Replies", testRepliesDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("OutboxEvents", testOutboxEventsQueryDeleteAll)
	t.Run("Replies", testRepliesQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("OutboxEvents", testOutboxEventsSliceDeleteAll)
	t.Run("Replies", testRepliesSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Comments", testCommentsExists)
	t.Run("OutboxEvents", testOutboxEventsExists)
	t.Run("Replies", testRepliesExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
}

func TestFind(t *testing.T) {
	t.Run("Comments", testCommentsFind)
	t.Run("OutboxEvents", testOutboxEventsFind)
	t.Run("Replies", testRepliesFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
}

func TestBind(t *testing.T) {
	t.Run("Comments", testCommentsBind)
	t.Run("OutboxEvents", testOutboxEventsBind)
	t.Run("Replies", testRepliesBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
}

func TestOne(t *testing.T) {
	t.Run("Comments", testCommentsOne)
	t.Run("OutboxEvents", testOutboxEventsOne)
	t.Run("Replies", testRepliesOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
}

func TestAll(t *testing.T) {
	t.Run("Comments", testCommentsAll)
	t.Run("OutboxEvents", testOutboxEventsAll)
	t.Run("Replies", testRepliesAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
}

func TestCount(t *testing.T) {
	t.Run("Comments", testCommentsCount)
	t.Run("OutboxEvents", testOutboxEventsCount)
	t.Run("Replies", testRepliesCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Comments", testCommentsHooks)
	t.Run("OutboxEvents", testOutboxEventsHooks)
	t.Run("Replies", testRepliesHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("OutboxEvents", testOutboxEventsInsert)
	t.Run("OutboxEvents", testOutboxEventsInsertWhitelist)
	t.Run("Replies", testRepliesInsert)
	t.Run("Replies", testRepliesInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Comments", testCommentsReload)
	t.Run("OutboxEvents", testOutboxEventsReload)
	t.Run("Replies", testRepliesReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Comments", testCommentsReloadAll)
	t.Run("OutboxEvents", testOutboxEventsReloadAll)
	t.Run("Replies", testRepliesReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Comments", testCommentsSelect)
	t.Run("OutboxEvents", testOutboxEventsSelect)
	t.Run("Replies", testRepliesSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Comments", testCommentsUpdate)
	t.Run("OutboxEvents", testOutboxEventsUpdate)
	t.Run("Replies", testRepliesUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("OutboxEvents", testOutboxEventsSliceUpdateAll)
	t.Run("Replies", testRepliesSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
}
