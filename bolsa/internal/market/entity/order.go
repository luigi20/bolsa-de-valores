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

// ==================== GETTERS ====================

func (o *Order) GetID() string {
	return o.id
}

func (o *Order) GetInvestor() *Investor {
	return o.investor
}

func (o *Order) GetAsset() *Asset {
	return o.asset
}

func (o *Order) GetShares() int {
	return o.shares
}

func (o *Order) GetPendingShares() int {
	return o.pendingShares
}

func (o *Order) GetPrice() float64 {
	return o.price
}

func (o *Order) GetOrderType() string {
	return o.orderType
}

func (o *Order) GetStatus() string {
	return o.status
}

func (o *Order) GetTransactions() []*Transaction {
	return o.transactions
}

// ==================== SETTERS ====================

func (o *Order) SetID(id string) {
	o.id = id
}

func (o *Order) SetInvestor(investor *Investor) {
	o.investor = investor
}

func (o *Order) SetAsset(asset *Asset) {
	o.asset = asset
}

func (o *Order) SetShares(shares int) {
	o.shares = shares
}

func (o *Order) SetPendingShares(pendingShares int) {
	o.pendingShares = pendingShares
}

func (o *Order) SetPrice(price float64) {
	o.price = price
}

func (o *Order) SetOrderType(orderType string) {
	o.orderType = orderType
}

func (o *Order) SetStatus(status string) {
	o.status = status
}

func (o *Order) SetTransactions(transactions []*Transaction) {
	o.transactions = transactions
}

// Opcional: método auxiliar muito útil para adicionar uma transação individualmente ao slice
func (o *Order) AddTransaction(transaction *Transaction) {
	o.transactions = append(o.transactions, transaction)
}
