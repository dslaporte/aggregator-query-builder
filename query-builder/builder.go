package querybuilder

import (
	"errors"
	"strings"
)

type aggregatorQueryBuilder struct {
	selectClause string
	fields       []string
	tableName    string
}

// NewAggregatorQueryBuilder is a constructor to start a new QueryBuilder
// Use the method Build() to create a complete SQL syntax.
func NewAggregatorQueryBuilder() *aggregatorQueryBuilder {
	return &aggregatorQueryBuilder{
		selectClause: "",
		fields:       []string{},
		tableName:    "",
	}
}

// isValid() should be called on every method to check
// if the aggregatorQueryBuilder pointer is not nil
func (q *aggregatorQueryBuilder) isValid() error {
	if q == nil {
		panic(errors.New(ErrInvalidQueryBuilder))
	}
	return nil
}

// Select adds "SELECT" clause on internal SQL Builder
func (q *aggregatorQueryBuilder) Select() *aggregatorQueryBuilder {
	q.isValid()

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
	q.isValid()
	q.fields = append(q.fields, fields...)
	return q
}

/*
  From adds a Table used in select sentence. E.g.:

   queryBuilder := NewaggregatorQueryBuilder()
   fmt.Println(queryBuilder.Select().From("TESTE").Build())

   "SELECT * FROM TESTE"
*/
func (q *aggregatorQueryBuilder) From(tableName string) *aggregatorQueryBuilder {
	q.isValid()
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
	if len(q.fields) > 0 {
		fields = strings.Join(q.fields, ",")
	}

	result := []string{
		q.selectClause,
		fields,
		"FROM",
		q.tableName,
	}
	return strings.TrimSpace(strings.Join(result, " "))

}
