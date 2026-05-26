package zhttp

import (
	"golang.org/x/net/html"
)

type (
	QueryHTML struct {
		node   *html.Node
		filter []*html.Node
	}
	Els []QueryHTML
)

func HTMLParse(HTML []byte) (doc QueryHTML, err error) {
	_ = "STUB: not implemented"
	return *new(QueryHTML), nil
}

func (r *QueryHTML) getNode() *html.Node { _ = "STUB: not implemented"; return nil }

func (r QueryHTML) SelectChild(el string, args ...map[string]string) QueryHTML {
	_ = "STUB: not implemented"
	return *new(QueryHTML)
}

func (r QueryHTML) SelectAllChild(el string, args ...map[string]string) (arr Els) {
	_ = "STUB: not implemented"
	return *new(Els)
}

// Deprecated: please use SelectAllChild("")
// Child All child elements
func (r QueryHTML) Child() (childs []QueryHTML) { _ = "STUB: not implemented"; return nil }

func (r QueryHTML) ForEachChild(f func(index int, child QueryHTML) bool) {
	_ = "STUB: not implemented"
	return
}

func (r QueryHTML) NthChild(index int) QueryHTML { _ = "STUB: not implemented"; return *new(QueryHTML) }

func (r QueryHTML) Select(el string, args ...map[string]string) QueryHTML {
	_ = "STUB: not implemented"
	return *new(QueryHTML)
}

func (r QueryHTML) SelectAll(el string, args ...map[string]string) (arr Els) {
	_ = "STUB: not implemented"
	return *new(Els)
}

func (r QueryHTML) SelectBrother(el string, args ...map[string]string) QueryHTML {
	_ = "STUB: not implemented"
	return *new(QueryHTML)
}

func (r QueryHTML) SelectParent(el string, args ...map[string]string) QueryHTML {
	_ = "STUB: not implemented"
	return *new(QueryHTML)
}

func (r QueryHTML) Find(text string) QueryHTML { _ = "STUB: not implemented"; return *new(QueryHTML) }

func (r QueryHTML) Filter(el ...QueryHTML) QueryHTML {
	_ = "STUB: not implemented"
	return *new(QueryHTML)
}

func parseSelector(text string) []*selector { _ = "STUB: not implemented"; return nil }
