package pojo

type Pageable interface {
	Current() int
	QueryCount() int
	Desc() *string
	Aes() *string
}

type Page struct {
	Current    int           `json:"current"`
	QueryCount int           `json:"queryCount"`
	Total      int           `json:"total"`
	Records    []interface{} `json:"records"`
}

type DefaultPageable struct {
	current    int
	queryCount int
	desc       *string
	aes        *string
}

func NewDefaultPageable(current, queryCount int) DefaultPageable {
	return DefaultPageable{current: current, queryCount: queryCount}
}

func (p DefaultPageable) Current() int {
	return p.current
}

func (p DefaultPageable) QueryCount() int {
	return p.queryCount
}

func (p DefaultPageable) Desc() *string {
	return p.desc
}

func (p DefaultPageable) Aes() *string {
	return p.aes
}
