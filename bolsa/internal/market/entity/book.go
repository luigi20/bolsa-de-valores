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
	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()
	heap.Init(buyOrders)
	heap.Init(sellOrders)
	for order := range b.ordersChan {
		{
			if order.orderType == "BUY" {
				buyOrders.Push(order)
				if sellOrders.Len() > 0 && sellOrders.orders[0].price <= order.price {
					sellOrder := sellOrders.Pop().(*Order)
					if sellOrder.pendingShares > 0 {
						transaction := NewTransaction(sellOrder, order, order.shares, sellOrder.price)
						b.AddTransaction(transaction, b.wg)
						sellOrder.transactions = append(sellOrder.transactions, transaction)
						order.transactions = append(order.transactions, transaction)
						b.ordersChanOut <- sellOrder
						b.ordersChanOut <- order
						if sellOrder.pendingShares > 0 {
							sellOrders.Push(sellOrder)
						}
					}
				}
			} else if order.orderType == "SELL" {
				sellOrders.Push(order)
				if buyOrders.Len() > 0 && buyOrders.orders[0].price >= order.price {
					buyOrder := buyOrders.Pop().(*Order)
					if buyOrder.pendingShares > 0 {
						transaction := NewTransaction(order, buyOrder, order.shares, buyOrder.price)
						b.AddTransaction(transaction, b.wg)
						buyOrder.transactions = append(buyOrder.transactions, transaction)
						order.transactions = append(order.transactions, transaction)
						b.ordersChanOut <- buyOrder
						b.ordersChanOut <- order
						if buyOrder.pendingShares > 0 {
							buyOrders.Push(buyOrder)
						}
					}
				}
			}
		}
	}
}

func (b *Book) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	sellingShares := transaction.sellingOrder.pendingShares
	buyingShares := transaction.buyingOrder.pendingShares
	minShares := sellingShares
	if buyingShares < sellingShares {
		minShares = buyingShares
	}
	transaction.sellingOrder.investor.UpdateAssetPosition(transaction.sellingOrder.asset.id, -minShares)
	transaction.AddSellOrderPendingShares(-minShares)
	transaction.buyingOrder.investor.UpdateAssetPosition(transaction.buyingOrder.asset.id, minShares)
	transaction.AddBuyOrderPendingShares(-minShares)
	transaction.CalculateTotal(transaction.shares, transaction.buyingOrder.price)
	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()
	b.transactions = append(b.transactions, transaction)
}
