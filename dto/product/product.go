package productdto

type CreateProduct struct {
	Title string `json:"title" validate:"required"`
	Brand string `json:"brand" validate:"required"`
	Image string `json:"image"`
	Price int    `json:"price" validate:"required"`
}

type UpdateProduct struct {
	Title string `json:"title"`
	Brand string `json:"brand"`
	Image string `json:"image"`
	Price int    `json:"price"`
}
