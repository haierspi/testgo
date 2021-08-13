#### offline_center_online.purchasehub_supplier_order_stockout_type_index 
分单中心供应商切分订单缺货类型关联表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 分单中心ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | stockout_type | 商品缺货类型 0: 不缺货 1: 连带缺货 2: 采购缺货 | tinyint(1) |  | YES |  | 0 |
| 3 | phsoid | 分单中心供应商订单ID | int(10) unsigned |  | NO |  |  |
| 4 | updated_at | 更新时间 | timestamp |  | YES |  |  |
| 5 | created_at | 创建时间 | timestamp | MUL | YES |  |  |
