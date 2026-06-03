package repository

import (
	"context"
	"math"

	"Eccomerce-GO--recap-project-/utils"
	"gorm.io/gorm"
)

// AbstractRepository implements generic CRUD, batching, and pagination patterns.
type AbstractRepository[T any] struct {
	DB *gorm.DB
}

// NewAbstractRepository initializes the generic engine
func NewAbstractRepository[T any](db *gorm.DB) *AbstractRepository[T] {
	return &AbstractRepository[T]{DB: db}
}

// Exists checks if a record matching the filter exists
func (r *AbstractRepository[T]) Exists(ctx context.Context, filter any) (bool, error) {
	var count int64
	var model T
	err := r.DB.WithContext(ctx).Model(&model).Where(filter).Count(&count).Error
	return count > 0, err
}

// FindByID retrieves a single entity by its ID
func (r *AbstractRepository[T]) FindByID(ctx context.Context, id any) (*T, error) {
	var dest T
	err := r.DB.WithContext(ctx).First(&dest, id).Error
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

// FindOne returns the first record matching a flexible filter query (map[string]any or struct)
func (r *AbstractRepository[T]) FindOne(ctx context.Context, filter any) (*T, error) {
	var dest T
	err := r.DB.WithContext(ctx).Where(filter).First(&dest).Error
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

// Find retrieves a collection matching specified filters
func (r *AbstractRepository[T]) Find(ctx context.Context, filter any) ([]T, error) {
	var dest []T
	err := r.DB.WithContext(ctx).Where(filter).Find(&dest).Error
	return dest, err
}

// Create handles a single insertion
func (r *AbstractRepository[T]) Create(ctx context.Context, data *T) (*T, error) {
	err := r.DB.WithContext(ctx).Create(data).Error
	return data, err
}

// CreateMany replicates insertMany batch operations safely [GORM Guides - Create]
func (r *AbstractRepository[T]) CreateMany(ctx context.Context, data []T) ([]T, error) {
	// GORM splits batch insertions automatically into chunks if specified in gorm.Config [GORM Guides - Create]
	err := r.DB.WithContext(ctx).Create(&data).Error
	return data, err
}

// Count returns the direct number of matching documents
func (r *AbstractRepository[T]) Count(ctx context.Context, filter any) (int64, error) {
	var count int64
	var model T
	err := r.DB.WithContext(ctx).Model(&model).Where(filter).Count(&count).Error
	return count, err
}

// FindAndPaginate handles offset limits, parallel count tasks, and structures clean metadata [GORM Guides - Query]
func (r *AbstractRepository[T]) FindAndPaginate(ctx context.Context, filter any, opts *utils.PaginateOptions, sortStr string) (*utils.PaginateResult[T], error) {
	// Replicate default fallback assignment
	limit := 1000
	page := 1
	if opts != nil {
		if opts.Limit > 0 {
			limit = opts.Limit
		}
		if opts.Page > 0 {
			page = opts.Page
		}
	}

	var data []T
	var total int64
	var model T

	// 1. Get the total count for metadata matching criteria
	if err := r.DB.WithContext(ctx).Model(&model).Where(filter).Count(&total).Error; err != nil {
		return nil, err
	}

	// 2. Fetch the paginated dataset [GORM Guides - Query]
	query := r.DB.WithContext(ctx).Where(filter).Limit(limit).Offset((page - 1) * limit)
	if sortStr != "" {
		query = query.Order(sortStr)
	}
	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}

	return &utils.PaginateResult[T]{
		Data: data,
		Meta: r.getMeta(total, len(data), limit, page),
	}, nil
}

// FindOneAndUpdate updates an entity matched by conditional criteria filters [GORM Guides - Update]
func (r *AbstractRepository[T]) FindOneAndUpdate(ctx context.Context, filter any, updateData any) (*T, error) {
	var dest T
	// Updates using mapping criteria and loads the freshly altered record [GORM Guides - Update]
	err := r.DB.WithContext(ctx).Model(&dest).Where(filter).Updates(updateData).First(&dest).Error
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

// Update matches Mongoose findByIdAndUpdate behavior precisely by ID target mapping [GORM Guides - Update]
func (r *AbstractRepository[T]) Update(ctx context.Context, id any, updateData any) (*T, error) {
	var dest T
	err := r.DB.WithContext(ctx).Model(&dest).Where("id = ?", id).Updates(updateData).First(&dest).Error
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

// UpdateMany executes conditional mass global updates, returning a truthy response on structural mutations [GORM Guides - Update]
func (r *AbstractRepository[T]) UpdateMany(ctx context.Context, filter any, updateData any) (bool, error) {
	var model T
	tx := r.DB.WithContext(ctx).Model(&model).Where(filter).Updates(updateData)
	return tx.RowsAffected >= 1, tx.Error
}

// Delete strips records by direct numerical key (Triggers soft-delete automatically if DeletedAt is present) [GORM Guides - Delete]
func (r *AbstractRepository[T]) Delete(ctx context.Context, id any) error {
	var model T
	return r.DB.WithContext(ctx).Delete(&model, id).Error
}

// DeleteMany executes bulk structural purging based on conditions [GORM Guides - Delete]
func (r *AbstractRepository[T]) DeleteMany(ctx context.Context, filter any) (bool, error) {
	var model T
	tx := r.DB.WithContext(ctx).Where(filter).Delete(&model)
	return tx.RowsAffected >= 1, tx.Error
}

// BuildSort replicates the key-value text builder logic for SQL generation
func (r *AbstractRepository[T]) BuildSort(sortKey string, order string) string {
	if sortKey == "" {
		return ""
	}
	if order == "desc" || order == "DESC" {
		return sortKey + " DESC"
	}
	return sortKey + " ASC"
}

// Internal private pagination metadata calculator
func (r *AbstractRepository[T]) getMeta(total int64, count int, limit int, page int) utils.Meta {
	totalPages := 1
	if limit > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
	}
	return utils.Meta{
		TotalItems:   total,
		Count:        count,
		ItemsPerPage: limit,
		CurrentPage:  page,
		TotalPages:   totalPages,
	}
}


func (r *AbstractRepository[T]) WithTransaction(ctx context.Context, fn func(txRepo *AbstractRepository[T]) error) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := NewAbstractRepository[T](tx)
		return fn(txRepo)
	})
}