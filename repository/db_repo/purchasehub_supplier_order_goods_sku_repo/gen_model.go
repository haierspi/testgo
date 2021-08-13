package purchasehub_supplier_order_goods_sku_repo

import "time"

// 采购中心供应商切分订单商品SKU
//go:generate gormgen -structs PurchasehubSupplierOrderGoodsSku -input .
type PurchasehubSupplierOrderGoodsSku struct {
	Id                    int32     // 采购中心ID
	Phsogid               int32     // 订单内单独的商品ID
	Phsoid                int32     // 采购中心供应商订单ID
	Phoid                 int32     // 采购中心ID
	GoodsId               int32     // 商品ID
	GoodsSn               string    // 商品货号
	SkuId                 string    // 商品货号
	SkuNumSn              string    // 商品货号
	UpdatedAt             time.Time `gorm:"time"` // 更新时间
	CreatedAt             time.Time `gorm:"time"` // 创建时间
	Price                 float64   // 商品采购价
	InPrice               float64   // 商品进货价
	PurchaseQuantity      int32     // 采购数量
	PurchaseAmount        float64   // SKU商品采购小计金额
	PurchaseOrderQuantity int32     // 采购接单总计数量
	PurchaseOrderPrice    float64   // SKU接单单价
	PurchaseOrderAmount   float64   // SKU接单金额
	DeliveryNum           int32     // 发货数量
	SuppSkuSn             string    // 供应商SKU(供应商货号+尺码名)
	StockNum              string    // 入库数量
	SizeEn                string    // 尺码EN名
	Status                int32     // 1,正常 2 挂起 3 删除
	HangupQuantity        int32     // 挂起数量
	IsHangup              int32     // 是否挂起 0：未挂起 1：挂起
	StockoutType          int32     // 商品缺货类型 0: 不缺货 1: 连带缺货 2: 采购缺货
}
