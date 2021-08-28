package models

type (
	FilteringQueryParams struct {
		Eq   interface{}
		Like interface{}
		Gte  interface{}
		Lte  interface{}
	}

	PagingQueryParams struct {
		Page  interface{}
		Limit interface{}
	}

	SortQueryParams struct {
		By   interface{}
		Mode interface{}
	}
)
