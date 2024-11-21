package service

import (
	"context"

	"github.com/YiD11/gomall/app/order/biz/dal/mysql"
	"github.com/YiD11/gomall/app/order/biz/model"
	order "github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.Items) == 0 {
		err = kerrors.NewBizStatusError(500001, "items is empty")
		return nil, err
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
		}

		if req.Address != nil {
			addr := req.Address
			o.Consignee.StreetAddress = addr.StreetAddress
			o.Consignee.City = addr.City
			o.Consignee.State = addr.State
			o.Consignee.Country = addr.Country
			o.Consignee.ZipCode = addr.ZipCode
		}

		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var items []model.OrderItem
		for _, it := range req.Items {
			items = append(items, model.OrderItem{
				OrderIdRefer: orderId.String(),
				ProductId:    it.Item.ProductId,
				Quantity:     it.Item.Quantity,
				Cost:         it.Cost,
			})
		}

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})

	return
}
