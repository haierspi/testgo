package purchasehub_supplier_order_real_details_repo

import "time"

// 分单中心供应商切分订单真实采购需求
//go:generate gormgen -structs PurchasehubSupplierOrderRealDetails -input .
type PurchasehubSupplierOrderRealDetails struct {
	Id               int32     // 自增ID
	Phsoid           int32     // 分单中心供应商订单ID
	Details          string    // SKU清单详情
	SkuCount         int32     // SKU种类数量
	GoodsCount       int32     // 商品数量
	PurchaseQuantity int32     // 采购总计数量
	PurchaseAmount   float64   // 采购总计金额
	Status           int32     // 处理状态 0:待处理 1:已处理
	IsChange         int32     // 是否变动 0:已变更 1:无变更
	CreatedAt        time.Time `gorm:"time"` // 创建时间
	UpdatedAt        time.Time `gorm:"time"` // 更新时间
}
