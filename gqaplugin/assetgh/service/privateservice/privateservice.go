package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

func GetAssetList(getAssetList model.RequestAssetList, username string) (err error, asset []model.GqaPluginAssetGh, total int64) {
	pageSize := getAssetList.PageSize
	offset := getAssetList.PageSize * (getAssetList.Page - 1)
	var assetList []model.GqaPluginAssetGh
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
		return err, assetList, 0
	}
	//配置搜索
	if getAssetList.AssetCode != "" {
		db = db.Where("asset_code = ?", getAssetList.AssetCode)
	}
	if getAssetList.AssetName != "" {
		db = db.Where("asset_name like ?", "%"+getAssetList.AssetName+"%")
	}
	if getAssetList.AssetCatalog1 != "" {
		db = db.Where("asset_catalog1 = ?", getAssetList.AssetCatalog1)
	}
	if getAssetList.AssetCatalog2 != "" {
		db = db.Where("asset_catalog2 = ?", getAssetList.AssetCatalog2)
	}
	if getAssetList.AssetCatalog3 != "" {
		db = db.Where("asset_catalog3 = ?", getAssetList.AssetCatalog3)
	}
	if getAssetList.AssetCatalog4 != "" {
		db = db.Where("asset_catalog4 = ?", getAssetList.AssetCatalog4)
	}
	if getAssetList.UseStatus != "" {
		db = db.Where("use_status = ?", getAssetList.UseStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, assetList, 0
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(getAssetList.SortBy, getAssetList.Desc)).Preload("CreatedByUser").Find(&assetList).Error
	return err, assetList, total
}

func EditAsset(toEditAsset model.GqaPluginAssetGh, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
		return err
	}
	var asset model.GqaPluginAssetGh
	if err = db.Where("id = ?", toEditAsset.Id).First(&asset).Error; err != nil {
		return err
	}
	var originalValue = toEditAsset.OriginalValue
	var usefulLife = toEditAsset.UsefulLife
	var number = toEditAsset.Number
	//计算月计提折旧
	toEditAsset.DepreciationMonth, _ = decimal.NewFromFloat(originalValue * float64(number) / float64(usefulLife*12)).Round(2).Float64()
	toEditAsset.DepreciationMonthLast, _ = decimal.NewFromFloat(originalValue*float64(number) - toEditAsset.DepreciationMonth*float64(usefulLife*12-1)).Round(2).Float64()
	//err = gqaGlobal.GqaDb.Updates(&toEditAsset).Error
	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEditAsset.GqaModel), map[string]interface{}{
		"asset_code":              toEditAsset.AssetCode,
		"asset_name":              toEditAsset.AssetName,
		"asset_catalog1":          toEditAsset.AssetCatalog1,
		"asset_catalog2":          toEditAsset.AssetCatalog2,
		"asset_catalog3":          toEditAsset.AssetCatalog3,
		"asset_catalog4":          toEditAsset.AssetCatalog4,
		"entry_date":              toEditAsset.EntryDate,
		"scrap_date":              toEditAsset.ScrapDate,
		"number":                  toEditAsset.Number,
		"original_value":          toEditAsset.OriginalValue,
		"depreciation_month":      toEditAsset.DepreciationMonth,
		"depreciation_month_last": toEditAsset.DepreciationMonthLast,
		"useful_life":             toEditAsset.UsefulLife,
		"user_dept":               toEditAsset.UserDept,
		"storage_location":        toEditAsset.StorageLocation,
		"custodian":               toEditAsset.Custodian,
		"use_status":              toEditAsset.UseStatus,
	})).Error
	return err
}

func AddAsset(toAddAsset model.GqaPluginAssetGh, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
		return err
	}
	var originalValue = toAddAsset.OriginalValue
	var usefulLife = toAddAsset.UsefulLife
	var number = toAddAsset.Number
	//计算月计提折旧
	toAddAsset.DepreciationMonth, _ = decimal.NewFromFloat(originalValue * float64(number) / float64(usefulLife*12)).Round(2).Float64()
	toAddAsset.DepreciationMonthLast, _ = decimal.NewFromFloat(originalValue*float64(number) - toAddAsset.DepreciationMonth*float64(usefulLife*12-1)).Round(2).Float64()
	err = db.Create(&toAddAsset).Error
	return err
}

func DeleteAssetById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
		return err
	}
	var asset model.GqaPluginAssetGh
	if err = db.Where("id = ?", id).First(&asset).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&asset).Error
	return err
}

func QueryAssetById(id uint, username string) (err error, assetInfo model.GqaPluginAssetGh) {
	var asset model.GqaPluginAssetGh
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
		return err, asset
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&asset, "id = ?", id).Error
	return err, asset
}

