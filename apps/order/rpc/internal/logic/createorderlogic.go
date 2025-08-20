package logic

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"shopGo/apps/order/rpc/internal/svc"
	"shopGo/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	if in.Uid == 0 || in.Pid == 0 {
		logx.Errorf("CreateOrder 参数不合法 uid:%d pid:%d", in.Uid, in.Pid)
		return nil, errors.New("参数不合法")
	}
	oid := genOrderID(time.Now())
	err := l.svcCtx.OrderModel.CreateOrder(l.ctx, oid, in.Uid, in.Pid)
	if err != nil {
		logx.Errorf("CreateOrder OrderModel.CreateOrder oid:%s uid:%s pid:%d", oid, in.Uid, in.Pid)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &order.CreateOrderResponse{}, nil
}

var num int64

func genOrderID(t time.Time) string {
	//s := t.Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s", ms, ps, rs)

	return n
}

func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	if len(m) < n {
		m = strings.Repeat("0", n-len(m)) + m
	}

	return m
}
