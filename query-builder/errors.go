package querybuilder

const (
	ErrInvalidQueryBuilder   = "invalid aggregator query builder"
	ErrTwiceSelectInRow      = "cannot call select clause twice in a row"
	ErrEmptyAggregationField = "cannot make a aggregate operation with empty fields"
)
