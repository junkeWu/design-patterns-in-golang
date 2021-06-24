package builder

// Gun build of complex object
type Gun struct {
	Type string
	Size string
}
// SetType object attribute's function
func (g *Gun) SetType(tp string) {
	g.Type = tp
}
func (g *Gun) SetSize(size string)  {
	g.Size = size
}
func (g *Gun) GetType() string {
	return g.Type
}
func (g *Gun) GetSize() string {
	return g.Size
}

// IGunBuilder builder interface in order to standardise Builder
type IGunBuilder interface {
	SetType(tp string) IGunBuilder
	SetSize(size string) IGunBuilder
	Build() *Gun
}
// GunBuilder concrete builder implement struct
type GunBuilder struct {
	gun *Gun
}
func (b *GunBuilder) SetType(tp string) IGunBuilder {
	if b.gun == nil {
		b.gun = &Gun{}
	}
	switch tp {
	case "AK47":
			tp = "AK47升级版"
	default:
		tp = "手枪"
	}
	b.gun.SetType(tp)
	return b
}

func (b *GunBuilder) SetSize(size string) IGunBuilder {
	if b.gun == nil {
		b.gun = &Gun{}
	}
	b.gun.SetSize(size)
	return b
}

func (b GunBuilder) Build() *Gun {
 	return b.gun
}

// GunDirector in order to standard process of builder build
type GunDirector struct {
	builder IGunBuilder
}

func (gd GunDirector) Create(tp string, size string) *Gun {
	builder := gd.builder.SetSize(size).Build()
	gd.builder.SetType(tp).Build()
	return builder
}
func NewGun(tp, size string) *Gun {
	return GunDirector{builder: &GunBuilder{}}.Create(tp,size)
}

