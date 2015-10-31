package cachet

// Meta can contain extra information such as the current paging attributes or links to other pages.
// Some responses will contain a meta object.
//
// Docs: https://docs.cachethq.io/docs/meta
type Meta struct {
	Pagination Pagination `json:"pagination"`
}

// Pagination will contain detail information about pages.
//
// Docs: https://docs.cachethq.io/docs/meta
type Pagination struct {
	Total       int   `json:"total"`
	Count       int   `json:"count"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Links       Links `json:"links"`
}

// Links will contain urls about the next and previous page.
//
// Docs: https://docs.cachethq.io/docs/meta
type Links struct {
	NextPage     string `json:"next_page"`
	PreviousPage string `json:"previous_page"`
}
