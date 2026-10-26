package entity

type Asset struct {
	id            string
	name          string
	market_volume int
}

func NewAsset(id string, name string, marketVolume int) *Asset {
	return &Asset{
		id:            id,
		name:          name,
		market_volume: marketVolume,
	}
}

// ==================== GETTERS ====================

func (a *Asset) GetID() string {
	return a.id
}

func (a *Asset) GetName() string {
	return a.name
}

func (a *Asset) GetMarketVolume() int {
	return a.market_volume
}

// ==================== SETTERS ====================

func (a *Asset) SetID(id string) {
	a.id = id
}

func (a *Asset) SetName(name string) {
	a.name = name
}

func (a *Asset) SetMarketVolume(marketVolume int) {
	a.market_volume = marketVolume
}
