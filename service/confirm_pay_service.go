//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-17 10:04:39
 */
package service

import (
	"cmall/model"
	"cmall/pkg/logging"
	"os"
)

// ConfirmPayService 接收FM支付回调接口
type ConfirmPayService struct {
	MerchantNum     string `form:"merchantNum" json:"merchantNum" `
	OrderNo         string `form:"orderNo" json:"orderNo" `
	PlatformOrderNo string `form:"platformOrderNo" json:"platformOrderNo"`
	Amount          string `form:"amount" json:"amount" `
	ActualPayAmount string `form:"actualPayAmount" json:"actualPayAmount" `
	State           int    `form:"state" json:"state" `
	Attch           string `form:"attch" json:"attch" `
	PayTime         string `form:"payTime" json:"payTime" `
	Sign            string `form:"sign" json:"sign" `
}

// Confirm 接收FM支付回调 详情请查阅FM支付文档
func (service *ConfirmPayService) Confirm() {
	if service.Attch == os.Getenv("FM_Pay_attch") {
		if service.State == 1 {
			if err := model.DB.Model(model.Order{}).Where("order_num=?", service.OrderNo).Update("type", 2).Error; err != nil {
				logging.Info(err)
			}
		}
	}
}
