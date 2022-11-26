// Package toml provides stubs for github.com/BurntSushi/toml.
package toml

import (
	"io"
	"io/fs"
)

func Unmarshal(_ []byte, _ any) error { return nil }

type MetaData struct{}

func Decode(_ string, _ any) (MetaData, error)            { return MetaData{}, nil }
func DecodeFS(_ fs.FS, _ string, _ any) (MetaData, error) { return MetaData{}, nil }
func DecodeFile(_ string, _ any) (MetaData, error)        { return MetaData{}, nil }

type Encoder struct{}

func NewEncoder(_ io.Writer) *Encoder { return nil }
func (*Encoder) Encode(_ any) error   { return nil }

type Decoder struct{}

func NewDecoder(_ io.Reader) *Decoder { return nil }
func (*Decoder) Decode(_ any) error   { return nil }
