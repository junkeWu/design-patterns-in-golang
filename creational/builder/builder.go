package builder

import (
	"fmt"
)

// complex objects
type MilkTea struct {
	Type string
	Size string
	Pearl bool
	Ice bool
}

func (mk *MilkTea) SetType(tp string){
	mk.Type = tp
}
func (mk *MilkTea) SetSize(size string){
	mk.Size = size
}
func (mk *MilkTea) SetPearl(pearl bool){
	mk.Pearl = pearl
}
func (mk *MilkTea) SetIce(ice bool){
	mk.Ice = ice
}

func (mk *MilkTea) GetType() string{
	return mk.Type
}
func (mk *MilkTea) GetSize() string{
	return mk.Size
}
func (mk *MilkTea) GetPearl() bool{
	return mk.Pearl
}
func (mk *MilkTea) GetIce() bool{
	return mk.Ice
}

// builder interface.
// builder is an abstract thing, which is mainly used to regulate our builders.
type Builder interface {
	SetType(tp string) Builder
	SetSize(size string) Builder
	SetPearl(pearl bool) Builder
	SetIce(ice bool) Builder

	// Build create specified object
	Build() *MilkTea
}
// create object according different needs.
type MilkTeaBuilder struct {
	milkTea *MilkTea
}
func (mtb *MilkTeaBuilder) SetType(tp string) Builder{
	if mtb.milkTea == nil {
		mtb.milkTea = &MilkTea{}
	}
	mtb.milkTea.SetType(tp)
	return mtb
}
func (mtb *MilkTeaBuilder) SetSize(size string) Builder{
	if mtb.milkTea == nil {
		mtb.milkTea = &MilkTea{}
	}
	mtb.milkTea.SetSize(size)
	return mtb
}
func (mtb *MilkTeaBuilder) SetPearl(pearl bool) Builder{
	if mtb.milkTea == nil {
		mtb.milkTea = &MilkTea{}
	}
	mtb.milkTea.SetPearl(pearl)
	return mtb
}
func (mtb *MilkTeaBuilder) SetIce(ice bool) Builder{
	if mtb.milkTea == nil {
		mtb.milkTea = &MilkTea{}
	}
	mtb.milkTea.SetIce(ice)
	return mtb
}
func (mtb *MilkTeaBuilder) Build() *MilkTea{
	return mtb.milkTea
}

type MilkTeaDirector struct {
	builder Builder
}

func (director MilkTeaDirector) Create(tp, size string, pearl, ice bool) *MilkTea{
	return director.builder.SetType(tp).SetSize(size).SetPearl(pearl).SetIce(ice).Build()
}

func (mk *MilkTea) string() string {
	var pearl = "不加珍珠"
	var ice = "不加冰"
	if mk.GetPearl() == true {
		pearl = "加珍珠"
	}
	if mk.GetIce() == true {
		ice = "加冰"
	}
	return fmt.Sprintf("某号客人奶茶是：%s %s %v %v ",mk.GetType(),mk.GetSize(),pearl,ice)
}

func buildMilkTea (tp, size string, pearl, ice bool) *MilkTea {
	builder := &MilkTeaBuilder{}
	milkDirector := &MilkTeaDirector{builder: builder}
	milkTea := milkDirector.Create(tp, size, pearl, ice)
	fmt.Println(milkTea.string())
	return milkTea
}
