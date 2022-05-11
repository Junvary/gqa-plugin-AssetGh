package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginAssetGhSysRoleApi = new(sysRoleApi)

type sysRoleApi struct{}

func (s *sysRoleApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysRoleApi{}).Where("api_group = ?", "plugin-AssetGh").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_api 表中AssetGh插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_role_api 表中AssetGh插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> AssetGh插件初始数据进入 sys_role_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> AssetGh插件初始数据进入 sys_role_api 表成功！")
		return nil
	})
}

var sysRoleApiData = []gqaModel.SysRoleApi{
	//asset
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/get-asset-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/edit-asset"},
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/add-asset"},
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/delete-asset-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/query-asset-by-id"},
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/get-settlement-list"},
	{RoleCode: "super-admin", ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/set-settlement"},
}
