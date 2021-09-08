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

	// DeleteQueryParams represents payload as LHS bracket for deleted rows selection
	DeleteQueryParams struct {
		Include bool
		Only    bool
	}

	// ManyToManyQueryParams represents payload as LHS bracket for many to many relation
	ManyToManyQueryParams struct {
		Eq string
		In string
	}
)
