package bridge

import "fmt"

// abstract interface
type IShape interface {
	Draw()
	SetColor(data IColor)
}

type Rectangle struct {
	IColor
	color string
}
type Round struct {
	IColor
	color string
}
type Triangle struct {
	IColor
	color string
}

func (c *Rectangle) Draw() {
	fmt.Println("绘制"+c.color+"矩形")
}
func (c *Rectangle) SetColor(data IColor){
	c.color = data.GetColor()
}
func (c *Round) SetColor(data IColor){
	c.color = data.GetColor()
}
func (c *Triangle) SetColor(data IColor){
	c.color = data.GetColor()
}
func (c *Round) Draw() {
	fmt.Println("绘制"+c.color+"圆形")
}

func (c *Triangle) Draw() {
	fmt.Println("绘制"+c.color+"三角形")
}
type IColor interface {
	GetColor() string
}

type Red struct {
}
type Blue struct {
}
type Yellow struct {
}
type Green struct {
}

func (r Red) GetColor() string {
	return "红色"
}

func (b Blue) GetColor() string {
	return "蓝色"
}

func (y Yellow) GetColor() string {
	return "蓝色"
}

func (g Green) GetColor() string {
	return "绿色"
}



