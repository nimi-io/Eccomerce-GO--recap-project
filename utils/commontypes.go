package utils

// PaginateOptions maps to IDefaultPaginationOptions
type PaginateOptions struct {
	Limit int
	Page  int
}

// Meta maps to IMeta
type Meta struct {
	TotalItems  int64 `json:"totalItems"`
	Count       int   `json:"count"`
	ItemsPerPage int   `json:"itemsPerPage"`
	CurrentPage int   `json:"currentPage"`
	TotalPages  int   `json:"totalPages"`
}

// PaginateResult maps to IPaginateResult
type PaginateResult[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}
