package factory

type Producto interface{
	Usar() string
}

type ProductoA struct{}

func (p ProductoA) Usar() string {
	return "Producto A"
}

type ProductoB struct{}

func (p ProductoB) Usar() string {
	return "Producto B"
}

func CrearProducto(tipo string) Producto {
	if tipo == "A" {
		return ProductoA{}
	} else if tipo == "B" {
		return ProductoB{}
	}

	return nil
}