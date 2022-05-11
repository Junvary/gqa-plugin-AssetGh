package privaterouter

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/api/privateapi"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/middleware"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(privateGroup *gin.RouterGroup) {
	//插件路由在注册的时候被分配为 PluginCode() 分组，无须再次分组。
	assetGhRouter := privateGroup.Use(middleware.LogOperationHandler())
	{
		//获取asset列表
		assetGhRouter.POST("get-asset-list", privateapi.GetAssetList)
		//编辑asset信息
		assetGhRouter.POST("edit-asset", privateapi.EditAsset)
		//新增asset
		assetGhRouter.POST("add-asset", privateapi.AddAsset)
		//删除asset
		assetGhRouter.POST("delete-asset-by-id", privateapi.DeleteAssetById)
		//根据ID查找asset
		assetGhRouter.POST("query-asset-by-id", privateapi.QueryAssetById)
		//获取月结列表
		assetGhRouter.POST("get-settlement-list", privateapi.GetSettlementList)
		//执行月结
		assetGhRouter.POST("set-settlement", privateapi.SetSettlement)
	}
}
