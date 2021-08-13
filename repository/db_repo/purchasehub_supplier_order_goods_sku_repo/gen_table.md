#### offline_center_online.purchasehub_supplier_order_goods_sku 
采购中心供应商切分订单商品SKU

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 采购中心ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | phsogid | 订单内单独的商品ID | int(10) unsigned | MUL | NO |  |  |
| 3 | phsoid | 采购中心供应商订单ID | int(10) unsigned | MUL | NO |  |  |
| 4 | phoid | 采购中心ID | int(11) unsigned | MUL | YES |  |  |
| 5 | goods_id | 商品ID | int(10) unsigned |  | NO |  |  |
| 6 | goods_sn | 商品货号 | varchar(64) |  | NO |  |  |
| 7 | sku_id | 商品货号 | varchar(64) |  | NO |  |  |
| 8 | sku_num_sn | 商品货号 | varchar(64) |  | NO |  |  |
| 9 | updated_at | 更新时间 | timestamp |  | YES |  |  |
| 10 | created_at | 创建时间 | timestamp | MUL | YES |  |  |
| 11 | price | 商品采购价 | decimal(10,4) unsigned |  | NO |  |  |
| 12 | in_price | 商品进货价 | decimal(10,4) unsigned |  | NO |  |  |
| 13 | purchase_quantity | 采购数量 | int(10) unsigned |  | NO |  |  |
| 14 | purchase_amount | SKU商品采购小计金额 | decimal(10,4) unsigned |  | NO |  |  |
| 15 | purchase_order_quantity | 采购接单总计数量 | int(10) unsigned |  | NO |  |  |
| 16 | purchase_order_price | SKU接单单价 | decimal(10,4) unsigned |  | NO |  |  |
| 17 | purchase_order_amount | SKU接单金额 | decimal(10,4) unsigned |  | NO |  |  |
| 18 | delivery_num | 发货数量 | int(10) unsigned |  | NO |  |  |
| 19 | supp_sku_sn | 供应商SKU(供应商货号+尺码名) | varchar(255) |  | NO |  |  |
| 20 | stock_num | 入库数量 | varchar(255) |  | NO |  | 0 |
| 21 | size_en | 尺码EN名 | varchar(60) |  | NO |  |  |
| 22 | status | 1,正常 2 挂起 3 删除 | int(11) |  | YES |  |  |
| 23 | hangup_quantity | 挂起数量 | tinyint(1) |  | YES |  | 0 |
| 24 | is_hangup | 是否挂起 0：未挂起 1：挂起 | tinyint(1) |  | YES |  | 0 |
| 25 | stockout_type | 商品缺货类型 0: 不缺货 1: 连带缺货 2: 采购缺货 | tinyint(1) |  | YES |  | 0 |
