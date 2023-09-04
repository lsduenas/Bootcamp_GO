package main

import (
	"ejercicios/internal/product"
	"fmt"
)

const (
	ProSmall  = "small"
	ProMedium = "medium"
	ProLarge  = "large"
)

func main() {
	// dependencies
	// Config product small
	cfg:= &product.ConfigSmall{
		Precio: 34.8, // mayus importante porque viene de otro package !!
	}

	pro1 := newProduct(ProSmall, cfg)

	// calculate price
	fmt.Println(pro1.CalcularPrecio())

	// Config product medium
	cfg2:= &product.ConfigMedium{
		Precio: 50.67,
	}

	pro2 := newProduct(ProMedium, cfg2)

	// calculate price
	fmt.Println(pro2.CalcularPrecio())

	// Config product large
	cfg3:= &product.ConfigLarge{
		Precio: 78.8,
	}

	pro3 := newProduct(ProLarge, cfg3)

	// calculate price
	fmt.Println(pro3.CalcularPrecio())
}

func newProduct(proType string, cfg any) product.Product {
	switch proType {
	case ProSmall:
		{
			// creacion config de producto small
			c := cfg.(*product.ConfigSmall)
			return product.NewSmall(c)
		}
	case ProMedium:
		{
			c := cfg.(*product.ConfigMedium)
			return product.NewMedium(c)
		}
	case ProLarge:
		{
			c := cfg.(*product.ConfigLarge)
			return product.NewLarge(c)
		}
	default:
		return nil
	}
}
