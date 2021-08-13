package purchasehub_supplier_order_log_repo

import "time"

// 采购单操作日志
//go:generate gormgen -structs PurchasehubSupplierOrderLog -input .
type PurchasehubSupplierOrderLog struct {
	Id            int32     //
	Phsoid        int32     //
	ErpBillnumber string    // 普元单号
	Type          string    // 操作类型
	Operate       string    // 操作类型
	OperateId     string    //
	OperateRes    int32     // 操作结果  1 成功 2 驳回
	Note          string    // 备注
	RawData       string    //
	TimeOut       string    //
	CreatedAt     time.Time `gorm:"time"` //
	UpdatedAt     time.Time `gorm:"time"` //
}
