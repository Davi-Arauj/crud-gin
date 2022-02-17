package enums

type TipoCliente int

const (
	FISICO = iota + 1
	JURIDICO
	ESPECIAL
)

func (t TipoCliente) TipoClienteString() string {
	return [...]string{"Fisico", "Juridico", "Especial"}[t-1]
}

func (t TipoCliente) TipoClienteInt() int {
	return int(t)
}
