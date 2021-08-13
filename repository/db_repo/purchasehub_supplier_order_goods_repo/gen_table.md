#### offline_center_online.purchasehub_supplier_order_goods 
分单中心供应商切分订单商品清单

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | phsogid | 订单内单独的商品ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | phsoid | 分单中心供应商订单ID | int(10) unsigned | MUL | NO |  |  |
| 3 | phoid | 分单中心ID | int(11) unsigned | MUL | YES |  |  |
| 4 | erp_billnumber | 普源采购单 | char(255) | MUL | YES |  |  |
| 5 | goods_id | 商品ID | int(10) unsigned | MUL | NO |  |  |
| 6 | goods_sn | 商品货号 | varchar(64) |  | NO |  |  |
| 7 | goods_name | 商品名称 | varchar(64) |  | NO |  |  |
| 8 | goods_img | 商品图片 | varchar(255) |  | NO |  |  |
| 9 | on_sale_time | 上架时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 10 | color | 商品颜色 | varchar(50) |  | NO |  |  |
| 11 | price | 商品采购价 | decimal(10,4) unsigned |  | NO |  |  |
| 12 | in_price | 商品进货价 | decimal(10,4) unsigned |  | NO |  |  |
| 13 | product_line | 产品线(lw,pop) | varchar(30) |  | YES |  |  |
| 14 | purchase_quantity | 采购数量 | int(10) unsigned |  | NO |  |  |
| 15 | purchase_amount | 商品采购小计金额 | decimal(10,4) unsigned |  | NO |  |  |
| 16 | purchase_order_quantity | 商品接单数量 | int(10) unsigned |  | NO |  |  |
| 17 | purchase_order_price | 货号接单单价 | decimal(10,4) unsigned |  | NO |  |  |
| 18 | purchase_order_amount | 商品接单总计金额 | decimal(10,4) unsigned |  | NO |  |  |
| 19 | sku_count | 货号下SKU种类数量 | int(11) unsigned |  | NO |  | 0 |
| 20 | developer_uid | 开发员UID | int(11) |  | NO |  | 0 |
| 21 | developer_name | 开发员名字 | varchar(100) |  | NO |  |  |
| 22 | purchaser_id | 采购员id | int(11) |  | NO |  | 0 |
| 23 | purchaser_uid | 采购员UID ly_admin_user表id | int(11) |  | NO |  | 0 |
| 24 | purchaser_name | 采购员名字 | varchar(100) |  | NO |  |  |
| 25 | supply_type | 供应类型  0 现货 1生产 | int(11) |  | NO |  | 0 |
| 26 | supply_status | 供应状态 0 停供 1 正常 2 缺货 | int(11) |  | YES |  | 1 |
| 27 | supp_sn | 款号(供应商货号) | varchar(255) |  | NO |  |  |
| 28 | supp_url | 供应商商品链接 | varchar(255) |  | NO |  |  |
| 29 | suppliers_name | 供应商名称 | varchar(255) |  | NO |  |  |
| 30 | suppliers_goods_img | 供应商商品图片 | varchar(255) |  | YES |  |  |
| 31 | updated_at | 更新时间 | timestamp |  | YES |  |  |
| 32 | created_at | 创建时间 | timestamp |  | YES |  |  |
| 33 | product_type | 商品新旧类型 0:未定义; 1:新新品 2:新品 3:老品 | tinyint(1) unsigned |  | NO |  | 0 |
| 34 | first_order | 是否为首单 0:去SCM 1:首单; 2:返单 | tinyint(1) unsigned |  | NO |  | 0 |
| 35 | is_examine_price | 商品是否需要核价 0:不需要; 1:需要 | tinyint(1) unsigned |  | NO |  | 0 |
| 36 | examine_price_result | 进入核价原因 | varchar(255) |  | NO |  |  |
| 37 | examine_price_key | 进入核价原因关键Key,对应 PurchaseHubWaitingExaminePriceLogic 的常量 | varchar(100) |  | NO |  |  |
| 38 | examine_price_data | 进入核价原因涉及的原因值 | text |  | YES |  |  |
| 39 | stockout_type | 商品缺货类型 0: 不缺货 1: 连带缺货 2: 采购缺货 | tinyint(1) |  | YES |  | 0 |
