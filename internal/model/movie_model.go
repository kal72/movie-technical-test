package model

type MovieRequest struct {
	ID          int     `json:"-"`
	Title       string  `json:"title" validate:"required,max=255"`
	Description string  `json:"description" validate:"required,max=255"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}

type MovieResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
