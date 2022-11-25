package examples

import (
	"encoding/json"
	"encoding/xml"
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
