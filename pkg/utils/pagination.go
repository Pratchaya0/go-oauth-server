package utils

// Token based pagination
// Option can use Offset based pagination if need

type (
	PaginateRequest struct {
		Page          int `query:"page" validate:"max=64"`
		RecordPerPage int `query:"recordPerPage" validate:"required,min=2,max=10"`
	}

	PaginateResponse[T any] struct {
		Data          *T    `json:"data"`
		Page          int   `json:"page"`
		RecordPerPage int64 `json:"recordPerPage"`
		TotalPage     int64 `json:"totalPage"`
	}
)
