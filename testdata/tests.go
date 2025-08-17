package testdata

import "encoding/json"

func namedType() {
	type Foo struct {
		NoTag string
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func anonymousType() {
	var foo struct {
		NoTag string
	}
	json.Marshal(foo)                    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)                   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(struct{ NoTag int }{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&struct{ NoTag int }{}) // want "the given struct should be annotated with the `json` tag"
}

func nestedType() {
	type Bar struct {
		NoTag string
	}
	type Foo struct {
		Bar Bar `json:"bar"`
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func embeddedType() {
	type Bar struct {
		NoTag string
	}
	type Foo struct {
		Bar
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func nestedArrayType() {
	type Bar struct {
		NoTag string
	}
	type Foo struct {
		Bars [5]Bar `json:"bars"`
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func nestedSliceType() {
	type Bar struct {
		NoTag string
	}
	type Foo struct {
		Bars []Bar `json:"bars"`
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func nestedMapType() {
	type Bar struct {
		NoTag string
	}
	type Foo struct {
		Bars map[string]Bar `json:"bars"`
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func nestedComplexType() {
	type Bar struct {
		NoTag string
	}
	type Foo struct {
		Bars **[][]map[string][][5][5]map[string]*Bar `json:"bars"`
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func recursiveType() {
	// should not cause panic; see issue #16.
	type Foo struct {
		Foo   *Foo `json:"foo"`
		NoTag string
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
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

func nestedTypeWithInterface() {
	type Foo struct {
		Nested Marshaler `json:"nested"`
	}
	var foo Foo
	json.Marshal(foo)
	json.Marshal(&foo)
	json.Marshal(Foo{})
	json.Marshal(&Foo{})
	json.Unmarshal(nil, &foo)
	json.Unmarshal(nil, &Foo{})
}

func ignoredNestedType() {
	type Nested struct {
		NoTag string
	}
	type Foo struct {
		Ignored  Nested `json:"-"`
		Exported string `json:"exported"`
	}
	var foo Foo
	json.Marshal(foo)
	json.Marshal(&foo)
	json.Marshal(Foo{})
	json.Marshal(&Foo{})
}

func ignoredNestedTypeWithSubsequentNoTagField() {
	type Nested struct {
		NoTag string
	}
	type Foo struct {
		Ignored  Nested `json:"-"`
		Exported string `json:"exported"`
		NoTag    string
	}
	var foo Foo
	json.Marshal(foo)    // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&foo)   // want "the given struct should be annotated with the `json` tag"
	json.Marshal(Foo{})  // want "the given struct should be annotated with the `json` tag"
	json.Marshal(&Foo{}) // want "the given struct should be annotated with the `json` tag"
}

func interfaceSliceType() {
	type WithMarshallableSlice struct {
		List []Marshaler `json:"marshallable"`
	}
	var withMarshallableSlice WithMarshallableSlice

	json.Marshal(withMarshallableSlice)
	json.MarshalIndent(withMarshallableSlice, "", "")
	json.NewEncoder(nil).Encode(withMarshallableSlice)
}
