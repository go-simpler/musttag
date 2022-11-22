package testdata

import (
	"encoding/json"
)

func MarshalJSON() {
	type X struct{ NoTag int } // want `exported fields should be annotated with the "json" tag`

	var x X
	json.Marshal(x)
	json.Marshal(&x)
	json.Marshal(X{})
	json.Marshal(&X{})

	var y struct{ NoTag int } // want `exported fields should be annotated with the "json" tag`
	json.Marshal(y)
	json.Marshal(&y)

	json.Marshal(struct{ NoTag int }{})  // want `exported fields should be annotated with the "json" tag`
	json.Marshal(&struct{ NoTag int }{}) // want `exported fields should be annotated with the "json" tag`

	json.MarshalIndent(struct{ NoTag int }{}, "", "")  // want `exported fields should be annotated with the "json" tag`
	json.NewEncoder(nil).Encode(struct{ NoTag int }{}) // want `exported fields should be annotated with the "json" tag`

	json.Marshal(0)
	json.Marshal("")
	json.Marshal(nil)
}

func UnmarshalJSON() {
	empty := []byte("{}")

	type X struct{ NoTag int } // want `exported fields should be annotated with the "json" tag`

	var x X
	json.Unmarshal(empty, &x)
	json.Unmarshal(empty, &X{})

	var y struct{ NoTag int } // want `exported fields should be annotated with the "json" tag`
	json.Unmarshal(empty, &y)

	json.Unmarshal(empty, &struct{ NoTag int }{}) // want `exported fields should be annotated with the "json" tag`

	json.NewDecoder(nil).Decode(&struct{ NoTag int }{}) // want `exported fields should be annotated with the "json" tag`

	json.Unmarshal(empty, &[]int{})
	json.Unmarshal(empty, &map[int]int{})
	json.Unmarshal(empty, nil)
}
