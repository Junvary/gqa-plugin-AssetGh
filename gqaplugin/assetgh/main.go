package assetgh

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/data"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/router/privaterouter"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/assetgh/router/publicrouter"
	"github.com/gin-gonic/gin"
)

var PluginAssetGh = new(assetGh)

type assetGh struct{}

func (f *assetGh) PluginCode() string {
	return "plugin-AssetGh"
}

func (f *assetGh) PluginName() string {
	return "工会固定资产管理"
}

func (f *assetGh) PluginVersion() string {
	return "v0.0.1"
}

func (f *assetGh) PluginMemo() string {
	return "这是工会固定资产管理插件"
}

func (f *assetGh) PluginRouterPublic(publicGroup *gin.RouterGroup) {
	publicrouter.InitPublicRouter(publicGroup)
}

func (f *assetGh) PluginRouterPrivate(privateGroup *gin.RouterGroup) {
	privaterouter.InitPrivateRouter(privateGroup)
}

func (f *assetGh) PluginMigrate() []interface{} {
	var ModelList = []interface{}{
		model.GqaPluginAssetGh{},
		model.GqaPluginAssetGhSettlement{},
	}
	return ModelList
}

func (f *assetGh) PluginData() []interface{ LoadData() (err error) } {
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginAssetGhSysApi,
		data.PluginAssetGhSysRoleApi,
		data.PluginAssetGhSysMenu,
		data.PluginAssetGhSysRoleMenu,
		data.PluginAssetGhSysDict,
	}
	return DataList
}
