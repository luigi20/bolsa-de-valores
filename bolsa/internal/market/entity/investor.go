package entity

type Investor struct {
	id             string
	name           string
	asset_position []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		id:             id,
		asset_position: []*InvestorAssetPosition{},
	}
}

type InvestorAssetPosition struct {
	asset_id string
	shares   int
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		asset_id: assetId,
		shares:   shares,
	}
}

// ==================== MÉTODOS ORIGINAIS DO INVESTOR ====================

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.asset_position = append(i.asset_position, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetId string, qtdShares int) {
	asset_position := i.GetAssetPosition(assetId)
	if asset_position == nil {
		i.asset_position = append(i.asset_position, NewInvestorAssetPosition(assetId, qtdShares))
	} else {
		// Como estamos dentro do mesmo pacote (entity), podemos acessar .shares diretamente,
		// mas caso queira usar setters/getters, você pode usar asset_position.SetShares(...)
		asset_position.shares += qtdShares
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

// ==================== GETTERS E SETTERS - INVESTOR ====================

func (i *Investor) GetID() string {
	return i.id
}

func (i *Investor) SetID(id string) {
	i.id = id
}

func (i *Investor) GetName() string {
	return i.name
}

func (i *Investor) SetName(name string) {
	i.name = name
}

func (i *Investor) GetAssetPositionList() []*InvestorAssetPosition {
	return i.asset_position
}

func (i *Investor) SetAssetPositionList(assetPosition []*InvestorAssetPosition) {
	i.asset_position = assetPosition
}

// ==================== GETTERS E SETTERS - INVESTOR ASSET POSITION ====================

func (iap *InvestorAssetPosition) GetAssetID() string {
	return iap.asset_id
}

func (iap *InvestorAssetPosition) SetAssetID(assetId string) {
	iap.asset_id = assetId
}

func (iap *InvestorAssetPosition) GetShares() int {
	return iap.shares
}

func (iap *InvestorAssetPosition) SetShares(shares int) {
	iap.shares = shares
}
