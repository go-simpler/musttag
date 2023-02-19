package tests

import (
	"encoding/json"
	"encoding/xml"

	// these packages are generated before the tests are run.
	"example.com/custom"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// TODO(junk1tm): drop `reportOnce` and test each builtin function.

func testJSON() {
	var user struct { // want `exported fields should be annotated with the "json" tag`
		Name  string
		Email string `json:"email"`
	}
	json.Marshal(user)
	json.MarshalIndent(user, "", "")
	json.Unmarshal(nil, &user)
	json.NewEncoder(nil).Encode(user)
	json.NewDecoder(nil).Decode(&user)
}

func testXML() {
	var user struct { // want `exported fields should be annotated with the "xml" tag`
		Name  string
		Email string `xml:"email"`
	}
	xml.Marshal(user)
	xml.MarshalIndent(user, "", "")
	xml.Unmarshal(nil, &user)
	xml.NewEncoder(nil).Encode(user)
	xml.NewDecoder(nil).Decode(&user)
	xml.NewEncoder(nil).EncodeElement(user, xml.StartElement{})
	xml.NewDecoder(nil).DecodeElement(&user, &xml.StartElement{})
}

func testYAML() {
	var user struct { // want `exported fields should be annotated with the "yaml" tag`
		Name  string
		Email string `yaml:"email"`
	}
	yaml.Marshal(user)
	yaml.Unmarshal(nil, &user)
	yaml.NewEncoder(nil).Encode(user)
	yaml.NewDecoder(nil).Decode(&user)
}

func testTOML() {
	var user struct { // want `exported fields should be annotated with the "toml" tag`
		Name  string
		Email string `toml:"email"`
	}
	toml.Unmarshal(nil, &user)
	toml.Decode("", &user)
	toml.DecodeFS(nil, "", &user)
	toml.DecodeFile("", &user)
	toml.NewEncoder(nil).Encode(user)
	toml.NewDecoder(nil).Decode(&user)
}

func testMapstructure() {
	var user struct { // want `exported fields should be annotated with the "mapstructure" tag`
		Name  string
		Email string `mapstructure:"email"`
	}
	mapstructure.Decode(nil, &user)
	mapstructure.DecodeMetadata(nil, &user, nil)
	mapstructure.WeakDecode(nil, &user)
	mapstructure.WeakDecodeMetadata(nil, &user, nil)
}

func testCustom() {
	var user struct { // want `exported fields should be annotated with the "custom" tag`
		Name  string
		Email string `custom:"email"`
	}
	custom.Marshal(user)
	custom.Unmarshal(nil, &user)
}
