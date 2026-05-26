package zjson

type JSONSyntaxError struct {
	Message  string
	Position int
}

func (e *JSONSyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

type RepairOptions struct {
	AllowComments       bool
	AllowTrailingCommas bool
	AllowSingleQuotes   bool
	AllowUnquotedKeys   bool
}

var defaultRepairOptions = &RepairOptions{
	AllowComments:       true,
	AllowTrailingCommas: true,
	AllowSingleQuotes:   true,
	AllowUnquotedKeys:   true,
}

func Repair(src string, opt ...func(*RepairOptions)) (dst string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func removeComments(src string) string { _ = "STUB: not implemented"; return "" }

type jsonParser struct {
	err       error
	options   *RepairOptions
	container string
	marker    []string
	index     int
}

func newJSONParser(in string, opts *RepairOptions) *jsonParser {
	_ = "STUB: not implemented"
	return nil
}

func (p *jsonParser) parseJSON() interface{} { _ = "STUB: not implemented"; return nil }

func (p *jsonParser) setError(message string) { _ = "STUB: not implemented"; return }

func (p *jsonParser) parseObject() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (p *jsonParser) parseArray() []interface{} { _ = "STUB: not implemented"; return nil }

func (p *jsonParser) parseString() interface{} { _ = "STUB: not implemented"; return nil }

func (p *jsonParser) parseNumber() interface{} { _ = "STUB: not implemented"; return nil }

func (p *jsonParser) parseBooleanOrNull() interface{} { _ = "STUB: not implemented"; return nil }

func (p *jsonParser) tryMatch(target string) bool { _ = "STUB: not implemented"; return false }

func (p *jsonParser) getByte(count int) (byte, bool) { _ = "STUB: not implemented"; return 0, false }

func (p *jsonParser) skipWhitespaces() { _ = "STUB: not implemented"; return }

func (p *jsonParser) setMarker(in string) { _ = "STUB: not implemented"; return }

func (p *jsonParser) resetMarker() { _ = "STUB: not implemented"; return }

func (p *jsonParser) getMarker() string { _ = "STUB: not implemented"; return "" }

func isLetter(c byte) bool { _ = "STUB: not implemented"; return false }

func isNumber(c byte) bool { _ = "STUB: not implemented"; return false }

func toLowerCase(c byte) byte { _ = "STUB: not implemented"; return 0 }
