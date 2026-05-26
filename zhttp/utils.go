package zhttp

import (
	"net/http"

	"golang.org/x/net/html"
)

type (
	selector struct {
		Attr    map[string]string
		Name    string
		i       int
		Child   bool
		Brother bool
	}
)

func (s *selector) appendAttr(key, val string, index int) { _ = "STUB: not implemented"; return }

// ConvertCookie Parse Cookie String
func ConvertCookie(cookiesRaw string) map[string]*http.Cookie {
	_ = "STUB: not implemented"
	return nil
}

// BodyJSON make the object be encoded in json format and set it to the request body
func BodyJSON(v interface{}) *bodyJson { _ = "STUB: not implemented"; return nil }

// BodyXML make the object be encoded in xml format and set it to the request body
func BodyXML(v interface{}) *bodyXml { _ = "STUB: not implemented"; return nil }

func File(path string, field ...string) interface{} { _ = "STUB: not implemented"; return nil }

var UserAgentLists = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Linux; U; Android 2.3.6; zh-cn; GT-S5660 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MicroMessenger/4.5.255",
	"Mozilla/5.0 (X11; OpenBSD i386) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1944.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 5.1) Gecko/20100101 Firefox/14.0 Opera/12.0",
	"Mozilla/5.0 (compatible; Googlebot/2.1;+http://www.google.com/bot.html)",
}

func RandomUserAgent() Header { _ = "STUB: not implemented"; return *new(Header) }

func matchElName(n *html.Node, name string) bool { _ = "STUB: not implemented"; return false }

func arr2Attr(args []map[string]string) map[string][]string { _ = "STUB: not implemented"; return nil }

func getAttrValue(attributes []html.Attribute) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func findAttrValue(attr html.Attribute, attribute string, value []string) bool {
	_ = "STUB: not implemented"
	return false
}

// todo optimization

func getElText(r QueryHTML, full bool) string { _ = "STUB: not implemented"; return "" }

func forChild(node *html.Node, fn func(n *html.Node) bool) { _ = "STUB: not implemented"; return }

func matchEl(n *html.Node, el string, args map[string][]string) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func findChild(node *html.Node, el string, args []map[string]string, multiple bool) (elArr []*html.Node) {
	_ = "STUB: not implemented"
	return nil
}
