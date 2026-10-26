package transformer

import (
	"github.com/luigi20/bolsa_de_valores/bolsa/internal/market/dto"
	"github.com/luigi20/bolsa_de_valores/bolsa/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID)
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)
	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}
	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID:    order.GetID(),
		InvestorID: order.GetInvestor().GetID(),
		AssetID:    order.GetAsset().GetID(),
		OrderType:  order.GetOrderType(),
		Status:     order.GetStatus(),
		Partial:    order.GetPendingShares(),
		Shares:     order.GetShares(),
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, t := range order.GetTransactions() {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.GetID(),
			BuyerID:       t.GetBuyingOrder().GetInvestor().GetID(),
			SellerID:      t.GetSellingOrder().GetInvestor().GetID(),
			AssetID:       t.GetSellingOrder().GetAsset().GetID(),
			Price:         t.GetPrice(),
			Shares:        t.GetSellingOrder().GetShares() - t.GetSellingOrder().GetPendingShares(),
		}
		transactionsOutput = append(transactionsOutput, transactionOutput)
	}
	output.TransactionsOutput = transactionsOutput
	return output
}
