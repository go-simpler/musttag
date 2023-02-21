package builtins

import (
	"encoding/json"
	"encoding/xml"

	// these packages are generated before the tests are run.
	"example.com/custom"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type User struct { /* want
	"`User` should be annotated with the `json` tag as it is passed to `json.Marshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `json` tag as it is passed to `json.MarshalIndent` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `json` tag as it is passed to `json.Unmarshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `json` tag as it is passed to `json.Encoder.Encode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `json` tag as it is passed to `json.Decoder.Decode` at testdata/src/builtins/builtins.go"

	"`User` should be annotated with the `xml` tag as it is passed to `xml.Marshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.MarshalIndent` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Unmarshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Encoder.Encode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Decoder.Decode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Encoder.EncodeElement` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Decoder.DecodeElement` at testdata/src/builtins/builtins.go"

	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Marshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Unmarshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Encoder.Encode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Decoder.Decode` at testdata/src/builtins/builtins.go"

	"`User` should be annotated with the `toml` tag as it is passed to `toml.Unmarshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.Decode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.DecodeFS` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.DecodeFile` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.Encoder.Encode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.Decoder.Decode` at testdata/src/builtins/builtins.go"

	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.Decode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.DecodeMetadata` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.WeakDecode` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.WeakDecodeMetadata` at testdata/src/builtins/builtins.go"

	"`User` should be annotated with the `custom` tag as it is passed to `custom.Marshal` at testdata/src/builtins/builtins.go"
	"`User` should be annotated with the `custom` tag as it is passed to `custom.Unmarshal` at testdata/src/builtins/builtins.go"
	*/
	Name  string
	Email string `json:"email" xml:"email" yaml:"email" toml:"email" mapstructure:"email" custom:"email"`
}

func testJSON() {
	var user User
	json.Marshal(user)
	json.MarshalIndent(user, "", "")
	json.Unmarshal(nil, &user)
	json.NewEncoder(nil).Encode(user)
	json.NewDecoder(nil).Decode(&user)
}

func testXML() {
	var user User
	xml.Marshal(user)
	xml.MarshalIndent(user, "", "")
	xml.Unmarshal(nil, &user)
	xml.NewEncoder(nil).Encode(user)
	xml.NewDecoder(nil).Decode(&user)
	xml.NewEncoder(nil).EncodeElement(user, xml.StartElement{})
	xml.NewDecoder(nil).DecodeElement(&user, &xml.StartElement{})
}

func testYAML() {
	var user User
	yaml.Marshal(user)
	yaml.Unmarshal(nil, &user)
	yaml.NewEncoder(nil).Encode(user)
	yaml.NewDecoder(nil).Decode(&user)
}

func testTOML() {
	var user User
	toml.Unmarshal(nil, &user)
	toml.Decode("", &user)
	toml.DecodeFS(nil, "", &user)
	toml.DecodeFile("", &user)
	toml.NewEncoder(nil).Encode(user)
	toml.NewDecoder(nil).Decode(&user)
}

func testMapstructure() {
	var user User
	mapstructure.Decode(nil, &user)
	mapstructure.DecodeMetadata(nil, &user, nil)
	mapstructure.WeakDecode(nil, &user)
	mapstructure.WeakDecodeMetadata(nil, &user, nil)
}

func testCustom() {
	var user User
	custom.Marshal(user)
	custom.Unmarshal(nil, &user)
}
