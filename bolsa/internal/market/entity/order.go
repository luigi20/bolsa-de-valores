package entity

type Order struct {
	id            string
	investor      *Investor
	asset         *Asset
	shares        int
	pendingShares int
	price         float64
	orderType     string
	status        string
	transactions  []*Transaction
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		id:            id,
		investor:      investor,
		asset:         asset,
		shares:        shares,
		pendingShares: shares,
		price:         price,
		orderType:     orderType,
		status:        "open",
		transactions:  []*Transaction{},
	}
}
