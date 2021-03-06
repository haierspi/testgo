///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package purchasehub_supplier_order_log_repo

import (
	"fmt"
	"time"

	"go_service/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PurchasehubSupplierOrderLog {
	return new(PurchasehubSupplierOrderLog)
}

func NewQueryBuilder() *purchasehubSupplierOrderLogRepoQueryBuilder {
	return new(purchasehubSupplierOrderLogRepoQueryBuilder)
}

func (t *PurchasehubSupplierOrderLog) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type purchasehubSupplierOrderLogRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PurchasehubSupplierOrderLog{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PurchasehubSupplierOrderLog{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PurchasehubSupplierOrderLog{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) First(db *gorm.DB) (*PurchasehubSupplierOrderLog, error) {
	ret := &PurchasehubSupplierOrderLog{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) QueryOne(db *gorm.DB) (*PurchasehubSupplierOrderLog, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*PurchasehubSupplierOrderLog, error) {
	var ret []*PurchasehubSupplierOrderLog
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) Limit(limit int) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) Offset(offset int) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereIdIn(value []int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereIdNotIn(value []int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderById(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WherePhsoid(p db_repo.Predicate, value int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WherePhsoidIn(value []int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WherePhsoidNotIn(value []int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByPhsoid(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "phsoid "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereErpBillnumber(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "erp_billnumber", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereErpBillnumberIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "erp_billnumber", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereErpBillnumberNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "erp_billnumber", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByErpBillnumber(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "erp_billnumber "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereType(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereTypeIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereTypeNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByType(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperate(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByOperate(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "operate "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateId(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_id", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateIdIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_id", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateIdNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByOperateId(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "operate_id "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateRes(p db_repo.Predicate, value int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_res", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateResIn(value []int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_res", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereOperateResNotIn(value []int32) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "operate_res", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByOperateRes(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "operate_res "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereNote(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereNoteIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereNoteNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "note", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByNote(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "note "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereRawData(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "raw_data", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereRawDataIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "raw_data", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereRawDataNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "raw_data", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByRawData(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "raw_data "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereTimeOut(p db_repo.Predicate, value string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "time_out", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereTimeOutIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "time_out", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereTimeOutNotIn(value []string) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "time_out", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByTimeOut(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "time_out "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereCreatedAt(p db_repo.Predicate, value time.Time) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByCreatedAt(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereUpdatedAt(p db_repo.Predicate, value time.Time) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *purchasehubSupplierOrderLogRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderLogRepoQueryBuilder) OrderByUpdatedAt(asc bool) *purchasehubSupplierOrderLogRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}
