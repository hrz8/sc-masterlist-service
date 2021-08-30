package models

type (
	FilteringQueryParams struct {
		Eq   string
		Like string
		Gte  interface{}
		Lte  interface{}
	}

	PagingQueryParams struct {
		Page  interface{}
		Limit interface{}
	}

	SortQueryParams struct {
		By   string
		Mode string
	}
)
