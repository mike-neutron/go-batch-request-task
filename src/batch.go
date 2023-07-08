package src

// Batch is a batch of items.
type Batch []Item

// Item is some abstract item.
type Item struct{
	Product Product
	Stock int32
}

// Product
type Product struct {
	Name string
}