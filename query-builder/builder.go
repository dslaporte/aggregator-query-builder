package querybuilder

import (
	"errors"
	"fmt"
	"strings"
)

type aggregationFn string

const (
	fnSUM   aggregationFn = "SUM"
	fnAVG   aggregationFn = "AVG"
	fnCOUNT aggregationFn = "COUNT"
	fnMIN   aggregationFn = "MIN"
	fnMAX   aggregationFn = "MAX"
)

type aggregatorQueryBuilder struct {
	selectClause    string
	selectFields    []string
	tableName       string
	aggregateFields []string
}

// NewAggregatorQueryBuilder is a constructor to start a new QueryBuilder
// Use the method Build() to create a complete SQL syntax.
func NewAggregatorQueryBuilder() *aggregatorQueryBuilder {
	return &aggregatorQueryBuilder{
		selectClause:    "",
		selectFields:    []string{},
		tableName:       "",
		aggregateFields: []string{},
	}
}

// isValid() should be called on every method to check
// if the aggregatorQueryBuilder pointer is not nil
func (q *aggregatorQueryBuilder) validate() {
	if q == nil {
		panic(errors.New(ErrInvalidQueryBuilder))
	}
}

// Select adds "SELECT" clause on internal SQL Builder
func (q *aggregatorQueryBuilder) Select() *aggregatorQueryBuilder {
	q.validate()

	if !IsEmpty(q.selectClause) {
		panic(errors.New(ErrTwiceSelectInRow))
	}
	q.selectClause = "SELECT"
	return q
}

/*
		Fields is a list of string names.
		Can be used with alias like "tb.FieldName as Name".
		When fields are not informed, will be considered the character "*"

		 queryBuilder := NewaggregatorQueryBuilder()
		 fmt.Println(queryBuilder.Select().From("TESTE").Build())
	 	 "SELECT * FROM TESTE"

	 	 fmt.Println(queryBuilder.Select("FIELD1", "FIELD2", "FIELD3").From("TESTE").Build())
	 	 "SELECT FIELD1,FIELD2,FIELD3 FROM TESTE"
*/
func (q *aggregatorQueryBuilder) Fields(fields ...string) *aggregatorQueryBuilder {
	q.validate()
	q.selectFields = append(q.selectFields, fields...)
	return q
}

/*
From adds a Table used in select sentence. E.g.:

	queryBuilder := NewaggregatorQueryBuilder()
	fmt.Println(queryBuilder.Select().From("TESTE").Build())

	"SELECT * FROM TESTE"
*/
func (q *aggregatorQueryBuilder) From(tableName string) *aggregatorQueryBuilder {
	q.validate()
	q.tableName = tableName
	return q
}

/*
Build is used to "compile" all the internal called methods,
resulting in a SQL sentence. E.g.:

	queryBuilder := NewaggregatorQueryBuilder()
	fmt.Println(queryBuilder.Select().From("TESTE").Build())

Expected result:

	"SELECT * FROM TESTE"
*/
func (q *aggregatorQueryBuilder) Build() string {
	fields := "*"
	groupedFields := ""
	if len(q.selectFields) > 0 {
		fields = strings.Join(q.selectFields, ",")
	}
	if len(q.aggregateFields) > 0 {
		groupedFields = fmt.Sprintf("%s %s", " GROUP BY ", strings.Join(q.selectFields, ","))
		if IsEmpty(fields) {
			fields = strings.Join(q.aggregateFields, ",")
		}
		fields += "," + strings.Join(q.aggregateFields, ",")
	}

	result := []string{
		q.selectClause,
		fields,
		"FROM",
		q.tableName,
		groupedFields,
	}
	return strings.TrimSpace(strings.Join(result, " "))
}

// addAggregateFn is an internal method that will be called for each aggregation method
// centralizing all the logic
func (q *aggregatorQueryBuilder) addAggregateFn(aggFunction aggregationFn, fields ...string) *aggregatorQueryBuilder {
	q.validate()
	if len(fields) == 0 {
		panic(errors.New(ErrEmptyAggregationField))
	}
	fieldSintax := ""
	switch aggFunction {
	case fnSUM:
		fieldSintax = "SUM(%s) as %s"
	case fnAVG:
		fieldSintax = "AVG(%s) as %s"
	case fnCOUNT:
		fieldSintax = "COUNT(%s) as %s"
	case fnMAX:
		fieldSintax = "MAX(%s) as %s"
	case fnMIN:
		fieldSintax = "MIN(%s) as %s"
	}
	for _, field := range fields {
		q.aggregateFields = append(q.aggregateFields, fmt.Sprintf(fieldSintax, field, field))
	}
	return q
}

// Sum allow you aggregate lot of fields using SQL SUM clause
//
// If you use Sum method with common query fields, it will
// automatically set the fields to group by clause
func (q *aggregatorQueryBuilder) Sum(fields ...string) *aggregatorQueryBuilder {
	return q.addAggregateFn(fnSUM, fields...)
}

// AVG allow you aggregate lot of fields using SQL AVG clause
//
// If you use Avg method with common query fields, it will
// automatically set the fields to group by clause
func (q *aggregatorQueryBuilder) Avg(fields ...string) *aggregatorQueryBuilder {
	return q.addAggregateFn(fnAVG, fields...)
}

// Count allow you aggregate lot of fields using SQL COUNT clause
//
// If you use Count method with common query fields, it will
// automatically set the fields to group by clause
func (q *aggregatorQueryBuilder) Count(fields ...string) *aggregatorQueryBuilder {
	return q.addAggregateFn(fnCOUNT, fields...)
}

// Min allow you aggregate lot of fields using SQL MIN clause
//
// If you use Min method with common query fields, it will
// automatically set the fields to group by clause
func (q *aggregatorQueryBuilder) Min(fields ...string) *aggregatorQueryBuilder {
	return q.addAggregateFn(fnMIN, fields...)
}

// Max allow you aggregate lot of fields using SQL MAX clause
//
// If you use Max method with common query fields, it will
// automatically set the fields to group by clause
func (q *aggregatorQueryBuilder) Max(fields ...string) *aggregatorQueryBuilder {
	return q.addAggregateFn(fnMAX, fields...)
}

func (q *aggregatorQueryBuilder) SubQuery() *aggregatorQueryBuilder {
	//to be implemented
	return q
}

func (q *aggregatorQueryBuilder) OptionalParam() *aggregatorQueryBuilder {
	//to be implemented
	return q
}

func (q *aggregatorQueryBuilder) Cast(field string, dataType string) string {
	//TODO: MAKE TESTS
	return fmt.Sprintf("CAST( %s as %s ) as %s", field, dataType, field)
}
