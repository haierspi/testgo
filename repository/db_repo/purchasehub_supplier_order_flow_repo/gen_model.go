package purchasehub_supplier_order_flow_repo

import "time"

// 采购单关键节点表
//go:generate gormgen -structs PurchasehubSupplierOrderFlow -input .
type PurchasehubSupplierOrderFlow struct {
	Id            int32     //
	ErpBillnumber string    // 采购单号
	Name          string    // 节点名称
	Operate       string    // 节点操作者
	OperateId     string    // 节点操作者id
	Note          string    // 备注
	StartTime     time.Time `gorm:"time"` // 操作开始时间
	EndTime       time.Time `gorm:"time"` // 操作结束时间
	CreatedAt     time.Time `gorm:"time"` //
	UpdatedAt     time.Time `gorm:"time"` //
}
