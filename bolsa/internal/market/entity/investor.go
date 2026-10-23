package entity

func NewInvestor(id string) *Investor {
	return &Investor{
		id:             id,
		asset_position: []*InvestorAssetPosition{},
	}
}

type Investor struct {
	id             string
	name           string
	asset_position []*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	asset_id string
	shares   int
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.asset_position = append(i.asset_position, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetId string, qtdShares int) {
	asset_position := i.GetAssetPosition(assetId)
	if asset_position == nil {
		i.asset_position = append(i.asset_position, NewInvestorAssetPosition(assetId, qtdShares))
	} else {
		asset_position.shares += qtdShares
	}
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		asset_id: assetId,
		shares:   shares,
	}
}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetPosition := range i.asset_position {
		if assetPosition.asset_id == assetId {
			return assetPosition
		}
	}
	return nil
}
