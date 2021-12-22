package database

import "gorm.io/gorm"

type TokenPriceModel struct {
	gorm.Model

	TokenName   string `gorm:"uniqueIndex,size=255"`
	TimeRecord  string
	RecordAt    string
	MarketPrice string
	Volume      string
	OpenPrice   string
	ClosePrice  string
	HighPrice   string
	LowPrice    string
	CandleSize  string
}

func (a TokenPriceModel) TableName() string {
	return "token_price_data"
}
