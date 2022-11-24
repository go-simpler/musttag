package testdata

import (
	"encoding/json"
)

func namedType() {
	type X struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal`
		`\Qjson.Encoder.Encode`
		`\Qjson.Decoder.Decode` */
		NoTag int
	}
	var x X
	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(X{})
	json.NewDecoder(nil).Decode(&X{})
}

func nestedNamedType() {
	type Y struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal`
		`\Qjson.Encoder.Encode`
		`\Qjson.Decoder.Decode` */
		NoTag int
	}
	type X struct {
		Y Y `json:"y"`
	}
	var x X
	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(X{})
	json.NewDecoder(nil).Decode(&X{})
}

func anonymousType() {
	var x struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal` */
		NoTag int
	}
	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(struct{ NoTag int }{})  // want `\Qjson.Encoder.Encode`
	json.NewDecoder(nil).Decode(&struct{ NoTag int }{}) // want `\Qjson.Decoder.Decode`
}

func nestedAnonymousType() {
	var x struct { /* want
		`\Qjson.Marshal`
		`\Qjson.MarshalIndent`
		`\Qjson.Unmarshal` */
		Y *struct{ NoTag int } `json:"y"`
	}
	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(struct{ Y struct{ NoTag int } }{})  // want `\Qjson.Encoder.Encode`
	json.NewDecoder(nil).Decode(&struct{ Y struct{ NoTag int } }{}) // want `\Qjson.Decoder.Decode`
}

func typeWithAllTags() {
	type X struct {
		Y       int      `json:"y"`
		Z       int      `json:"z"`
		Nested  struct{} `json:"nested"`
		private int
	}
	var x X
	json.Marshal(x)
	json.MarshalIndent(x, "", "")
	json.Unmarshal(nil, &x)
	json.NewEncoder(nil).Encode(x)
	json.NewDecoder(nil).Decode(&x)
}

func nonStructArgument() {
	json.Marshal(0)
	json.MarshalIndent("", "", "")
	json.Unmarshal(nil, &[]int{})
	json.NewEncoder(nil).Encode(map[int]int{})
	json.NewDecoder(nil).Decode(&map[int]int{})
}
