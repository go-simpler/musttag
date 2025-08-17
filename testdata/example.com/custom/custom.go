package custom

func Function(any) ([]byte, error) { return nil, nil }

type Struct struct{}

func (Struct) Method(any) ([]byte, error) { return nil, nil }

type Interface interface{ Method(any) ([]byte, error) }
