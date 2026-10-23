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
