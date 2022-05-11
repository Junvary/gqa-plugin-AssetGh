package privateapi

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/service/privateservice"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAssetList(c *gin.Context) {
	var getAssetList model.RequestAssetList
	if err := gqaModel.RequestShouldBindJSON(c, &getAssetList); err != nil {
		return
	}
	if err, asset, total := privateservice.GetAssetList(getAssetList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取固定资产列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取固定资产列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records:  asset,
			Page:     getAssetList.Page,
			PageSize: getAssetList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditAsset(c *gin.Context) {
	var toEditAsset model.GqaPluginAssetGh
	if err := gqaModel.RequestShouldBindJSON(c, &toEditAsset); err != nil {
		return
	}
	toEditAsset.UpdatedBy = gqaUtils.GetUsername(c)
	if err := privateservice.EditAsset(toEditAsset, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("编辑固定资产失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("编辑固定资产失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData("编辑固定资产成功！", c)
	}
}

func AddAsset(c *gin.Context) {
	var toAddAsset model.RequestAddAsset
	if err := gqaModel.RequestShouldBindJSON(c, &toAddAsset); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = gqaModel.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: gqaGlobal.GqaModel{
			CreatedBy: gqaUtils.GetUsername(c),
			Status:    toAddAsset.Status,
			Sort:      toAddAsset.Sort,
			Memo:      toAddAsset.Memo,
		},
	}
	addAsset := &model.GqaPluginAssetGh{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		AssetCode:                         toAddAsset.AssetCode,
		AssetName:                         toAddAsset.AssetName,
		AssetCatalog1:                     toAddAsset.AssetCatalog1,
		AssetCatalog2:                     toAddAsset.AssetCatalog2,
		AssetCatalog3:                     toAddAsset.AssetCatalog3,
		AssetCatalog4:                     toAddAsset.AssetCatalog4,
		EntryDate:                         toAddAsset.EntryDate,
		Number:                            toAddAsset.Number,
		OriginalValue:                     toAddAsset.OriginalValue,
		UsefulLife:                        toAddAsset.UsefulLife,
		UserDept:                          toAddAsset.UserDept,
		StorageLocation:                   toAddAsset.StorageLocation,
		Custodian:                         toAddAsset.Custodian,
		UseStatus:                         toAddAsset.UseStatus,
	}
	if err := privateservice.AddAsset(*addAsset, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("添加固定资产失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("添加固定资产失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData("添加固定资产成功！", c)
	}
}

func DeleteAssetById(c *gin.Context) {
	var toDeleteId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := privateservice.DeleteAssetById(toDeleteId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("删除固定资产失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("删除固定资产失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData("删除固定资产成功！", c)
	}
}

func QueryAssetById(c *gin.Context) {
	var toQueryId gqaModel.RequestQueryById
	if err := gqaModel.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := privateservice.QueryAssetById(toQueryId.Id, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("查找固定资产失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("查找固定资产失败，"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessMessageData(gin.H{"records": dept}, "查找固定资产成功！", c)
	}
}

func GetSettlementList(c *gin.Context) {
	var getSettlementList model.RequestAssetSettlementList
	if err := gqaModel.RequestShouldBindJSON(c, &getSettlementList); err != nil {
		return
	}
	if err, asset := privateservice.GetSettlementList(getSettlementList, gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("获取固定资产列表失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("获取固定资产列表失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData(gqaModel.ResponsePage{
			Records: asset,
		}, c)
	}
}

func SetSettlement(c *gin.Context) {
	if err := privateservice.SetSettlement(gqaUtils.GetUsername(c)); err != nil {
		gqaGlobal.GqaLogger.Error("执行月结失败！", zap.Any("err", err))
		gqaModel.ResponseErrorMessage("执行月结失败！"+err.Error(), c)
	} else {
		gqaModel.ResponseSuccessData("执行月结成功！", c)
	}
}
