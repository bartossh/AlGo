package constrains

// Sortable constrain types to those that can preserve order based on comparable value
type Sortable interface {
	~int | ~rune | ~int64 | ~byte | ~uint | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type NaturalNumber interface {
	~int | ~rune | ~int64 | ~byte | ~uint | ~uint16 | ~uint32 | ~uint64
}
