package service

import (
	"context"
	"strconv"
	"time"

	"github.com/YiD11/gomall/app/payment/biz/dal/mysql"
	"github.com/YiD11/gomall/app/payment/biz/model"
	payment "github.com/YiD11/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	creditCard := req.CreditCard
	card := creditcard.Card{
		Number: creditCard.Number,
		Cvv:    strconv.Itoa(int(creditCard.Cvv)),
		Month:  strconv.Itoa(int(creditCard.ExpirationMonth)),
		Year:   strconv.Itoa(int(creditCard.ExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}

	tid, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}

	err = model.CreatePayMentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: tid.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}

	return &payment.ChargeResp{
		TransactionId: tid.String(),
	}, nil
}
