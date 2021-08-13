#### offline_center_online.purchasehub_supplier_order_log 
采购单操作日志

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | phsoid |  | int(11) |  | YES |  |  |
| 3 | erp_billnumber | 普元单号 | char(50) |  | YES |  |  |
| 4 | type | 操作类型 | char(50) |  | YES |  |  |
| 5 | operate | 操作类型 | char(50) |  | YES |  |  |
| 6 | operate_id |  | char(50) |  | YES |  |  |
| 7 | operate_res | 操作结果  1 成功 2 驳回 | int(11) |  | YES |  |  |
| 8 | note | 备注 | varchar(255) |  | YES |  |  |
| 9 | raw_data |  | text |  | YES |  |  |
| 10 | time_out |  | varchar(255) |  | YES |  |  |
| 11 | created_at |  | timestamp |  | YES |  |  |
| 12 | updated_at |  | timestamp |  | YES |  |  |
