package database

type Car struct {
	Name     string `json:"name"`
	Price    uint64 `json:"price"`
	Discount uint64 `json:"discount"`
}
