package purchasehub_supplier_order_callback_repo

import "time"

// 供应商切分订单回调相关数据
//go:generate gormgen -structs PurchasehubSupplierOrderCallback -input .
type PurchasehubSupplierOrderCallback struct {
	Id        int32     // 自增ID
	Phsoid    int32     // 分单中心供应商订单ID
	Type      string    // 回调类型
	Status    int32     // 处理状态 0:待处理 1:已处理
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
