package models

type (
	// FilteringQueryParams represents payload as LHS bracket for querying query params
	FilteringQueryParams struct {
		Eq   string
		Like string
		Gte  interface{}
		Lte  interface{}
	}

	// PagingQueryParams represents payload as LHS bracket for paging query params
	PagingQueryParams struct {
		Page  interface{}
		Limit interface{}
	}

	// SortQueryParams represents payload as LHS bracket for ordering/sorting query params
	SortQueryParams struct {
		By   string
		Mode string
	}
)
