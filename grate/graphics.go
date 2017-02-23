package grate

type Image interface {
	Load()
	Draw()
	Translate(float64, float64)
	Scale(float64, float64)
	Rotate(float64)
	Width() float64
	Height() float64
}

type Graphics interface {
	Image(string) Image
	Width() float64
	Height() float64
	Translate(x, y float64)
}


