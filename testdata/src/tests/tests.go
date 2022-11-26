package tests

import (
	"encoding/json"
	"encoding/xml"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

var xmlSE xml.StartElement

func namedType() {
	type X struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal`
		`\Qjson.Encoder.Encode`
		`\Qjson.Decoder.Decode`

		`\Qxml.Marshal`
		`\Qxml.MarshalIndent`
		`\Qxml.Unmarshal`
		`\Qxml.Encoder.Encode`
		`\Qxml.Decoder.Decode`
		`\Qxml.Encoder.EncodeElement`
		`\Qxml.Decoder.DecodeElement`

		`\Qyaml.v3.Marshal`
		`\Qyaml.v3.Unmarshal`
		`\Qyaml.v3.Encoder.Encode`
		`\Qyaml.v3.Decoder.Decode`

		`\Qtoml.Unmarshal`
		`\Qtoml.Decode`
		`\Qtoml.DecodeFS`
		`\Qtoml.DecodeFile`
		`\Qtoml.Encoder.Encode`
		`\Qtoml.Decoder.Decode` */
		NoTag int
	}
	var x X

	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(X{})
	json.NewDecoder(nil).Decode(&X{})

	xml.Marshal(x)
	xml.MarshalIndent(x, "", "")
	xml.Unmarshal(nil, &x)
	xml.NewEncoder(nil).Encode(X{})
	xml.NewDecoder(nil).Decode(&X{})
	xml.NewEncoder(nil).EncodeElement(X{}, xmlSE)
	xml.NewDecoder(nil).DecodeElement(&X{}, &xmlSE)

	yaml.Marshal(x)
	yaml.Unmarshal(nil, &x)
	yaml.NewEncoder(nil).Encode(X{})
	yaml.NewDecoder(nil).Decode(&X{})

	toml.Unmarshal(nil, &x)
	toml.Decode("", &x)
	toml.DecodeFS(nil, "", &x)
	toml.DecodeFile("", &x)
	toml.NewEncoder(nil).Encode(X{})
	toml.NewDecoder(nil).Decode(&X{})
}

func nestedNamedType() {
	type Y struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal`
		`\Qjson.Encoder.Encode`
		`\Qjson.Decoder.Decode`

		`\Qxml.Marshal`
		`\Qxml.MarshalIndent`
		`\Qxml.Unmarshal`
		`\Qxml.Encoder.Encode`
		`\Qxml.Decoder.Decode`
		`\Qxml.Encoder.EncodeElement`
		`\Qxml.Decoder.DecodeElement`

		`\Qyaml.v3.Marshal`
		`\Qyaml.v3.Unmarshal`
		`\Qyaml.v3.Encoder.Encode`
		`\Qyaml.v3.Decoder.Decode`

		`\Qtoml.Unmarshal`
		`\Qtoml.Decode`
		`\Qtoml.DecodeFS`
		`\Qtoml.DecodeFile`
		`\Qtoml.Encoder.Encode`
		`\Qtoml.Decoder.Decode` */
		NoTag int
	}
	type X struct {
		Y Y `json:"y" xml:"y" yaml:"y" toml:"y"`
	}
	var x X

	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(X{})
	json.NewDecoder(nil).Decode(&X{})

	xml.Marshal(x)
	xml.MarshalIndent(x, "", "")
	xml.Unmarshal(nil, &x)
	xml.NewEncoder(nil).Encode(X{})
	xml.NewDecoder(nil).Decode(&X{})
	xml.NewEncoder(nil).EncodeElement(X{}, xmlSE)
	xml.NewDecoder(nil).DecodeElement(&X{}, &xmlSE)

	yaml.Marshal(x)
	yaml.Unmarshal(nil, &x)
	yaml.NewEncoder(nil).Encode(X{})
	yaml.NewDecoder(nil).Decode(&X{})

	toml.Unmarshal(nil, &x)
	toml.Decode("", &x)
	toml.DecodeFS(nil, "", &x)
	toml.DecodeFile("", &x)
	toml.NewEncoder(nil).Encode(X{})
	toml.NewDecoder(nil).Decode(&X{})
}

