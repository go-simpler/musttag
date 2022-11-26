// Package yaml provides stubs for gopkg.in/yaml.v3.
package yaml

import "io"

func Marshal(_ any) ([]byte, error)   { return nil, nil }
func Unmarshal(_ []byte, _ any) error { return nil }

type Encoder struct{}

func NewEncoder(w io.Writer) *Encoder { return nil }
func (*Encoder) Encode(_ any) error   { return nil }

type Decoder struct{}

func NewDecoder(_ io.Reader) *Decoder { return nil }
func (*Decoder) Decode(_ any) error   { return nil }
