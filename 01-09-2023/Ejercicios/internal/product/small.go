package product

type ConfigSmall struct {
	Precio float64
}

type Small struct {
	precio float64
}

func NewSmall( cfg *ConfigSmall) (s *Small){
	s = &Small {
		precio: cfg.Precio,
	}
	return
}

func (s Small) CalcularPrecio() (precio float64) {
	precio = s.precio
	return
}