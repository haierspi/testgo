#### offline_center_online.purchasehub_supplier_order_flow 
采购单关键节点表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | erp_billnumber | 采购单号 | char(255) | MUL | NO |  |  |
| 3 | name | 节点名称 | varchar(50) |  | YES |  |  |
| 4 | operate | 节点操作者 | varchar(50) |  | YES |  |  |
| 5 | operate_id | 节点操作者id | varchar(50) |  | YES |  |  |
| 6 | note | 备注 | varchar(50) |  | YES |  |  |
| 7 | start_time | 操作开始时间 | timestamp |  | YES |  |  |
| 8 | end_time | 操作结束时间 | timestamp |  | YES |  |  |
| 9 | created_at |  | timestamp |  | YES |  |  |
| 10 | updated_at |  | timestamp |  | YES |  |  |
