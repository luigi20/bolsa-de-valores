package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	id           string
	sellingOrder *Order
	buyingOrder  *Order
	shares       int
	price        float64
	total        float64
	dateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		id:           uuid.NewString(),
		sellingOrder: sellingOrder,
		buyingOrder:  buyingOrder,
		shares:       shares,
		price:        price,
		total:        total,
		dateTime:     time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int, price float64) {
	t.total = float64(t.shares) * t.price
}

func (t *Transaction) AddBuyOrderPendingShares(shares int) {
	t.buyingOrder.SetPendingShares(t.buyingOrder.GetPendingShares() + shares)
}

func (t *Transaction) AddSellOrderPendingShares(shares int) {
	t.sellingOrder.SetPendingShares(t.sellingOrder.GetPendingShares() + shares)
}

func (t *Transaction) CloseBuyOrder() {
	if t.buyingOrder.GetPendingShares() == 0 {
		t.buyingOrder.SetStatus("CLOSED")
	}
}

func (t *Transaction) CloseSellOrder() {
	if t.sellingOrder.GetPendingShares() == 0 {
		t.sellingOrder.SetStatus("CLOSED")
	}
}

// ==================== GETTERS ====================

func (t *Transaction) GetID() string {
	return t.id
}

func (t *Transaction) GetSellingOrder() *Order {
	return t.sellingOrder
}

func (t *Transaction) GetBuyingOrder() *Order {
	return t.buyingOrder
}

func (t *Transaction) GetShares() int {
	return t.shares
}

func (t *Transaction) GetPrice() float64 {
	return t.price
}

func (t *Transaction) GetTotal() float64 {
	return t.total
}

func (t *Transaction) GetDateTime() time.Time {
	return t.dateTime
}

// ==================== SETTERS ====================

func (t *Transaction) SetID(id string) {
	t.id = id
}

func (t *Transaction) SetSellingOrder(sellingOrder *Order) {
	t.sellingOrder = sellingOrder
}

func (t *Transaction) SetBuyingOrder(buyingOrder *Order) {
	t.buyingOrder = buyingOrder
}

func (t *Transaction) SetShares(shares int) {
	t.shares = shares
}

func (t *Transaction) SetPrice(price float64) {
	t.price = price
}

func (t *Transaction) SetTotal(total float64) {
	t.total = total
}

func (t *Transaction) SetDateTime(dateTime time.Time) {
	t.dateTime = dateTime
}
