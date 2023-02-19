package tests

import "encoding/json"

func namedType() {
	type Foo struct { // want `json.Marshal` `json.Unmarshal` `json.Encoder.Encode` `json.Decoder.Decode`
		NoTag string
	}
	var foo Foo
	json.Marshal(foo)
	json.Unmarshal(nil, &foo)
	json.NewEncoder(nil).Encode(Foo{})
	json.NewDecoder(nil).Decode(&Foo{})
}

func anonymousType() {
	var foo struct { // want `json.Marshal` `json.Unmarshal`
		NoTag string
	}
	json.Marshal(foo)
	json.Unmarshal(nil, &foo)
	json.NewEncoder(nil).Encode(struct{ NoTag int }{})  // want `json.Encoder.Encode`
	json.NewDecoder(nil).Decode(&struct{ NoTag int }{}) // want `json.Decoder.Decode`
}

func nestedType() {
	type Bar struct { // want `json.Marshal` `json.Unmarshal` `json.Encoder.Encode` `json.Decoder.Decode`
		Baz struct{ NoTag string } `json:"baz"`
	}
	type Foo struct {
		Bar ***Bar `json:"bar"`
	}
	var foo Foo
	json.Marshal(foo)
	json.Unmarshal(nil, &foo)
	json.NewEncoder(nil).Encode(Foo{})
	json.NewDecoder(nil).Decode(&Foo{})
}

func embeddedType() {
	type Bar struct { // want `json.Marshal` `json.Unmarshal` `json.Encoder.Encode` `json.Decoder.Decode`
		NoTag string
	}
	type Foo struct {
		Bar
	}
	var foo Foo
	json.Marshal(foo)
	json.Unmarshal(nil, &foo)
	json.NewEncoder(nil).Encode(Foo{})
	json.NewDecoder(nil).Decode(&Foo{})
}

// should not cause panic; see issue #16.
func recursiveType() {
	type Foo struct { // want `json.Marshal` `json.Unmarshal` `json.Encoder.Encode` `json.Decoder.Decode`
		Foo2  *Foo `json:"foo2"`
		NoTag string
	}
	var foo Foo
	json.Marshal(foo)
	json.Unmarshal(nil, &foo)
	json.NewEncoder(nil).Encode(Foo{})
	json.NewDecoder(nil).Decode(&Foo{})
}

func shouldBeIgnored() {
	type Foo struct {
		NoTag int
	}
	var foo Foo
	marshalJSON := json.Marshal
	marshalJSON(foo)  // a non-static call.
	json.Marshal(0)   // a non-struct argument.
	json.Marshal(nil) // nil argument, see issue #20.
}

func nothingToReport() {
	type Bar struct {
		B   string `json:"b"`
		Bar struct {
			C string `json:"c"`
		} `json:"bar"`
	}
	type Foo struct {
		Bar
		A       string `json:"a"`
		private string
		_       string
	}
	var foo Foo
	json.Marshal(foo)
	json.Unmarshal(nil, &foo)
	json.NewEncoder(nil).Encode(Foo{})
	json.NewDecoder(nil).Decode(&Foo{})
}
