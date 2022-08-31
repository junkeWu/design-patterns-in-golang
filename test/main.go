package main

import "fmt"

// OrderCreate 定义一个创建订单行为的接口
type OrderCreate interface {
	Do(robot *RobotOrderCreate) error
}

// Param 参数校验
type Param struct{}

func (behavior *Param) Do(robot *RobotOrderCreate) error {
	fmt.Println("参数校验")
	return nil
}

type Address struct{}

func (behavior *Address) Do(robot *RobotOrderCreate) error {
	fmt.Println("地址校验")
	return nil
}

type Order struct{}

func (behavior *Order) Do(robot *RobotOrderCreate) error {
	fmt.Println("写订单信息到数据库")
	return nil
}

type OrderItem struct{}

func (behavior *OrderItem) Do(robot *RobotOrderCreate) error {
	fmt.Println("写订单商品信息到数据库")
	return nil
}

type Request struct{}

type RobotOrderCreate struct {
	InfoUser    interface{}
	InfoProduct interface{}
	InfoAddress interface{}
	InfoCoupon  interface{}
	behaviors   []OrderCreate
}

func (robot *RobotOrderCreate) New(r *Request) *RobotOrderCreate {
	fmt.Println("订单机器人生成初始化订单:", r)
	return robot
}

// RegisterBehavior 注册行为
func (robot *RobotOrderCreate) RegisterBehavior(behavior ...OrderCreate) *RobotOrderCreate {
	robot.behaviors = append(robot.behaviors, behavior...)
	return robot
}

func (robot *RobotOrderCreate) Create() error {
	for _, behavior := range robot.behaviors {
		behavior.Do(robot)
	}
	return nil
}

func main() {
	(&RobotOrderCreate{}).New(&Request{}).
		RegisterBehavior(
			&Param{},
			&Address{},
			&Order{},
			&OrderItem{},
		).Create()
}
