package types

type OrderItem struct {
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float32
}

type Order struct {
	OrderId   string
	CreatDate string
	Cost      float32
	Items     []OrderItem
}
