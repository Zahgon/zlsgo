package zvalid

// Trim remove leading and trailing spaces
func (v Engine) Trim() Engine { _ = "STUB: not implemented"; return *new(Engine) }

// RemoveSpace remove all spaces
func (v Engine) RemoveSpace() Engine { _ = "STUB: not implemented"; return *new(Engine) }

// Replace replace text
func (v Engine) Replace(old, new string, n int) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// ReplaceAll replace all text
func (v Engine) ReplaceAll(old, new string) Engine { _ = "STUB: not implemented"; return *new(Engine) }

// XSSClean clean html tag
func (v Engine) XSSClean() Engine { _ = "STUB: not implemented"; return *new(Engine) }

// SnakeCaseToCamelCase snakeCase To CamelCase: hello_world => helloWorld
func (v Engine) SnakeCaseToCamelCase(ucfirst bool, delimiter ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// CamelCaseToSnakeCase camelCase To SnakeCase helloWorld/HelloWorld => hello_world
func (v Engine) CamelCaseToSnakeCase(delimiter ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// EncryptPassword encrypt the password
func (v Engine) EncryptPassword(cost ...int) Engine { _ = "STUB: not implemented"; return *new(Engine) }
