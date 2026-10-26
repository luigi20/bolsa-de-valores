package entity

type OrderQueue struct {
	orders []*Order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

func (oq *OrderQueue) Less(i, j int) bool {
	return oq.orders[i].GetPrice() < oq.orders[j].GetPrice()
}

func (oq *OrderQueue) Swap(i, j int) {
	oq.orders[i], oq.orders[j] = oq.orders[j], oq.orders[i]
}

func (oq *OrderQueue) Len() int {
	return len(oq.orders)
}

func (oq *OrderQueue) Push(order any) {
	oq.orders = append(oq.orders, order.(*Order))
}

func (oq *OrderQueue) Pop() any {
	if len(oq.orders) == 0 {
		return nil
	}
	order := oq.orders[len(oq.orders)-1]
	oq.orders = oq.orders[:len(oq.orders)-1]
	return order
}
