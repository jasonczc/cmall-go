/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 15:40:28
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 11:26:43
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/serializer"
)

// CreateCartService 购物车创建的服务
type CreateCartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Create 创建购物车
func (service *CreateCartService) Create() serializer.Response {
	var product model.Products
	var code int
	err := model.DB.First(&product, service.ProductID).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if product == (model.Products{}) {
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var cart model.Carts
	model.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&cart)
	//如果不存在该购物车则创建
	if cart == (model.Carts{}) {
		cart = model.Carts{
			UserID:    service.UserID,
			ProductID: service.ProductID,
			Num:       1,
			MaxNum:    10,
			Check:     false,
		}

		err = model.DB.Create(&cart).Error
		if err != nil {
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Data: serializer.BuildCart(cart, product),
		}
	} else if cart.Num < cart.MaxNum { //如果存在该购物车且num小于maxnum
		cart.Num++
		err = model.DB.Save(&cart).Error
		if err != nil {
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Status: 1,
			Data:   serializer.BuildCart(cart, product),
		}
	} else {
		return serializer.Response{
			Status: 2,
			Msg:    "超过最大上限",
		}
	}
}
