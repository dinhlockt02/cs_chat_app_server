package pchatmdl

type PagingOrder string

var (
	DESC PagingOrder = "desc"
	ASC              = "asc"
)

var defaultLimit int64 = 20
var defaultOrder = DESC

type Paging struct {
	LastId *string `form:"last_id"`
	Order  *string `form:"order"`
	Limit  *int64  `form:"limit"`
}

func (p *Paging) Process() {
	if p.Order == nil || *p.Order != ASC {
		p.Order = (*string)(&defaultOrder)
	}
	if p.Limit == nil || *p.Limit <= 0 || *p.Limit > 50 {
		p.Limit = &defaultLimit
	}
}
