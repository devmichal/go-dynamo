package response

type Product struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Status    int     `json:"status"`
	CreatedAt string  `json:"createdAt"`
}
