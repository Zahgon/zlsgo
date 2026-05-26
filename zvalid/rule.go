package zvalid

// Regex regular expression match
func (v Engine) Regex(pattern string, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsBool boolean value
func (v Engine) IsBool(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsLower lowerCase letters
func (v Engine) IsLower(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsUpper uppercase letter
func (v Engine) IsUpper(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsLetter uppercase and lowercase letters
func (v Engine) IsLetter(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsNumber is number
func (v Engine) IsNumber(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsInteger is integer
func (v Engine) IsInteger(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsLowerOrDigit lowercase letters or numbers
func (v Engine) IsLowerOrDigit(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsUpperOrDigit uppercase letters or numbers
func (v Engine) IsUpperOrDigit(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsLetterOrDigit letters or numbers
func (v Engine) IsLetterOrDigit(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsChinese chinese character
func (v Engine) IsChinese(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsMobile chinese mobile
func (v Engine) IsMobile(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsMail email address
func (v Engine) IsMail(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsURL links
func (v Engine) IsURL(customError ...string) Engine { _ = "STUB: not implemented"; return *new(Engine) }

// IsIP ipv4 v6 address
func (v Engine) IsIP(customError ...string) Engine { _ = "STUB: not implemented"; return *new(Engine) }

// IsJSON valid json format
func (v Engine) IsJSON(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// IsChineseIDNumber mainland china id number
func (v Engine) IsChineseIDNumber(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

/* 'X' */

// MinLength minimum length
func (v Engine) MinLength(min int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MinUTF8Length utf8 encoding minimum length
func (v Engine) MinUTF8Length(min int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MaxLength the maximum length
func (v Engine) MaxLength(max int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MaxUTF8Length utf8 encoding maximum length
func (v Engine) MaxUTF8Length(max int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MinInt minimum integer value
func (v Engine) MinInt(min int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MaxInt maximum integer value
func (v Engine) MaxInt(max int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MinFloat minimum floating point value
func (v Engine) MinFloat(min float64, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// MaxFloat maximum floating point value
func (v Engine) MaxFloat(max float64, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// EnumString allow only values ​​in []string
func (v Engine) EnumString(slice []string, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// EnumInt allow only values ​​in []int
func (v Engine) EnumInt(i []int, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// EnumFloat64 allow only values ​​in []float64
func (v Engine) EnumFloat64(f []float64, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// CheckPassword check encrypt password
func (v Engine) CheckPassword(password string, customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}
