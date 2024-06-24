// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// OrderInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type OrderInfoDao struct {
	gmvc.M                   // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB           // DB is the raw underlying database management object.
	Table   string           // Table is the table name of the DAO.
	Columns orderInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// OrderInfoColumns defines and stores column names for table order_info.
type orderInfoColumns struct {
	Id               string //
	Number           string // 订单编号
	UserId           string // 用户id
	PayType          string // 支付方式 1微信 2支付宝 3云闪付
	CreatedAt        string //
	UpdatedAt        string //
	Remark           string // 备注
	PayAt            string // 支付时间
	Status           string // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价
	ConsigneeName    string // 收货人姓名
	ConsigneePhone   string // 收货人手机号
	ConsigneeAddress string // 收货人详细地址
	Price            string // 订单金额 单位分
	CouponPrice      string // 优惠券金额 单位分
	ActualPrice      string // 实际支付金额 单位分
}

func NewOrderInfoDao() *OrderInfoDao {
	return &OrderInfoDao{
		M:     g.DB("default").Model("order_info").Safe(),
		DB:    g.DB("default"),
		Table: "order_info",
		Columns: orderInfoColumns{
			Id:               "id",
			Number:           "number",
			UserId:           "user_id",
			PayType:          "pay_type",
			CreatedAt:        "created_at",
			UpdatedAt:        "updated_at",
			Remark:           "remark",
			PayAt:            "pay_at",
			Status:           "status",
			ConsigneeName:    "consignee_name",
			ConsigneePhone:   "consignee_phone",
			ConsigneeAddress: "consignee_address",
			Price:            "price",
			CouponPrice:      "coupon_price",
			ActualPrice:      "actual_price",
		},
	}
}
