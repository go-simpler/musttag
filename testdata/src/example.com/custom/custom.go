// Package custom is an example of using a custom package.
package custom

func Marshal(_ any) ([]byte, error)   { return nil, nil }
func Unmarshal(_ []byte, _ any) error { return nil }
