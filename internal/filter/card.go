package filter

type CardFilter struct {
	Filter map[string][]string
}

func NewCardFilter(codes []string, values []string, suits []string) *CardFilter {
	return &CardFilter{map[string][]string{"code": codes, "value": values, "suit": suits}}
}
