package nameshake

type Params struct {
	MinNamePrice uint16
}

const (
	DefaultMinPrice uint16 = 1
)

func NewParams(minPrice uint16) Params {
	return Params{minPrice}
}

func DefaultParams() Params {
	return Params{DefaultMinPrice}
}
