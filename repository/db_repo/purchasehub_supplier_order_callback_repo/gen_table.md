#### offline_center_online.purchasehub_supplier_order_callback 
供应商切分订单回调相关数据

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 自增ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | phsoid | 分单中心供应商订单ID | int(10) unsigned | MUL | NO |  |  |
| 3 | type | 回调类型 | varchar(255) |  | NO |  |  |
| 4 | status | 处理状态 0:待处理 1:已处理 | tinyint(1) unsigned |  | NO |  | 0 |
| 5 | created_at | 创建时间 | timestamp |  | YES |  |  |
| 6 | updated_at | 更新时间 | timestamp |  | YES |  |  |
