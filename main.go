package main

import (
	querybuilder "aggregator-query-builder/query-builder"
	"fmt"
)

func main() {
	querybuilder := querybuilder.NewAggregatorQueryBuilder()
	fmt.Println(`printing a "select * from teste" clause:`)
	fmt.Println(querybuilder.Select().From("TESTE").Build())
}
