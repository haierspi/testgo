package purchasehub_supplier_order_goods_repo

import "time"

// 分单中心供应商切分订单商品清单
//go:generate gormgen -structs PurchasehubSupplierOrderGoods -input .
type PurchasehubSupplierOrderGoods struct {
	Phsogid               int32     // 订单内单独的商品ID
	Phsoid                int32     // 分单中心供应商订单ID
	Phoid                 int32     // 分单中心ID
	ErpBillnumber         string    // 普源采购单
	GoodsId               int32     // 商品ID
	GoodsSn               string    // 商品货号
	GoodsName             string    // 商品名称
	GoodsImg              string    // 商品图片
	OnSaleTime            time.Time `gorm:"time"` // 上架时间
	Color                 string    // 商品颜色
	Price                 float64   // 商品采购价
	InPrice               float64   // 商品进货价
	ProductLine           string    // 产品线(lw,pop)
	PurchaseQuantity      int32     // 采购数量
	PurchaseAmount        float64   // 商品采购小计金额
	PurchaseOrderQuantity int32     // 商品接单数量
	PurchaseOrderPrice    float64   // 货号接单单价
	PurchaseOrderAmount   float64   // 商品接单总计金额
	SkuCount              int32     // 货号下SKU种类数量
	DeveloperUid          int32     // 开发员UID
	DeveloperName         string    // 开发员名字
	PurchaserId           int32     // 采购员id
	PurchaserUid          int32     // 采购员UID ly_admin_user表id
	PurchaserName         string    // 采购员名字
	SupplyType            int32     // 供应类型  0 现货 1生产
	SupplyStatus          int32     // 供应状态 0 停供 1 正常 2 缺货
	SuppSn                string    // 款号(供应商货号)
	SuppUrl               string    // 供应商商品链接
	SuppliersName         string    // 供应商名称
	SuppliersGoodsImg     string    // 供应商商品图片
	UpdatedAt             time.Time `gorm:"time"` // 更新时间
	CreatedAt             time.Time `gorm:"time"` // 创建时间
	ProductType           int32     // 商品新旧类型 0:未定义; 1:新新品 2:新品 3:老品
	FirstOrder            int32     // 是否为首单 0:去SCM 1:首单; 2:返单
	IsExaminePrice        int32     // 商品是否需要核价 0:不需要; 1:需要
	ExaminePriceResult    string    // 进入核价原因
	ExaminePriceKey       string    // 进入核价原因关键Key,对应 PurchaseHubWaitingExaminePriceLogic 的常量
	ExaminePriceData      string    // 进入核价原因涉及的原因值
	StockoutType          int32     // 商品缺货类型 0: 不缺货 1: 连带缺货 2: 采购缺货
}
