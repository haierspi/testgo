///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package purchasehub_supplier_order_sku_merge_purchase_repo

import (
	"fmt"
	"time"

	"go_service/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PurchasehubSupplierOrderSkuMergePurchase {
	return new(PurchasehubSupplierOrderSkuMergePurchase)
}

func NewQueryBuilder() *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	return new(purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder)
}

func (t *PurchasehubSupplierOrderSkuMergePurchase) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PurchasehubSupplierOrderSkuMergePurchase{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PurchasehubSupplierOrderSkuMergePurchase{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PurchasehubSupplierOrderSkuMergePurchase{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) First(db *gorm.DB) (*PurchasehubSupplierOrderSkuMergePurchase, error) {
	ret := &PurchasehubSupplierOrderSkuMergePurchase{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) QueryOne(db *gorm.DB) (*PurchasehubSupplierOrderSkuMergePurchase, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*PurchasehubSupplierOrderSkuMergePurchase, error) {
	var ret []*PurchasehubSupplierOrderSkuMergePurchase
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) Limit(limit int) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) Offset(offset int) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereIdIn(value []int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereIdNotIn(value []int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderById(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePhsoid(p db_repo.Predicate, value int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePhsoidIn(value []int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePhsoidNotIn(value []int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phsoid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderByPhsoid(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "phsoid "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuPurchaseQuantity(p db_repo.Predicate, value string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_purchase_quantity", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuPurchaseQuantityIn(value []string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_purchase_quantity", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuPurchaseQuantityNotIn(value []string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_purchase_quantity", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderBySkuPurchaseQuantity(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sku_purchase_quantity "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuPurchasePrice(p db_repo.Predicate, value string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_purchase_price", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuPurchasePriceIn(value []string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_purchase_price", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuPurchasePriceNotIn(value []string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_purchase_price", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderBySkuPurchasePrice(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sku_purchase_price "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuGoodsDetails(p db_repo.Predicate, value string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_goods_details", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuGoodsDetailsIn(value []string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_goods_details", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereSkuGoodsDetailsNotIn(value []string) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sku_goods_details", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderBySkuGoodsDetails(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sku_goods_details "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePurchaseQuantity(p db_repo.Predicate, value int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "purchase_quantity", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePurchaseQuantityIn(value []int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "purchase_quantity", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePurchaseQuantityNotIn(value []int32) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "purchase_quantity", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderByPurchaseQuantity(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "purchase_quantity "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePurchaseAmount(p db_repo.Predicate, value float64) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "purchase_amount", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePurchaseAmountIn(value []float64) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "purchase_amount", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WherePurchaseAmountNotIn(value []float64) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "purchase_amount", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderByPurchaseAmount(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "purchase_amount "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereCreatedAt(p db_repo.Predicate, value time.Time) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderByCreatedAt(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereUpdatedAt(p db_repo.Predicate, value time.Time) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder) OrderByUpdatedAt(asc bool) *purchasehubSupplierOrderSkuMergePurchaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}