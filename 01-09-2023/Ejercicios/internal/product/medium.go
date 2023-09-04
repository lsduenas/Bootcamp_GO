package product

type ConfigMedium struct {
	Precio float64
}

type Medium struct {
	precio float64
}

func NewMedium(cfg *ConfigMedium) (m *Medium){
	m = &Medium{
		precio: cfg.Precio,
	}
	return
}

func (m *Medium) CalcularPrecio() (precio float64) {
	precio = (m.precio + (m.precio*0.03))
	return
}