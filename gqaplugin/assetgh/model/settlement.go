package model

type GqaPluginAssetGhSettlement struct {
	SetYearMonth      string  `json:"setYearMonth" gorm:"comment:'执行年月';not null;index"`
	AssetCatalog1     string  `json:"assetCatalog1" gorm:"comment:'资产类别1';not null;index"`
	OriginalValue     float64 `json:"originalValue" gorm:"comment:'原值';not null"`
	MonthDepreciation float64 `json:"monthDepreciation" gorm:"comment:'月折旧';not null"`
	TotalDepreciation float64 `json:"totalDepreciation" gorm:"comment:'累计折旧';not null"`
	NetWorth          float64 `json:"netWorth" gorm:"comment:'净值';not null"`
}

type RequestAssetSettlementList struct {
	SetYearMonth string `json:"setYearMonth"`
}
