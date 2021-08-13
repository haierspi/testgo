///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package purchasehub_supplier_order_callback_repo

import (
	"fmt"
	"time"

	"go_service/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PurchasehubSupplierOrderCallback {
	return new(PurchasehubSupplierOrderCallback)
}

func NewQueryBuilder() *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	return new(purchasehubSupplierOrderCallbackRepoQueryBuilder)
}

func (t *PurchasehubSupplierOrderCallback) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type purchasehubSupplierOrderCallbackRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PurchasehubSupplierOrderCallback{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PurchasehubSupplierOrderCallback{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PurchasehubSupplierOrderCallback{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) First(db *gorm.DB) (*PurchasehubSupplierOrderCallback, error) {
	ret := &PurchasehubSupplierOrderCallback{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) QueryOne(db *gorm.DB) (*PurchasehubSupplierOrderCallback, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*PurchasehubSupplierOrderCallback, error) {
	var ret []*PurchasehubSupplierOrderCallback
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) Limit(limit int) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) Offset(offset int) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereIdIn(value []int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereIdNotIn(value []int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) OrderById(asc bool) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WherePhsoid(p db_repo.Predicate, value int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WherePhsoidIn(value []int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WherePhsoidNotIn(value []int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) OrderByPhsoid(asc bool) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "phsoid "+order)
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereType(p db_repo.Predicate, value string) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereTypeIn(value []string) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereTypeNotIn(value []string) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) OrderByType(asc bool) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereStatus(p db_repo.Predicate, value int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereStatusIn(value []int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereStatusNotIn(value []int32) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) OrderByStatus(asc bool) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereCreatedAt(p db_repo.Predicate, value time.Time) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) OrderByCreatedAt(asc bool) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereUpdatedAt(p db_repo.Predicate, value time.Time) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderCallbackRepoQueryBuilder) OrderByUpdatedAt(asc bool) *purchasehubSupplierOrderCallbackRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}