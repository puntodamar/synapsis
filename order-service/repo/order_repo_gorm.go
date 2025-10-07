package repo

import (
	"context"
	"github.com/synapsis/order-service/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderGormRepo struct {
	db *gorm.DB
}

func NewOrderGormRepo(db *gorm.DB) *OrderGormRepo {
	return &OrderGormRepo{db: db}
}

func (r *OrderGormRepo) Create(ctx context.Context, o *domain.Order) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(o).Error; err != nil {
			return err
		}
		if len(o.Items) == 0 {
			return nil
		}
		for i := range o.Items {
			o.Items[i].OrderID = o.ID
		}
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "order_id"}, {Name: "sku"}},
			DoUpdates: clause.AssignmentColumns([]string{"qty"}),
		}).Create(&o.Items).Error
	})
}

func (r *OrderGormRepo) SetStatus(ctx context.Context, id, status string) error {
	res := r.db.Model(&domain.Order{}).
		Where("id = ?", id).
		Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *OrderGormRepo) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	var o domain.Order
	if err := r.db.Preload("Items").
		First(&o, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}
