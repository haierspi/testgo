package purchasehub_supplier_order_stockout_type_index_repo

import "time"

// 分单中心供应商切分订单缺货类型关联表
//go:generate gormgen -structs PurchasehubSupplierOrderStockoutTypeIndex -input .
type PurchasehubSupplierOrderStockoutTypeIndex struct {
	Id           int32     // 分单中心ID
	StockoutType int32     // 商品缺货类型 0: 不缺货 1: 连带缺货 2: 采购缺货
	Phsoid       int32     // 分单中心供应商订单ID
	UpdatedAt    time.Time `gorm:"time"` // 更新时间
	CreatedAt    time.Time `gorm:"time"` // 创建时间
}
