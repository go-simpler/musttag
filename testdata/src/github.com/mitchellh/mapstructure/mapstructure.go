// Package mapstructure provides stubs for github.com/mitchellh/mapstructure.
package mapstructure

// mapstructure.NewDecoder wants the struct pointer to be passed via
// mapstructure.DecoderConfig, that's why we do not support it for now.

type Metadata struct{}

func Decode(_, _ any) error                          { return nil }
func DecodeMetadata(_, _ any, _ *Metadata) error     { return nil }
func WeakDecode(_, _ any) error                      { return nil }
func WeakDecodeMetadata(_, _ any, _ *Metadata) error { return nil }
