package tests

import (
	"encoding/json"
	"encoding/xml"

	"example.com/custom"
	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Struct struct{ NoTag string }

type Marshaler struct{ NoTag string }

func (Marshaler) MarshalJSON() ([]byte, error)                               { return nil, nil }
func (*Marshaler) UnmarshalJSON([]byte) error                                { return nil }
func (Marshaler) MarshalXML(e *xml.Encoder, start xml.StartElement) error    { return nil }
func (*Marshaler) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error { return nil }
func (Marshaler) MarshalYAML() (any, error)                                  { return nil, nil }
func (*Marshaler) UnmarshalYAML(*yaml.Node) error                            { return nil }
func (*Marshaler) UnmarshalTOML(any) error                                   { return nil }

type TextMarshaler struct{ NoTag string }

func (TextMarshaler) MarshalText() ([]byte, error) { return nil, nil }
func (*TextMarshaler) UnmarshalText([]byte) error  { return nil }

func testJSON() {
	var st Struct
	json.Marshal(st)                 // want "the given struct should be annotated with the `json` tag"
	json.MarshalIndent(st, "", "")   // want "the given struct should be annotated with the `json` tag"
	json.Unmarshal(nil, &st)         // want "the given struct should be annotated with the `json` tag"
	json.NewEncoder(nil).Encode(st)  // want "the given struct should be annotated with the `json` tag"
	json.NewDecoder(nil).Decode(&st) // want "the given struct should be annotated with the `json` tag"

	var m Marshaler
	json.Marshal(m)
	json.MarshalIndent(m, "", "")
	json.Unmarshal(nil, &m)
	json.NewEncoder(nil).Encode(m)
	json.NewDecoder(nil).Decode(&m)

	var tm TextMarshaler
	json.Marshal(tm)
	json.MarshalIndent(tm, "", "")
	json.Unmarshal(nil, &tm)
	json.NewEncoder(nil).Encode(tm)
	json.NewDecoder(nil).Decode(&tm)
}

func testJSONIndirectSlice() {
	type WithMarshallableSlice struct {
		List []Marshaler `json:"marshallable"`
	}
	var withMarshallableSlice WithMarshallableSlice

	json.Marshal(withMarshallableSlice)
	json.MarshalIndent(withMarshallableSlice, "", "")
	json.NewEncoder(nil).Encode(withMarshallableSlice)
}

func testXML() {
	var st Struct
	xml.Marshal(st)                                             // want "the given struct should be annotated with the `xml` tag"
	xml.MarshalIndent(st, "", "")                               // want "the given struct should be annotated with the `xml` tag"
	xml.Unmarshal(nil, &st)                                     // want "the given struct should be annotated with the `xml` tag"
	xml.NewEncoder(nil).Encode(st)                              // want "the given struct should be annotated with the `xml` tag"
	xml.NewDecoder(nil).Decode(&st)                             // want "the given struct should be annotated with the `xml` tag"
	xml.NewEncoder(nil).EncodeElement(st, xml.StartElement{})   // want "the given struct should be annotated with the `xml` tag"
	xml.NewDecoder(nil).DecodeElement(&st, &xml.StartElement{}) // want "the given struct should be annotated with the `xml` tag"

	var m Marshaler
	xml.Marshal(m)
	xml.MarshalIndent(m, "", "")
	xml.Unmarshal(nil, &m)
	xml.NewEncoder(nil).Encode(m)
	xml.NewDecoder(nil).Decode(&m)
	xml.NewEncoder(nil).EncodeElement(m, xml.StartElement{})
	xml.NewDecoder(nil).DecodeElement(&m, &xml.StartElement{})

	var tm TextMarshaler
	xml.Marshal(tm)
	xml.MarshalIndent(tm, "", "")
	xml.Unmarshal(nil, &tm)
	xml.NewEncoder(nil).Encode(tm)
	xml.NewDecoder(nil).Decode(&tm)
	xml.NewEncoder(nil).EncodeElement(tm, xml.StartElement{})
	xml.NewDecoder(nil).DecodeElement(&tm, &xml.StartElement{})
}

func testYAML() {
	var st Struct
	yaml.Marshal(st)                 // want "the given struct should be annotated with the `yaml` tag"
	yaml.Unmarshal(nil, &st)         // want "the given struct should be annotated with the `yaml` tag"
	yaml.NewEncoder(nil).Encode(st)  // want "the given struct should be annotated with the `yaml` tag"
	yaml.NewDecoder(nil).Decode(&st) // want "the given struct should be annotated with the `yaml` tag"

	var m Marshaler
	yaml.Marshal(m)
	yaml.Unmarshal(nil, &m)
	yaml.NewEncoder(nil).Encode(m)
	yaml.NewDecoder(nil).Decode(&m)
}

func testTOML() {
	var st Struct
	toml.Unmarshal(nil, &st)         // want "the given struct should be annotated with the `toml` tag"
	toml.Decode("", &st)             // want "the given struct should be annotated with the `toml` tag"
	toml.DecodeFS(nil, "", &st)      // want "the given struct should be annotated with the `toml` tag"
	toml.DecodeFile("", &st)         // want "the given struct should be annotated with the `toml` tag"
	toml.NewEncoder(nil).Encode(st)  // want "the given struct should be annotated with the `toml` tag"
	toml.NewDecoder(nil).Decode(&st) // want "the given struct should be annotated with the `toml` tag"

	var m Marshaler
	toml.Unmarshal(nil, &m)
	toml.Decode("", &m)
	toml.DecodeFS(nil, "", &m)
	toml.DecodeFile("", &m)
	toml.NewDecoder(nil).Decode(&m)

	var tm TextMarshaler
	toml.Unmarshal(nil, &tm)
	toml.Decode("", &tm)
	toml.DecodeFS(nil, "", &tm)
	toml.DecodeFile("", &tm)
	toml.NewEncoder(nil).Encode(tm)
	toml.NewDecoder(nil).Decode(&tm)
}

func testMapstructure() {
	var st Struct
	mapstructure.Decode(nil, &st)                  // want "the given struct should be annotated with the `mapstructure` tag"
	mapstructure.DecodeMetadata(nil, &st, nil)     // want "the given struct should be annotated with the `mapstructure` tag"
	mapstructure.WeakDecode(nil, &st)              // want "the given struct should be annotated with the `mapstructure` tag"
	mapstructure.WeakDecodeMetadata(nil, &st, nil) // want "the given struct should be annotated with the `mapstructure` tag"
}

func testSQLX() {
	var st Struct
	sqlx.Get(nil, &st, "")                           // want "the given struct should be annotated with the `db` tag"
	sqlx.GetContext(nil, nil, &st, "")               // want "the given struct should be annotated with the `db` tag"
	sqlx.Select(nil, &st, "")                        // want "the given struct should be annotated with the `db` tag"
	sqlx.SelectContext(nil, nil, &st, "")            // want "the given struct should be annotated with the `db` tag"
	sqlx.StructScan(nil, &st)                        // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Conn).GetContext(nil, &st, "")          // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Conn).SelectContext(nil, &st, "")       // want "the given struct should be annotated with the `db` tag"
	new(sqlx.DB).Get(&st, "")                        // want "the given struct should be annotated with the `db` tag"
	new(sqlx.DB).GetContext(nil, &st, "")            // want "the given struct should be annotated with the `db` tag"
	new(sqlx.DB).Select(&st, "")                     // want "the given struct should be annotated with the `db` tag"
	new(sqlx.DB).SelectContext(nil, &st, "")         // want "the given struct should be annotated with the `db` tag"
	new(sqlx.NamedStmt).Get(&st, nil)                // want "the given struct should be annotated with the `db` tag"
	new(sqlx.NamedStmt).GetContext(nil, &st, nil)    // want "the given struct should be annotated with the `db` tag"
	new(sqlx.NamedStmt).Select(&st, nil)             // want "the given struct should be annotated with the `db` tag"
	new(sqlx.NamedStmt).SelectContext(nil, &st, nil) // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Row).StructScan(&st)                    // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Rows).StructScan(&st)                   // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Stmt).Get(&st)                          // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Stmt).GetContext(nil, &st)              // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Stmt).Select(&st)                       // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Stmt).SelectContext(nil, &st)           // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Tx).Get(&st, "")                        // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Tx).GetContext(nil, &st, "")            // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Tx).Select(&st, "")                     // want "the given struct should be annotated with the `db` tag"
	new(sqlx.Tx).SelectContext(nil, &st, "")         // want "the given struct should be annotated with the `db` tag"
}

func testCustom() {
	var st Struct
	custom.Marshal(st)         // want "the given struct should be annotated with the `custom` tag"
	custom.Unmarshal(nil, &st) // want "the given struct should be annotated with the `custom` tag"
}
