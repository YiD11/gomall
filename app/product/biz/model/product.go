package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return product, err
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ? ", "%"+q+"%", "%"+q+"%").Error
	return products, err
}

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cli *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  cli,
		prefix:       "shop",
	}
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedRes := c.cacheClient.Get(c.productQuery.ctx, cachedKey)

	err = func() error {
		if err := cachedRes.Err(); err != nil {
			return err
		}
		cachedResByte, err := cachedRes.Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(cachedResByte, &product)
		if err != nil {
			return err
		}
		return nil
	}()

	if err != nil { // cache miss
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}

		enc, err := json.Marshal(product)
		if err != nil {
			return product, err
		}

		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, enc, time.Hour)
	}
	return
}

// just call the original function without caching `ProductQuery.SearchProducts()`
func (c CachedProductQuery) SearchProducts(q string) (products []*Product, err error) {
	return c.productQuery.SearchProducts(q)
}

type ProductMutation struct {
	ctx context.Context
	db  *gorm.DB
	
}
