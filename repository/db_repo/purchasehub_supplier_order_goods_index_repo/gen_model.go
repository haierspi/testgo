package purchasehub_supplier_order_goods_index_repo

import "time"

// 分单中心供应商切分订单商品清单
//go:generate gormgen -structs PurchasehubSupplierOrderGoodsIndex -input .
type PurchasehubSupplierOrderGoodsIndex struct {
	Id        int32     // 分单中心ID
	GoodsId   int32     // 商品ID
	Phsoid    int32     // 分单中心供应商订单ID
	UpdatedAt time.Time `gorm:"time"` // 更新时间
	CreatedAt time.Time `gorm:"time"` // 创建时间
}
