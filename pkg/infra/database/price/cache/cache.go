package cache

// we are storing the prices in the system which is not convenient
// TODO we must refactor the approach of price handling

var itemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}
