package product

type ConfigLarge struct{
	Precio float64
}


type Large struct {
	precio float64
}

func NewLarge(cfg *ConfigLarge) (l *Large){
	l = &Large{
		precio: cfg.Precio,
	}
	return
}

func (l *Large) CalcularPrecio()(precio float64){
	precio = l.precio + (l.precio * 0.06) + 2500.0
	return
}