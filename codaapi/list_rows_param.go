package codaapi

import (
	"encoding/json"
	"fmt"
)

type ListRowsParam func(target *ListRowsParams)

var ListRows = struct {
	Query          func(columnNameOrID, value string) ListRowsParam
	SortBy         func(value RowsSortBy) ListRowsParam // Can be used with codaapi.RowsSortBy* constants to specify value
	UseColumnNames func(value bool) ListRowsParam
	VisibleOnly    func(value bool) ListRowsParam
	Limit          func(value int) ListRowsParam
	SyncToken      func(value string) ListRowsParam
	ValueFormat    func(value ValueFormat) ListRowsParam // Can be used with codaapi.ValueFormat* constants to specify value
	PageToken      func(value string) ListRowsParam
}{
	Query: func(columnNameOrID, value string) ListRowsParam {
		valueBs, _ := json.Marshal(value)
		query := fmt.Sprintf("%s:%s", columnNameOrID, valueBs)

		return func(target *ListRowsParams) {
			target.Query = &query
		}
	},
	SortBy: func(value RowsSortBy) ListRowsParam {
		return func(target *ListRowsParams) {
			target.SortBy = &value
		}
	},
	UseColumnNames: func(value bool) ListRowsParam {
		var v = UseColumnNames(value)
		return func(target *ListRowsParams) {
			target.UseColumnNames = &v
		}
	},
	ValueFormat: func(value ValueFormat) ListRowsParam {
		return func(target *ListRowsParams) {
			target.ValueFormat = &value
		}
	},
	VisibleOnly: func(value bool) ListRowsParam {
		return func(target *ListRowsParams) {
			target.VisibleOnly = &value
		}
	},
	Limit: func(value int) ListRowsParam {
		var v = Limit(value)
		return func(target *ListRowsParams) {
			target.Limit = &v
		}
	},
	PageToken: func(value string) ListRowsParam {
		var v = PageToken(value)
		return func(target *ListRowsParams) {
			target.PageToken = &v
		}
	},
	SyncToken: func(value string) ListRowsParam {
		var v = SyncToken(value)
		return func(target *ListRowsParams) {
			target.SyncToken = &v
		}
	},
}
