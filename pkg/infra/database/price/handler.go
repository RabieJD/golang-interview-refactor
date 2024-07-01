package price

type Handler interface {
	GetPrice(identifier string) (float64, error)
}
