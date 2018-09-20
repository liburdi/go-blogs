package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"blog/db"
)

// ModelMethod 接口
type ModelMethod interface {
	Query(offset int64, limit int64) (map[string]string, error)
	Create() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	QueryOne(int64) (map[string]string, error)
}

// ModelProperty 属性
type ModelProperty struct {
	TableName string
}

// WhereCondition where条件
type WhereCondition struct {
	Operator string
	Value    string
}

// WhereValues where条件值
type WhereValues map[string]WhereCondition

// UpdateValues update条件值
type UpdateValues map[string]string

// InsertValues Insert值
type InsertValues map[string]string

// SelectValues select条件值
type SelectValues map[string]interface{}

// SelectResult 结果值
type SelectResult map[string]interface{}

// Conn 连接
var Conn *sql.DB

func init() {
	// MySQL
	Conn = db.MySQLDB
}

// Query 获取数据
func (p *ModelProperty) Query(s SelectValues, where WhereValues, offset int64, limit int64) ([]SelectResult, error) {
	var selectString = s.MergeSelect()

	var whereString = where.MergeWhere()

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d, %d",
		selectString, p.TableName, whereString, offset, limit)

	rows, err := Conn.Query(sql)

	result := []SelectResult{}

	if err != nil {
		return result, err
	}

	selectField := make([]interface{}, len(s))

	var i = 0
	for _, v := range s {
		selectField[i] = v
		i++
	}

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return result, err
		}

		var i = 0
		var tmpResult = SelectResult{}

		for k, v := range s {
			var ref = reflect.ValueOf(v)
			var refv = ref.Elem()

			if refv.Kind() == reflect.Int64 {
				tmpResult[k] = refv.Int()
			} else if refv.Kind() == reflect.String {
				tmpResult[k] = refv.String()
			}

			i++
		}

		result = append(result, tmpResult)
	}

	return result, nil
}

// MergeSelect  合并select条件
func (s SelectValues) MergeSelect() string {
	value := []string{}
	for k := range s {
		v := fmt.Sprintf("`%s`", k)
		value = append(value, v)
	}

	return strings.Join(value, ", ")
}

// MergeWhere 合并where条件
func (w WhereValues) MergeWhere() string {
	where := []string{}
	if len(w) == 0 {
		return "1=1"
	}

	for k, v := range w {
		if v.Operator == "" {
			s := fmt.Sprintf("%s = %s", k, v.Value)
			where = append(where, s)
		} else {
			s := fmt.Sprintf("%s %s %s", k, v.Operator, v.Value)
			where = append(where, s)
		}
	}

	return strings.Join(where, " AND ")
}
