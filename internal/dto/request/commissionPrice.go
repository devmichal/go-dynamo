package request

type Product struct {
	Name   string  `json:"name" from:"name" validate:"required,min=3,max=32"`
	Price  float64 `json:"price" from:"price"`
	Status int     `json:"status" from:"status"`
}
