package purchasehub_supplier_order_goods_merge_repo

import "time"

// 分单中心-供应商分单-商品维度合单
//go:generate gormgen -structs PurchasehubSupplierOrderGoodsMerge -input .
type PurchasehubSupplierOrderGoodsMerge struct {
	Id        int32     // 自增ID
	Phsoid    int32     // 分单中心供应商订单ID
	Details   string    // 货号清单详情
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
