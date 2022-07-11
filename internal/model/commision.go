package model

const TableName = "commission_db"

type Commission struct {
	Id        string
	Name      string
	Status    int
	CreatedAt string
	Price     float64
}

func NewCommission(id string, name string, status int, price float64) *Commission {

	return &Commission{
		Id:        id,
		Name:      name,
		Status:    status,
		CreatedAt: "",
		Price:     price,
	}
}
