package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	order         []*Order
	transactions  []*Transaction
	ordersChan    chan *Order
	ordersChanOut chan *Order
	wg            *sync.WaitGroup
}

func NewBook(ordersChan chan *Order, ordersChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		order:         []*Order{},
		transactions:  []*Transaction{},
		ordersChan:    ordersChan,
		ordersChanOut: ordersChanOut,
		wg:            wg,
	}
}

func (b *Book) Trade() {
	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)
	for order := range b.ordersChan {
		asset := order.GetAsset().GetID()
		if buyOrders[asset] == nil {
			buyOrders[asset] = NewOrderQueue()
			heap.Init(buyOrders[asset])
		}
		if sellOrders[asset] == nil {
			sellOrders[asset] = NewOrderQueue()
			heap.Init(sellOrders[asset])
		}

		if order.GetOrderType() == "BUY" {
			buyOrders[asset].Push(order)
			if sellOrders[asset].Len() > 0 && sellOrders[asset].orders[0].GetPrice() <= order.GetPrice() {
				sellOrder := sellOrders[asset].Pop().(*Order)
				if sellOrder.GetPendingShares() > 0 {
					transaction := NewTransaction(sellOrder, order, order.GetShares(), sellOrder.GetPrice())
					b.AddTransaction(transaction, b.wg)
					sellOrder.SetTransactions(append(sellOrder.GetTransactions(), transaction))
					order.SetTransactions(append(order.GetTransactions(), transaction))
					b.ordersChanOut <- sellOrder
					b.ordersChanOut <- order
					if sellOrder.GetPendingShares() > 0 {
						sellOrders[asset].Push(sellOrder)
					}
				}
			}
		} else if order.GetOrderType() == "SELL" {
			sellOrders[asset].Push(order)
			if buyOrders[asset].Len() > 0 && buyOrders[asset].orders[0].GetPrice() >= order.GetPrice() {
				buyOrder := buyOrders[asset].Pop().(*Order)
				if buyOrder.GetPendingShares() > 0 {
					transaction := NewTransaction(order, buyOrder, order.GetShares(), buyOrder.GetPrice())
					b.AddTransaction(transaction, b.wg)
					buyOrder.SetTransactions(append(buyOrder.GetTransactions(), transaction))
					order.SetTransactions(append(order.GetTransactions(), transaction))
					b.ordersChanOut <- buyOrder
					b.ordersChanOut <- order
					if buyOrder.GetPendingShares() > 0 {
						buyOrders[asset].Push(buyOrder)
					}
				}
			}
		}
	}
}

func (b *Book) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	sellingShares := transaction.sellingOrder.GetPendingShares()
	buyingShares := transaction.buyingOrder.GetPendingShares()
	minShares := sellingShares
	if buyingShares < sellingShares {
		minShares = buyingShares
	}
	transaction.sellingOrder.GetInvestor().UpdateAssetPosition(transaction.sellingOrder.GetAsset().GetID(), -minShares)
	transaction.AddSellOrderPendingShares(-minShares)
	transaction.buyingOrder.GetInvestor().UpdateAssetPosition(transaction.buyingOrder.GetAsset().GetID(), minShares)
	transaction.AddBuyOrderPendingShares(-minShares)
	transaction.CalculateTotal(transaction.shares, transaction.buyingOrder.GetPrice())
	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()
	b.transactions = append(b.transactions, transaction)
}
