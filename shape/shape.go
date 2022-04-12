package shape

type Point struct {
	X, Y int
}

func (p *Point) Add(a int) {
	p.X += a
	p.Y += a
}

func (p Point) Mul(a int) {
	p.X *= a
	p.Y *= a
}

type Square struct {
	Width  float32
	Height float32
}

type Triangle struct {
	Width  float32
	Height float32
}

func (s Square) Area() float32 {
	return s.Width * s.Height
}

func (t Triangle) Area() float32 {
	return t.Width * t.Height / 2
}

type Line struct {
	Length float32
}

type Figure interface {
	Area() float32
}
