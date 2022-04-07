package models

type sqlWhere struct {
	selectTable string
	selectRaws  string
	where       string
	limit       string
	orderBy     string
	groupBy     string
	params      []interface {
	}
}
//
//var selectSql = "select `?` from `?` where ? limit ?"
//var updateSql = "select `?` from `?` where ? "
//var deleteSql = "select `?` from `?` where ? "
//
//func NewSqlBuilds() *sqlWhere {
//	return new(sqlWhere)
//}
//
//func (sql *sqlWhere) Table(table string) *sqlWhere {
//	sql.selectTable = table
//	return sql
//}
//
//func (sql *sqlWhere) Select(raw ...string) *sqlWhere {
//	sql.selectRaws = strings.Join(raw, ",")
//	return sql
//}
//
//func (sql *sqlWhere) Where(key, operation string, value interface{}) *sqlWhere {
//
//}
//
//func (sql *sqlWhere) WhereBetween() *sqlWhere {
//
//}
//
//func (sql *sqlWhere) WhereIn() *sqlWhere {
//
//}
//
//func (sql *sqlWhere) ToSql() {
//
//}
