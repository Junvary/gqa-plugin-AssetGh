package model

import (
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

type GqaPluginAssetGh struct {
	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	AssetCode             string  `json:"assetCode" gorm:"comment:资产编号;not null;index"`
	AssetName             string  `json:"assetName" gorm:"comment:资产名称;not null;index;"`
	AssetCatalog1         string  `json:"assetCatalog1" gorm:"comment:资产类别1;not null;"`
	AssetCatalog2         string  `json:"assetCatalog2" gorm:"comment:资产类别2;not null;"`
	AssetCatalog3         string  `json:"assetCatalog3" gorm:"comment:资产类别3;"`
	AssetCatalog4         string  `json:"assetCatalog4" gorm:"comment:资产类别4;"`
	EntryDate             string  `json:"entryDate" gorm:"comment:入账日期;not null;"`
	ScrapDate             string  `json:"scrapDate" gorm:"comment:报废日期;not null;"`
	Number                int     `json:"number" gorm:"comment:数量;not null;"`
	OriginalValue         float64 `json:"originalValue" gorm:"comment:资产原值;not null;"`
	DepreciationMonth     float64 `json:"depreciationMonth" gorm:"comment:折旧(元月);not null;"`
	DepreciationMonthLast float64 `json:"depreciationMonthLast" gorm:"comment:最后一个月折旧(元月);not null;"`
	UsefulLife            int     `json:"usefulLife" gorm:"comment:使用年限;not null;"`
	UserDept              string  `json:"userDept" gorm:"comment:使用部门;not null;"`
	StorageLocation       string  `json:"storageLocation" gorm:"comment:存放地点;not null;"`
	Custodian             string  `json:"custodian" gorm:"comment:保管人;not null;"`
	UseStatus             string  `json:"useStatus" gorm:"comment:使用状态;not null;default:assetGh_on_use;index"`
}

type RequestAddAsset struct {
	gqaModel.RequestAdd
	AssetCode       string  `json:"assetCode"`
	AssetName       string  `json:"assetName"`
	AssetCatalog1   string  `json:"assetCatalog1"`
	AssetCatalog2   string  `json:"assetCatalog2"`
	AssetCatalog3   string  `json:"assetCatalog3"`
	AssetCatalog4   string  `json:"assetCatalog4"`
	EntryDate       string  `json:"entryDate"`
	ScrapDate       string  `json:"scrapDate"`
	Number          int     `json:"number"`
	OriginalValue   float64 `json:"originalValue"`
	UsefulLife      int     `json:"usefulLife"`
	UserDept        string  `json:"userDept"`
	StorageLocation string  `json:"storageLocation"`
	Custodian       string  `json:"custodian"`
	UseStatus       string  `json:"useStatus"`
}

type RequestAssetList struct {
	gqaModel.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddAsset 中的字段
	AssetCode     string `json:"assetCode"`
	AssetName     string `json:"assetName"`
	AssetCatalog1 string `json:"assetCatalog1"`
	AssetCatalog2 string `json:"assetCatalog2"`
	AssetCatalog3 string `json:"assetCatalog3"`
	AssetCatalog4 string `json:"assetCatalog4"`
	UseStatus     string `json:"useStatus"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//GqaPluginAssetGh
}
