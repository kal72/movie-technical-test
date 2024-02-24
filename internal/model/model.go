package model

type Response[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type PageResponse[T any] struct {
	Status       string        `json:"status"`
	Data         []T           `json:"data"`
	PageMetadata *PageMetadata `json:"paging,omitempty"`
	Message      string        `json:"message"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int   `json:"total_page"`
}
