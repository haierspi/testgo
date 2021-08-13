package purchasehub_supplier_order_repo

import "time"

// 采购中心供应商切分订单
//go:generate gormgen -structs PurchasehubSupplierOrder -input .
type PurchasehubSupplierOrder struct {
	Phsoid                     int32     // 采购中心供应商订单ID
	Phoid                      int32     // 上下文目标ID
	SupplierId                 int32     // 供应商ID
	TargetContextId            int32     // 目标ID
	TargetContextType          int32     // 目标类型 1:1688自动下单 2.供应商平台 3.SCM下单 4.线下订单
	TargetContextCustom        string    // 目标自定义凭证
	ErpBillnumber              string    // 普源采购单
	AlibabaOrderId             string    // 1688订单号
	PurchaseQuantity           int32     // 采购总计数量
	PurchaseAmount             float64   // 接单总计金额
	PurchaseOrderQuantity      int32     // 接单总计数量
	PurchaseOrderAmount        float64   // 采购总计金额
	IsDeleted                  int32     // 是否删除 0:正常订单 1:已标记删除
	Status                     int32     // 采购商品状态 0:待处理 1:已处理
	CreatedUid                 int32     // 创建UID
	CreatedNickname            string    // 创建昵称
	UpdatedAt                  time.Time `gorm:"time"` // 更新时间
	CreatedAt                  time.Time `gorm:"time"` // 创建时间
	DeletedAt                  time.Time `gorm:"time"` // 标记删除时间
	Type                       int32     // 程序分配的目标平台 1:1688采购;2:供应商平台;3:SCM;4:线下单
	ProductLine                string    // 分配产品线
	NeedCallback               int32     // 需要回调
	CateId                     int32     // 订单类别ID
	CateText                   string    // 订单类别
	MergeTargetPhsoid          int32     // 合单目标订单 - 采购中心供应商订单ID
	MergeSourcePhsoid          int32     // 合并订单原单 - 采购中心供应商订单ID
	IsMerged                   int32     // 是否被合单 0;1
	IsMerge                    int32     // 是否为合单 0;1
	IsPull                     int32     // 1688拉取状态 0不需要1需要2已拉取
	StoreName                  string    // 仓库名
	DeletedNickname            string    // 删除人
	DeletedUid                 int32     // 删除人
	CallbackType               int32     // 回调类型,参考 ly_purchasehub_order.source_context_type
	DemandManageUid            int32     // 需求管理操作人UID [自动下需求 直接为0]
	DemandManageNickname       string    // 需求管理操作人昵称
	SupplierAbbreviation       string    // 供应商简称
	SupplierFullName           string    // 供应商全称
	SuppliersType              int32     // 供应商分类 1 生产商 2 贸易商 3 代理商 4工贸一体
	OrderTypeId                int32     // 接单类型：1线下接单 2平台接单 31688自动拍单
	SettlementMethodId         int32     // 结算方式：1先款后货 2货到付款 3预付 4 1688 虚拟拍单
	SkuCount                   int32     // SKU种类数量
	GoodsCount                 int32     // 商品数量
	StockInQuantity            int32     // 入库总计数量
	StockInAmount              float64   // 入库总计金额
	DiscountMoney              float64   // 减免金额
	ExpressFee                 float64   // 快递费用
	TimeFee                    float64   // 全检费
	RefundFee                  float64   // 退货费
	SpecialLogisticsCostsTotal float64   // 专线物流费用
	PurchaserId                int32     // 采购员id
	PurchaserUid               int32     // 采购员UID ly_admin_user表id
	PurchaserName              string    // 采购员名字
	FlowStatus                 int32     // 流程状态(二进制状态位) 1:
	CurFlowStatus              int32     // 当前状态(单状态)
	IsArchive                  int32     // 是否已归档 0:未归档 1:已归档 (已归档表示订单中心同步)
	IsRepeal                   int32     // 是否废除 0：未废除 1：已废除 对应普元 CheckFlag = 3
	Note                       string    // 订单备注
	ExaminePriceResult         string    // 进入核价原因
	ExaminePriceKey            string    // 进入核价原因关键Key,对应 PurchaseHubWaitingExaminePriceLogic 的常量
	ExaminePriceData           string    // 进入核价原因涉及的原因值
	PushErp                    int32     // 推送普源 0:不需要推送; 1:等待推送 2:已推送
	IsReal                     int32     // 是否拉取过真实接单数 0:没有; 1:已经拉取过
	SourceErpBillnumber        string    // 来源采购单订单号
	SourcePhsoid               int32     // 来源采购单ID
	IsSingle                   int32     // 是否为单货号订单 0：非单货号订单 1：单货号订单
	DeliveryDate               time.Time `gorm:"time"` // 到货日期
}
