package examples

import (
	"encoding/json"
	"encoding/xml"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// make sure each example contains multiple calls but only one report.

func exampleJSON() {
	var user struct { // want `\Qexported fields should be annotated with the "json" tag`
		Name  string
		Email string `json:"email"`
	}
	json.Marshal(user)
	json.Unmarshal(nil, &user)
}

func exampleXML() {
	var user struct { // want `\Qexported fields should be annotated with the "xml" tag`
		Name  string
		Email string `xml:"email"`
	}
	xml.Marshal(user)
	xml.Unmarshal(nil, &user)
}

func exampleYAML() {
	var user struct { // want `\Qexported fields should be annotated with the "yaml" tag`
		Name  string
		Email string `yaml:"email"`
	}
	yaml.Marshal(user)
	yaml.Unmarshal(nil, &user)
}

func exampleTOML() {
	var user struct { // want `\Qexported fields should be annotated with the "toml" tag`
		Name  string
		Email string `toml:"email"`
	}
	toml.Decode("", &user)
	toml.Unmarshal(nil, &user)
}