func anonymousType() {
	var x struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal`

		`\Qxml.Marshal`
		`\Qxml.MarshalIndent`
		`\Qxml.Unmarshal`

		`\Qyaml.v3.Marshal`
		`\Qyaml.v3.Unmarshal`

		`\Qtoml.Unmarshal`
		`\Qtoml.Decode`
		`\Qtoml.DecodeFS`
		`\Qtoml.DecodeFile` */
		NoTag int
	}

	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(struct{ NoTag int }{})  // want `\Qjson.Encoder.Encode`
	json.NewDecoder(nil).Decode(&struct{ NoTag int }{}) // want `\Qjson.Decoder.Decode`

	xml.Marshal(x)
	xml.MarshalIndent(x, "", "")
	xml.Unmarshal(nil, &x)
	xml.NewEncoder(nil).Encode(struct{ NoTag int }{})                 // want `\Qxml.Encoder.Encode`
	xml.NewDecoder(nil).Decode(&struct{ NoTag int }{})                // want `\Qxml.Decoder.Decode`
	xml.NewEncoder(nil).EncodeElement(struct{ NoTag int }{}, xmlSE)   // want `\Qxml.Encoder.EncodeElement`
	xml.NewDecoder(nil).DecodeElement(&struct{ NoTag int }{}, &xmlSE) // want `\Qxml.Decoder.DecodeElement`

	yaml.Marshal(x)
	yaml.Unmarshal(nil, &x)
	yaml.NewEncoder(nil).Encode(struct{ NoTag int }{})  // want `\Qyaml.v3.Encoder.Encode`
	yaml.NewDecoder(nil).Decode(&struct{ NoTag int }{}) // want `\Qyaml.v3.Decoder.Decode`

	toml.Unmarshal(nil, &x)
	toml.Decode("", &x)
	toml.DecodeFS(nil, "", &x)
	toml.DecodeFile("", &x)
	toml.NewEncoder(nil).Encode(struct{ NoTag int }{})  // want `\Qtoml.Encoder.Encode`
	toml.NewDecoder(nil).Decode(&struct{ NoTag int }{}) // want `\Qtoml.Decoder.Decode`
}

func nestedAnonymousType() {
	var x struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal`

		`\Qxml.Marshal`
		`\Qxml.MarshalIndent`
		`\Qxml.Unmarshal`

		`\Qyaml.v3.Marshal`
		`\Qyaml.v3.Unmarshal`

		`\Qtoml.Unmarshal`
		`\Qtoml.Decode`
		`\Qtoml.DecodeFS`
		`\Qtoml.DecodeFile` */
		Y *struct{ NoTag int } `json:"y" xml:"y" yaml:"y" toml:"y"`
	}

	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(struct{ Y struct{ NoTag int } }{})  // want `\Qjson.Encoder.Encode`
	json.NewDecoder(nil).Decode(&struct{ Y struct{ NoTag int } }{}) // want `\Qjson.Decoder.Decode`

	xml.Marshal(x)
	xml.MarshalIndent(x, "", "")
	xml.Unmarshal(nil, &x)
	xml.NewEncoder(nil).Encode(struct{ Y struct{ NoTag int } }{})                 // want `\Qxml.Encoder.Encode`
	xml.NewDecoder(nil).Decode(&struct{ Y struct{ NoTag int } }{})                // want `\Qxml.Decoder.Decode`
	xml.NewEncoder(nil).EncodeElement(struct{ Y struct{ NoTag int } }{}, xmlSE)   // want `\Qxml.Encoder.EncodeElement`
	xml.NewDecoder(nil).DecodeElement(&struct{ Y struct{ NoTag int } }{}, &xmlSE) // want `\Qxml.Decoder.DecodeElement`

	yaml.Marshal(x)
	yaml.Unmarshal(nil, &x)
	yaml.NewEncoder(nil).Encode(struct{ Y struct{ NoTag int } }{})  // want `\Qyaml.v3.Encoder.Encode`
	yaml.NewDecoder(nil).Decode(&struct{ Y struct{ NoTag int } }{}) // want `\Qyaml.v3.Decoder.Decode`

	toml.Unmarshal(nil, &x)
	toml.Decode("", &x)
	toml.DecodeFS(nil, "", &x)
	toml.DecodeFile("", &x)
	toml.NewEncoder(nil).Encode(struct{ Y struct{ NoTag int } }{})  // want `\Qtoml.Encoder.Encode`
	toml.NewDecoder(nil).Decode(&struct{ Y struct{ NoTag int } }{}) // want `\Qtoml.Decoder.Decode`
}

