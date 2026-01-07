package request

type CreateProduct struct {
	Name   string `json:"name" validate:"required,min=3,max=255"`
	Price  int64  `json:"price" validate:"required,gt=0"`
	Active *bool  `json:"active" validate:"required"`
}

type CreateProductOutput struct {
	Id string
}
