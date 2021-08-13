#### offline_center_online.purchasehub_supplier_order 
采购中心供应商切分订单

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | phsoid | 采购中心供应商订单ID | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | phoid | 上下文目标ID | int(11) unsigned | MUL | YES |  |  |
| 3 | supplier_id | 供应商ID | int(11) unsigned |  | YES |  |  |
| 4 | target_context_id | 目标ID | int(11) unsigned |  | YES |  |  |
| 5 | target_context_type | 目标类型 1:1688自动下单 2.供应商平台 3.SCM下单 4.线下订单 | int(11) unsigned | MUL | YES |  |  |
| 6 | target_context_custom | 目标自定义凭证 | char(255) | PRI | NO |  |  |
| 7 | erp_billnumber | 普源采购单 | char(255) | MUL | NO |  |  |
| 8 | alibaba_order_id | 1688订单号 | varchar(50) |  | YES |  |  |
| 9 | purchase_quantity | 采购总计数量 | int(10) unsigned |  | NO |  |  |
| 10 | purchase_amount | 接单总计金额 | decimal(10,4) unsigned |  | NO |  |  |
| 11 | purchase_order_quantity | 接单总计数量 | int(10) unsigned |  | NO |  |  |
| 12 | purchase_order_amount | 采购总计金额 | decimal(10,4) unsigned |  | NO |  |  |
| 13 | is_deleted | 是否删除 0:正常订单 1:已标记删除 | tinyint(1) unsigned | MUL | NO |  | 0 |
| 14 | status | 采购商品状态 0:待处理 1:已处理  | tinyint(1) unsigned |  | NO |  | 0 |
| 15 | created_uid | 创建UID | int(10) unsigned |  | NO |  |  |
| 16 | created_nickname | 创建昵称 | char(255) |  | YES |  |  |
| 17 | updated_at | 更新时间 | timestamp |  | YES |  |  |
| 18 | created_at | 创建时间 | timestamp | MUL | YES |  |  |
| 19 | deleted_at | 标记删除时间 | timestamp |  | YES |  |  |
| 20 | type | 程序分配的目标平台 1:1688采购;2:供应商平台;3:SCM;4:线下单  | tinyint(1) unsigned |  | NO |  | 0 |
| 21 | product_line | 分配产品线  | char(255) |  | NO |  |  |
| 22 | need_callback | 需要回调 | tinyint(1) unsigned |  | NO |  | 0 |
| 23 | cate_id | 订单类别ID | int(11) |  | NO |  |  |
| 24 | cate_text | 订单类别  | char(255) |  | YES |  |  |
| 25 | merge_target_phsoid | 合单目标订单 - 采购中心供应商订单ID | int(11) unsigned |  | YES |  |  |
| 26 | merge_source_phsoid | 合并订单原单 - 采购中心供应商订单ID | int(11) unsigned |  | YES |  |  |
| 27 | is_merged | 是否被合单 0;1 | tinyint(1) unsigned |  | NO |  | 0 |
| 28 | is_merge | 是否为合单 0;1 | tinyint(1) unsigned |  | NO |  | 0 |
| 29 | is_pull | 1688拉取状态 0不需要1需要2已拉取 | tinyint(1) |  | YES |  | 0 |
| 30 | store_name | 仓库名  | char(255) |  | NO |  |  |
| 31 | deleted_nickname | 删除人 | varchar(255) |  | NO |  |  |
| 32 | deleted_uid | 删除人 | mediumint(8) unsigned |  | NO |  | 0 |
| 33 | callback_type | 回调类型,参考 ly_purchasehub_order.source_context_type | int(11) unsigned |  | YES |  |  |
| 34 | demand_manage_uid | 需求管理操作人UID [自动下需求 直接为0] | int(10) unsigned |  | YES |  |  |
| 35 | demand_manage_nickname | 需求管理操作人昵称 | char(255) |  | YES |  |  |
| 36 | supplier_abbreviation | 供应商简称 | varchar(100) |  | YES |  |  |
| 37 | supplier_full_name | 供应商全称 | varchar(100) |  | YES |  |  |
| 38 | suppliers_type | 供应商分类 1 生产商 2 贸易商 3 代理商 4工贸一体 | int(11) |  | NO |  | 0 |
| 39 | order_type_id | 接单类型：1线下接单 2平台接单 31688自动拍单 | tinyint(3) unsigned |  | NO |  | 0 |
| 40 | settlement_method_id | 结算方式：1先款后货 2货到付款 3预付 4 1688 虚拟拍单 | tinyint(3) unsigned |  | NO |  | 0 |
| 41 | sku_count | SKU种类数量 | int(11) unsigned |  | NO |  | 0 |
| 42 | goods_count | 商品数量 | int(11) unsigned |  | NO |  | 0 |
| 43 | stock_in_quantity | 入库总计数量 | int(10) unsigned |  | NO |  | 0 |
| 44 | stock_in_amount | 入库总计金额 | decimal(10,4) unsigned |  | NO |  | 0.0000 |
| 45 | discount_money | 减免金额 | decimal(10,4) unsigned |  | NO |  | 0.0000 |
| 46 | express_fee | 快递费用 | decimal(10,4) unsigned |  | NO |  | 0.0000 |
| 47 | time_fee | 全检费 | decimal(10,4) unsigned |  | NO |  | 0.0000 |
| 48 | refund_fee | 退货费 | decimal(10,4) unsigned |  | NO |  | 0.0000 |
| 49 | special_logistics_costs_total | 专线物流费用 | decimal(10,4) unsigned |  | NO |  | 0.0000 |
| 50 | purchaser_id | 采购员id | int(11) |  | NO |  | 0 |
| 51 | purchaser_uid | 采购员UID ly_admin_user表id | int(11) |  | NO |  | 0 |
| 52 | purchaser_name | 采购员名字 | varchar(100) |  | NO |  |  |
| 53 | flow_status | 流程状态(二进制状态位) 1:     | int(64) unsigned |  | NO |  | 0 |
| 54 | cur_flow_status | 当前状态(单状态) | int(64) unsigned |  | NO |  | 0 |
| 55 | is_archive | 是否已归档 0:未归档 1:已归档 (已归档表示订单中心同步) | tinyint(1) unsigned |  | NO |  | 0 |
| 56 | is_repeal | 是否废除 0：未废除 1：已废除 对应普元 CheckFlag = 3 | tinyint(1) unsigned |  | YES |  | 0 |
| 57 | note | 订单备注 | varchar(255) |  | YES |  |  |
| 58 | examine_price_result | 进入核价原因 | varchar(255) |  | NO |  |  |
| 59 | examine_price_key | 进入核价原因关键Key,对应 PurchaseHubWaitingExaminePriceLogic 的常量 | varchar(100) |  | NO |  |  |
| 60 | examine_price_data | 进入核价原因涉及的原因值 | text |  | YES |  |  |
| 61 | push_erp | 推送普源 0:不需要推送; 1:等待推送 2:已推送 | tinyint(1) unsigned | MUL | NO |  | 0 |
| 62 | is_real | 是否拉取过真实接单数 0:没有; 1:已经拉取过 | tinyint(1) unsigned |  | NO |  | 0 |
| 63 | source_erp_billnumber | 来源采购单订单号 | varchar(255) |  | NO |  |  |
| 64 | source_phsoid | 来源采购单ID | int(10) unsigned |  | NO |  | 0 |
| 65 | is_single | 是否为单货号订单 0：非单货号订单 1：单货号订单 | tinyint(1) |  | YES |  | 0 |
| 66 | delivery_date | 到货日期 | timestamp |  | YES |  |  |