func GetSettlementList(getSettlementList model.RequestAssetSettlementList, username string) (err error, asset []model.GqaPluginAssetGhSettlement) {
	var settleList []model.GqaPluginAssetGhSettlement
	err = gqaGlobal.GqaDb.Where("set_year_month like ?", "%"+getSettlementList.SetYearMonth+"%").Find(&settleList).Error
	if err != nil {
		return err, nil
	}
	if len(settleList) != 0 || getSettlementList.SetYearMonth != time.Now().Format("200601") {
		return nil, settleList
	} else {
		var assetList []model.GqaPluginAssetGh
		var db *gorm.DB
		if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
			return err, nil
		}
		db = db.Where("use_status = ?", "assetGh_on_use")
		if err = db.Find(&assetList).Error; err != nil {
			return err, nil
		}
		var OriginalValue float64
		var MonthDepreciation float64
		var TotalDepreciation float64
		var NetWorth float64
		for _, v := range assetList {
			now := time.Now()
			entryDate, _ := time.Parse("2006-01-02", v.EntryDate)
			OriginalValue += v.OriginalValue * float64(v.Number)
			if v.ScrapDate[0:len(v.ScrapDate)-3] == now.Format("2006-01") {
				MonthDepreciation += v.DepreciationMonthLast
				TotalDepreciation += v.OriginalValue * float64(v.Number)
				NetWorth += 0
			} else {
				MonthDepreciation += v.DepreciationMonth
				TotalDepreciation += v.DepreciationMonth * float64(SubMonth(now, entryDate))
				NetWorth += v.OriginalValue*float64(v.Number) - v.DepreciationMonth
			}
		}
		err, dict := gqaUtils.GetDict("AssetGhCatalog")
		if err != nil {
			return err, nil
		}
		var assetMap = make(map[string][]model.GqaPluginAssetGh)
		for _, v := range assetList {
			for _, d := range dict {
				if v.AssetCatalog1 == d.DictCode {
					assetMap[d.DictLabel] = append(assetMap[d.DictLabel], v)
				}
			}
		}
		var assetSettlementList []model.GqaPluginAssetGhSettlement
		for k, v := range assetMap {
			var ov float64
			var md float64
			var td float64
			var nw float64
			for _, vv := range v {
				now := time.Now()
				entryDate, _ := time.Parse("2006-01-02", vv.EntryDate)
				if vv.ScrapDate[0:len(vv.ScrapDate)-3] == now.Format("2006-01") {
					ov += vv.OriginalValue * float64(vv.Number)
					md += vv.DepreciationMonthLast
					td += vv.OriginalValue * float64(vv.Number)
					nw += 0
				} else {
					ov += vv.OriginalValue * float64(vv.Number)
					md += vv.DepreciationMonth
					td += vv.DepreciationMonth * float64(SubMonth(now, entryDate))
					nw += vv.OriginalValue*float64(vv.Number) - vv.DepreciationMonth
				}
			}
			assetSettlementList = append(assetSettlementList, model.GqaPluginAssetGhSettlement{
				SetYearMonth:      getSettlementList.SetYearMonth,
				AssetCatalog1:     k,
				OriginalValue:     ov,
				MonthDepreciation: md,
				TotalDepreciation: td,
				NetWorth:          nw,
			})
		}
		assetSettlementList = append(assetSettlementList, model.GqaPluginAssetGhSettlement{
			SetYearMonth:      getSettlementList.SetYearMonth,
			AssetCatalog1:     "合计",
			OriginalValue:     OriginalValue,
			MonthDepreciation: MonthDepreciation,
			TotalDepreciation: TotalDepreciation,
			NetWorth:          NetWorth,
		})
		if len(assetSettlementList) != 0 {
			err = gqaGlobal.GqaDb.Create(&assetSettlementList).Error
			return err, assetSettlementList
		}
		return err, nil
	}
}

func SetSettlement(username string) (err error) {
	//先删除月结表中当月的数据
	err = gqaGlobal.GqaDb.Where("set_year_month = ?", time.Now().Format("200601")).Unscoped().Delete(&model.GqaPluginAssetGhSettlement{}).Error
	if err != nil {
		return err
	}
	var assetList []model.GqaPluginAssetGh
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPluginAssetGh{})); err != nil {
		return err
	}
	db = db.Where("use_status = ?", "assetGh_on_use")
	if err = db.Find(&assetList).Error; err != nil {
		return err
	}
	var newAssetList []model.GqaPluginAssetGh
	for _, ass := range assetList {
		ScrapDate, _ := time.Parse("2006-01-02", ass.ScrapDate)
		if ScrapDate.Before(time.Now()) {
			ass.UseStatus = "assetGh_scrap"
		}
		newAssetList = append(newAssetList, ass)
	}
	if err = db.Save(&newAssetList).Error; err != nil {
		return err
	}
	return nil
}

// 计算日期相差多少月
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return month
}
