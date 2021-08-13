package DemandPlanSearch

import (
	"goService/components"
	"time"

	"gorm.io/gorm"
)

// LyDemandPlanSearch 需求计划V2搜索列表
type DemandPlanSearch struct {
	gorm.Model
	ProductLine       string    `gorm:"column:product_line;type:varchar(10)"`                     // 产品线
	GoodsSn           string    `gorm:"index;column:goods_sn;type:varchar(20)"`                   // 商品货号
	MadeType          int       `gorm:"column:made_type;type:int(11)"`                            // 生产方式
	SalesRankID       int       `gorm:"column:sales_rank_id;type:int(11)"`                        // 销量等级
	SalesTrendID      int       `gorm:"column:sales_trend_id;type:int(11)"`                       // 销量趋势
	UpperDays         float64   `gorm:"column:upper_days;type:decimal(10,2) unsigned"`            // 上架天数
	StyleCategory     int8      `gorm:"column:style_category;type:tinyint(1)"`                    // 款式品类
	Season            int8      `gorm:"column:season;type:tinyint(1)"`                            // 季节
	TagMarketingID    int       `gorm:"column:tag_marketing_id;type:int(11) unsigned"`            // 营销标签
	TopCateID         int       `gorm:"column:top_cate_id;type:int(11) unsigned;not null"`        // 商品分类
	Sku               string    `gorm:"column:sku;type:varchar(255)"`                             // 商品SKU
	SupplierID        int       `gorm:"column:supplier_id;type:int(11)"`                          // 供应商
	SupplyType        int       `gorm:"column:supply_type;type:int(11)"`                          // 供应类型
	Day1SaleNum       float64   `gorm:"column:day1_sale_num;type:decimal(10,2) unsigned"`         // 昨日销售
	Day7AvgRevSaleNum float64   `gorm:"column:day7_avg_rev_sale_num;type:decimal(10,2) unsigned"` // 7日修复均销
	IsSkuSaleAbnormal int8      `gorm:"column:is_sku_sale_abnormal;type:tinyint(1);not null"`     // 是否删除 0：否 1：是
	IsGoodsSaleUprush int8      `gorm:"column:is_goods_sale_uprush;type:tinyint(1);not null"`     // 款销突增款 0：否 1：是
	IsBrandOperation  int8      `gorm:"column:is_brand_operation;type:tinyint(1);not null"`       // 品运判断款 0：否 1：是
	IsOther           int8      `gorm:"column:is_other;type:tinyint(1);not null"`                 // 正常款 0：否 1：是
	IsNewGoods        int8      `gorm:"column:is_new_goods;type:tinyint(1);not null"`             // 新品突出款 0：否 1：是
	IsChecked         int8      `gorm:"index;column:is_checked;type:tinyint(1)"`                  // 品运判断审核 0：不需要审核 1：需要审核 2 审核成功
	IsDeleted         int8      `gorm:"column:is_deleted;type:tinyint(1)"`                        // 是否删除 0：未删除 1：已删除
	DataCreatTime     time.Time `gorm:"column:data_creat_time;type:timestamp"`                    // 数据创建时间
	StockQty          int       `gorm:"column:stock_qty;type:int(11);not null"`                   // 商品备货数量
}

func NewModel() *DemandPlanSearch {
	return new(DemandPlanSearch)
}

func Get(id int) *DemandPlanSearch {
	model := new(DemandPlanSearch)
	components.Db.Where(id).First(model)
	return model
}
