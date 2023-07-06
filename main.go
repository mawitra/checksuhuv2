package checksuhu

type Suhu struct {
}

type Calculate interface {
	Konvert(number float64) (float64, float64, float64)
}

func NewSuhu() Calculate {
	return &Suhu{}
}

func (s *Suhu) Konvert(c float64) (float64, float64, float64) {

	reamur := (c * 4) / 5

	kelvin := c + 273.15

	fahrenheit := (c*9)/5 + 32

	return reamur, kelvin, fahrenheit
}
