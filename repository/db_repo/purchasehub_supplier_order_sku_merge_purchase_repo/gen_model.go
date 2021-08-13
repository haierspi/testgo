package purchasehub_supplier_order_sku_merge_purchase_repo

import "time"

// 分单中心-供应商分单-合单详情
//go:generate gormgen -structs PurchasehubSupplierOrderSkuMergePurchase -input .
type PurchasehubSupplierOrderSkuMergePurchase struct {
	Id                  int32     // 自增ID
	Phsoid              int32     // 分单中心供应商订单ID
	SkuPurchaseQuantity string    // SKU采购信息
	SkuPurchasePrice    string    // SKU采购价格
	SkuGoodsDetails     string    // SKU 货号详情
	PurchaseQuantity    int32     // 采购总计数量
	PurchaseAmount      float64   // 采购总计金额
	CreatedAt           time.Time `gorm:"time"` // 创建时间
	UpdatedAt           time.Time `gorm:"time"` // 更新时间
}
