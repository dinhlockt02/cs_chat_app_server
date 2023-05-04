package common

type Paging[T any] struct {
	Field string `form:"field"`
	Value T      `form:"value"`
	Order string `form:"order"`
	Limit int64  `form:"limit"`
}

var defaultLimit int64 = 20
var descOrder = "desc"
var ascOrder = "asc"

func (p *Paging[T]) MongoProcess() (err error) {
	if p.Limit <= 0 || p.Limit > 50 {
		p.Limit = defaultLimit
	}
	if p.Order != ascOrder {
		p.Order = descOrder
	}
	return err
}
