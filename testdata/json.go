package testdata

import (
	"encoding/json"
)

func MarshalJSON() {
	{
		// named type
		type X struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		var x X
		json.Marshal(x)
	}
	{
		// named type + pointer
		type X struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		var x X
		json.Marshal(&x)
	}
	{
		// named type + in-place declaration
		type X struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.Marshal(X{})
	}
	{
		// named type + in-place declaration + pointer
		type X struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.Marshal(&X{})
	}
	{
		// anonymous type
		var x struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.Marshal(x)
	}
	{
		// anonymous type + pointer
		var x struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.Marshal(&x)
	}
	{
		// anonymous type + in-place declaration
		json.Marshal(struct{ Y int }{}) // want `exported fields should be annotated with the "json" tag`
	}
	{
		// anonymous type + in-place declaration + pointer
		json.Marshal(&struct{ Y int }{}) // want `exported fields should be annotated with the "json" tag`
	}
	{
		// marshal indent
		var x struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.MarshalIndent(x, "", "")
	}
	{
		// new encoder
		var x struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.NewEncoder(nil).Encode(x)
	}
	{
		// anonymous nested struct
		var x struct { // want `exported fields should be annotated with the "json" tag`
			Y struct{ Z int } `json:"Y"`
		}
		json.Marshal(x)
	}
	{
		// nested struct of a named type
		type Y struct{ Z int } // want `exported fields should be annotated with the "json" tag`
		type X struct {
			Y Y `json:"y"`
		}
		var x X
		json.Marshal(x)
	}
	{
		// non-struct argument
		json.Marshal(0)
		json.Marshal(nil)
	}
}

func UnmarshalJSON() {
	{
		// named type
		type X struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		var x X
		json.Unmarshal(nil, &x)
	}
	{
		// named type + in-place declaration
		type X struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.Unmarshal(nil, &X{})
	}
	{
		// anonymous type
		var x struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.Unmarshal(nil, &x)
	}
	{
		// anonymous type + in-place declaration
		json.Unmarshal(nil, &struct{ Y int }{}) // want `exported fields should be annotated with the "json" tag`
	}
	{
		// new decoder
		var x struct{ Y int } // want `exported fields should be annotated with the "json" tag`
		json.NewDecoder(nil).Decode(&x)
	}
	{
		// anonymous nested struct
		var x struct { // want `exported fields should be annotated with the "json" tag`
			Y struct{ Z int } `json:"Y"`
		}
		json.Unmarshal(nil, &x)
	}
	{
		// nested struct of a named type
		type Y struct{ Z int } // want `exported fields should be annotated with the "json" tag`
		type X struct {
			Y Y `json:"y"`
		}
		var x X
		json.Unmarshal(nil, &x)
	}
	{
		// non-struct argument
		json.Unmarshal(nil, &[]int{})
		json.Unmarshal(nil, nil)
	}
}
