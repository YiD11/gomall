package types

type OrderItem struct {
	ProductName string
	Picture     string
	Quantity    uint32
	Cost        float32
}

type Order struct {
	OrderId     string
	CreatedDate string
	Cost        float32
	Items       []OrderItem
}
