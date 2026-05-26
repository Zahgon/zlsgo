package zvalid

// HasLetter must contain letters not case sensitive
func (v Engine) HasLetter(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasLower must contain lowercase letters
func (v Engine) HasLower(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasUpper must contain uppercase letters
func (v Engine) HasUpper(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasNumber must contain numbers
func (v Engine) HasNumber(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasSymbol must contain symbols
func (v Engine) HasSymbol(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasString must contain a specific string
func (v Engine) HasString(sub string, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasPrefix must contain the specified prefix string
func (v Engine) HasPrefix(sub string, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// HasSuffix contains the specified suffix string
func (v Engine) HasSuffix(sub string, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// Password Universal password (any visible character, length between 6 ~ 20)
func (v Engine) Password(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// StrongPassword Strong equal strength password (length is 6 ~ 20, must include uppercase and lowercase letters, numbers and special characters)
func (v Engine) StrongPassword(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}
