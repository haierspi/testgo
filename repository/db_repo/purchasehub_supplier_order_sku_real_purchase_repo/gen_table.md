#### offline_center_online.purchasehub_supplier_order_sku_real_purchase 
分单中心供应商切分订单真实采购需求

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 自增ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | phsoid | 分单中心供应商订单ID | int(10) unsigned | UNI | NO |  |  |
| 3 | sku_purchase_quantity | SKU采购信息 | text |  | YES |  |  |
| 4 | sku_purchase_price | SKU采购价格 | text |  | YES |  |  |
| 5 | sku_goods_details | SKU 货号详情 | text |  | YES |  |  |
| 6 | purchase_quantity | 采购总计数量 | int(10) unsigned |  | NO |  |  |
| 7 | purchase_amount | 采购总计金额 | decimal(10,4) unsigned |  | NO |  |  |
| 8 | status | 处理状态 0:待处理 1:已处理 | tinyint(1) unsigned |  | NO |  | 0 |
| 9 | is_change | 是否变动 0:已变更 1:无变更 | tinyint(1) unsigned |  | NO |  | 0 |
| 10 | created_at | 创建时间 | timestamp |  | YES |  |  |
| 11 | updated_at | 更新时间 | timestamp |  | YES |  |  |
