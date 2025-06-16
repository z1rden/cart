package cart_storage

type Cart struct {
	items map[int64]uint16 // Map: ключ - итем, uint16 - количество.
}
