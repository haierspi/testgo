package purchasehub_supplier_order_goods_sku_merge_repo

import "time"

// 分单中心-供应商分单-SKU维度合单详情
//go:generate gormgen -structs PurchasehubSupplierOrderGoodsSkuMerge -input .
type PurchasehubSupplierOrderGoodsSkuMerge struct {
	Id        int32     // 自增ID
	Phsoid    int32     // 分单中心供应商订单ID
	Details   string    // SKU清单详情
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
