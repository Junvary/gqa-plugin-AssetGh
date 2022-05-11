package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAssetGhSysApi = new(sysApi)

type sysApi struct{}

func (s *sysApi) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysApi{}).Where("api_group = ?", "plugin-AssetGh").Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_api 表中AssetGh插件数据已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_api 表中AssetGh插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysApiData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> AssetGh插件初始数据进入 sys_api 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> AssetGh插件初始数据进入 sys_api 表成功！")
		return nil
	})
}

var sysApiData = []gqaModel.SysApi{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1101, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：获取asset-list",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/get-asset-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1102, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：编辑asset信息",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/edit-asset"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1103, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：新增asset",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/add-asset"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1104, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：删除asset",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/delete-asset-by-id"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1105, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：根据ID查找asset",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/query-asset-by-id"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1106, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：获取月结列表",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/get-settlement-list"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Sort: 1107, Stable: "yes", CreatedBy: "admin", CreatedAt: time.Now(), Memo: "插件AssetGh：执行月结",
	}}, ApiGroup: "plugin-AssetGh", ApiMethod: "POST", ApiPath: "/plugin-AssetGh/set-settlement"},
}
