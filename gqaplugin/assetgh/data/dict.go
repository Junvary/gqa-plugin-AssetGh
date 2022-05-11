package data

import (
	"fmt"
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var PluginAssetGhSysDict = new(sysDict)

type sysDict struct{}

func (s *sysDict) LoadData() error {
	return gqaGlobal.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&gqaModel.SysDict{}).
			Where("parent_code = ?", "AssetGhCatalog").
			Or("dict_code = ?", "AssetGhCatalog").
			Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_dict 表中AssetGh插件菜单已存在，跳过初始化数据！数据量：", count)
			gqaGlobal.GqaLogger.Warn("[GQA-Plugin] --> sys_dict 表中AssetGh插件菜单已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysDictData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> AssetGh插件初始数据进入 sys_dict 表成功！")
		gqaGlobal.GqaLogger.Info("[GQA-Plugin] --> AssetGh插件初始数据进入 sys_dict 表成功！")
		return nil
	})
}

var sysDictData = []gqaModel.SysDict{
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1101, Memo: "工会固定资产分类", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, DictCode: "AssetGhCatalog", DictLabel: "工会固定资产分类"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1, Memo: "房屋及构筑物", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGhCatalog", DictCode: "AssetGh_FangWu", DictLabel: "房屋及构筑物"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1, Memo: "业务及管理用房-钢结构", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu__YeWu_GangJieGou", DictLabel: "业务及管理用房-钢结构", DictExt1: "50"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 2, Memo: "业务及管理用房-钢筋混凝土结构", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu_YeWu_GangJin", DictLabel: "业务及管理用房-钢筋混凝土结构", DictExt1: "50"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 3, Memo: "业务及管理用房-砖混结构", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu_YeWu_ZhuanHun", DictLabel: "业务及管理用房-砖混结构", DictExt1: "30"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 4, Memo: "业务及管理用房-砖木结构", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu_YeWu_ZhuanMu", DictLabel: "业务及管理用房-砖木结构", DictExt1: "30"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 5, Memo: "简易房", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu_JianYiFang", DictLabel: "简易房", DictExt1: "8"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 6, Memo: "房屋附属设施", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu_FuShu", DictLabel: "房屋附属设施", DictExt1: "8"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 7, Memo: "构筑物", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_FangWu", DictCode: "AssetGh_FangWu_GouZhu", DictLabel: "构筑物", DictExt1: "8"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 2, Memo: "通用设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGhCatalog", DictCode: "AssetGh_TongYong", DictLabel: "通用设备"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1, Memo: "计算机设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_JiSuanJi", DictLabel: "计算机设备", DictExt1: "6"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 2, Memo: "办公设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_BanGong", DictLabel: "办公设备", DictExt1: "6"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 3, Memo: "车辆", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_CheLiang", DictLabel: "车辆", DictExt1: "8"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 4, Memo: "图书档案设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_TuShu", DictLabel: "图书档案设备", DictExt1: "5"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 5, Memo: "机械设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_JiXie", DictLabel: "机械设备", DictExt1: "10"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 6, Memo: "电气设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_DianQi", DictLabel: "电气设备", DictExt1: "5"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 7, Memo: "雷达、无线电和卫星导航设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_LeiDa", DictLabel: "雷达、无线电和卫星导航设备", DictExt1: "10"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 8, Memo: "通信设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_TongXin", DictLabel: "通信设备", DictExt1: "5"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 9, Memo: "广播、电视、电影设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_GuangBo", DictLabel: "广播、电视、电影设备", DictExt1: "5"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 10, Memo: "仪表设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_YiBiao", DictLabel: "仪表设备", DictExt1: "5"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 11, Memo: "电子和通信测量设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_DianZi", DictLabel: "电子和通信测量设备", DictExt1: "5"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 12, Memo: "计量标准器具及量具、衡器", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_TongYong", DictCode: "AssetGh_TongYong_JiLiang", DictLabel: "计量标准器具及量具、衡器", DictExt1: "5"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 3, Memo: "专用设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGhCatalog", DictCode: "AssetGh_ZhuanYong", DictLabel: "专用设备"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1, Memo: "食品加工专用设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_ShiPin", DictLabel: "食品加工专用设备", DictExt1: "10", DictExt2: "15"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 2, Memo: "纺织设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_FangZhi", DictLabel: "纺织设备", DictExt1: "10", DictExt2: "15"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 3, Memo: "缝纫、服饰、制革和毛皮加工设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_FengRen", DictLabel: "缝纫、服饰、制革和毛皮加工设备", DictExt1: "10", DictExt2: "15"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 4, Memo: "医疗设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_YiLiao", DictLabel: "医疗设备", DictExt1: "5", DictExt2: "10"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 5, Memo: "安全生成设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_AnQuan", DictLabel: "安全生成设备", DictExt1: "10", DictExt2: "20"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 6, Memo: "环境污染防治设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_HuanJing", DictLabel: "环境污染防治设备", DictExt1: "10", DictExt2: "20"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 7, Memo: "文艺设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_WenYi", DictLabel: "文艺设备", DictExt1: "5", DictExt2: "15"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 8, Memo: "体育设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_TiYu", DictLabel: "体育设备", DictExt1: "5", DictExt2: "15"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 9, Memo: "娱乐设备", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_ZhuanYong", DictCode: "AssetGh_ZhuanYong_YuLe", DictLabel: "娱乐设备", DictExt1: "5", DictExt2: "15"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 4, Memo: "家具、用具及装具", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGhCatalog", DictCode: "AssetGh_JiaJu", DictLabel: "家具、用具及装具"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1, Memo: "家具", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_JiaJu", DictCode: "AssetGh_JiaJu_JiaJu", DictLabel: "家具", DictExt1: "15"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 2, Memo: "用具及装具", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGh_JiaJu", DictCode: "AssetGh_JiaJu_YongJu", DictLabel: "用具及装具", DictExt1: "5"},

	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1102, Memo: "工会固定资产使用状态", CreatedAt: time.Now(), CreatedBy: "admin"}},
		DictCode: "AssetGhUseStatus", DictLabel: "工会固定资产使用状态"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 1, Memo: "在用", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGhUseStatus", DictCode: "assetGh_on_use", DictLabel: "在用"},
	{GqaModelWithCreatedByAndUpdatedBy: gqaModel.GqaModelWithCreatedByAndUpdatedBy{GqaModel: gqaGlobal.GqaModel{
		Stable: "no", Sort: 2, Memo: "报废", CreatedAt: time.Now(), CreatedBy: "admin",
	}}, ParentCode: "AssetGhUseStatus", DictCode: "assetGh_scrap", DictLabel: "报废"},
}
