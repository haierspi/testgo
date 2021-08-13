#### offline_center_online.purchasehub_supplier_order_goods_index 
分单中心供应商切分订单商品清单

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 分单中心ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | goods_id | 商品ID | int(10) unsigned | MUL | NO |  |  |
| 3 | phsoid | 分单中心供应商订单ID | int(10) unsigned | MUL | NO |  |  |
| 4 | updated_at | 更新时间 | timestamp |  | YES |  |  |
| 5 | created_at | 创建时间 | timestamp |  | YES |  |  |
