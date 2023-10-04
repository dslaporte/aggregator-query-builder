package querybuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregatorQueryBuilder(t *testing.T) {
	t.Run("should return success when select is called", func(t *testing.T) {
		query := NewAggregatorQueryBuilder()
		assert.Equal(t, "SELECT * FROM", query.Select().Build())
	})

	t.Run("should return an error when select is called", func(t *testing.T) {
		expectedError := "cannot call select clause twice in a row"
		query := NewAggregatorQueryBuilder()
		assert.PanicsWithError(t, expectedError, func() {
			query.Select().Select().Build()
		})
	})

	t.Run("should return a complete sql sentence when method build is called", func(t *testing.T) {
		query := NewAggregatorQueryBuilder()
		expectedResult := "SELECT CAMPO_A,CAMPO_B,CAMPO_C FROM TableName"
		assert.Equal(t, expectedResult, query.Select().Fields("CAMPO_A", "CAMPO_B", "CAMPO_C").From("TableName").Build())
	})
}