// all good, nothing to report.
func typeWithAllTags() {
	var x struct {
		Y       int      `json:"y" xml:"y" yaml:"y" toml:"y"`
		Z       int      `json:"z" xml:"z" yaml:"z" toml:"z"`
		Nested  struct{} `json:"nested" xml:"nested" yaml:"nested" toml:"nested"`
		private int
	}

	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(x)
	json.NewDecoder(nil).Decode(&x)

	xml.Marshal(x)
	xml.MarshalIndent(x, "", "")
	xml.Unmarshal(nil, &x)
	xml.NewEncoder(nil).Encode(x)
	xml.NewDecoder(nil).Decode(&x)
	xml.NewEncoder(nil).EncodeElement(x, xmlSE)
	xml.NewDecoder(nil).DecodeElement(&x, &xmlSE)

	yaml.Marshal(x)
	yaml.Unmarshal(nil, &x)
	yaml.NewEncoder(nil).Encode(x)
	yaml.NewDecoder(nil).Decode(&x)

	toml.Unmarshal(nil, &x)
	toml.Decode("", &x)
	toml.DecodeFS(nil, "", &x)
	toml.DecodeFile("", &x)
	toml.NewEncoder(nil).Encode(x)
	toml.NewDecoder(nil).Decode(&x)
}

// non-static calls should be ignored.
func nonStaticCalls() {
	var x struct {
		NoTag int
	}

	marshalJSON := json.Marshal
	marshalJSON(x)

	marshalXML := xml.Marshal
	marshalXML(x)

	marshalYAML := yaml.Marshal
	marshalYAML(x)

	unmarshalTOML := toml.Unmarshal
	unmarshalTOML(nil, &x)
}

// non-struct argument calls should be ignored.
func nonStructArgument() {
	json.Marshal(0)
	json.MarshalIndent("", "", "")
	json.Unmarshal(nil, &[]int{})
	json.NewEncoder(nil).Encode(map[int]int{})
	json.NewDecoder(nil).Decode(&map[int]int{})

	xml.Marshal(0)
	xml.MarshalIndent("", "", "")
	xml.Unmarshal(nil, &[]int{})
	xml.NewEncoder(nil).Encode(map[int]int{})
	xml.NewDecoder(nil).Decode(&map[int]int{})
	xml.NewEncoder(nil).EncodeElement(map[int]int{}, xmlSE)
	xml.NewDecoder(nil).DecodeElement(&map[int]int{}, &xmlSE)

	yaml.Marshal(0)
	yaml.Unmarshal(nil, &[]int{})
	yaml.NewEncoder(nil).Encode(map[int]int{})
	yaml.NewDecoder(nil).Decode(&map[int]int{})

	toml.Unmarshal(nil, &[]int{})
	toml.Decode("", &[]int{})
	toml.DecodeFS(nil, "", &[]int{})
	toml.DecodeFile("", &[]int{})
	toml.NewEncoder(nil).Encode(map[int]int{})
	toml.NewDecoder(nil).Decode(&map[int]int{})
}
