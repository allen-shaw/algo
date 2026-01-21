package kody

import (
	"errors"
	"text/tabwriter"
)

// 题目名称：设计并实现一个购物车系统
// 要求：
// 添加商品 (Add Item)： 传入用户ID, 商品 ID、单价和数量。如果商品已存在，累加数量。
// 移除商品 (Remove Item)： 传入用户ID, 商品 ID，将其从购物车彻底移除。
// 修改数量 (Update Quantity)： 传入用户ID, 商品 ID 和新数量。
// 计算总价 (Calculate Total)： 传入用户ID, 返回当前购物车内所有商品的总金额。

type Item struct {
	quantity int64
	price int64
}

type UserCart struct {
	items map[string]*Item
	// tatal int64
}

type Cart struct {
	users map[string]*UserCart
}

func (c *Cart) AddItem(userID, itemID string, price int64, quantity int64) error {
	user, ok := c.users[userID]
	if !ok {
		return errors.New("user not existed")
	}
	item, ok := user.items[itemID]
	if !ok {
		item = &Item{
			price: price,
			quantity: 0,
		}
		user.items[itemID] = item
	}
	item.price = price
	item.quantity += quantity
	return nil
}


func (c *Cart) RemoveItem(userID, itemID string) error {
	user, ok := c.users[userID]
	if !ok {
		return errors.New("user not existed")
	}
	delete(user.items, itemID)
	return nil
}

func (c *Cart) UpdateQuantity(userID, itemID string, quantity int64) error {
	user, ok := c.users[userID]
	if !ok {
		return errors.New("user not existed")
	}
	item, ok := user.items[itemID]
	if !ok {
		return errors.New("item not existed")	
	}
	item.quantity = quantity
	return nil
}

func (c *Cart) CalculateTotal(userID string) (int64,error) {
	user, ok := c.users[userID]
	if !ok {
		return 0, errors.New("user not existed")
	}
	items := user.items
	total := int64(0)
	for _, item := range items {
		total += item.price * item.quantity
	}
	return total, nil
}