package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shopGo/apps/order/rpc/internal/config"
	"shopGo/apps/order/rpc/model"
)

type ServiceContext struct {
	Config         config.Config
	OrderModel     model.OrdersModel
	OrderItemModel model.OrderitemModel
	ShippingModel  model.ShippingModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		OrderModel:     model.NewOrdersModel(conn),
		OrderItemModel: model.NewOrderitemModel(conn),
		ShippingModel:  model.NewShippingModel(conn),
	}
}
