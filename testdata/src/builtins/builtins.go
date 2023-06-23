package builtins

import (
	"encoding/json"
	"encoding/xml"

	"example.com/custom"
	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type User struct { /* want
	"`User` should be annotated with the `json` tag as it is passed to `json.Marshal` at"
	"`User` should be annotated with the `json` tag as it is passed to `json.MarshalIndent` at"
	"`User` should be annotated with the `json` tag as it is passed to `json.Unmarshal` at"
	"`User` should be annotated with the `json` tag as it is passed to `json.Encoder.Encode` at"
	"`User` should be annotated with the `json` tag as it is passed to `json.Decoder.Decode` at"

	"`User` should be annotated with the `xml` tag as it is passed to `xml.Marshal` at"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.MarshalIndent` at"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Unmarshal` at"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Encoder.Encode` at"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Decoder.Decode` at"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Encoder.EncodeElement` at"
	"`User` should be annotated with the `xml` tag as it is passed to `xml.Decoder.DecodeElement` at"

	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Marshal` at"
	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Unmarshal` at"
	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Encoder.Encode` at"
	"`User` should be annotated with the `yaml` tag as it is passed to `yaml.v3.Decoder.Decode` at"

	"`User` should be annotated with the `toml` tag as it is passed to `toml.Unmarshal` at"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.Decode` at"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.DecodeFS` at"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.DecodeFile` at"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.Encoder.Encode` at"
	"`User` should be annotated with the `toml` tag as it is passed to `toml.Decoder.Decode` at"

	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.Decode` at"
	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.DecodeMetadata` at"
	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.WeakDecode` at"
	"`User` should be annotated with the `mapstructure` tag as it is passed to `mapstructure.WeakDecodeMetadata` at"

	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Get` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.GetContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Select` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.SelectContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.StructScan` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Conn.GetContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Conn.SelectContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.DB.Get` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.DB.GetContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.DB.Select` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.DB.SelectContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.NamedStmt.Get` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.NamedStmt.GetContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.NamedStmt.Select` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.NamedStmt.SelectContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Row.StructScan` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Rows.StructScan` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Stmt.Get` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Stmt.GetContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Stmt.Select` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Stmt.SelectContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Tx.Get` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Tx.GetContext` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Tx.Select` at"
	"`User` should be annotated with the `db` tag as it is passed to `sqlx.Tx.SelectContext` at"

	"`User` should be annotated with the `custom` tag as it is passed to `custom.Marshal` at"
	"`User` should be annotated with the `custom` tag as it is passed to `custom.Unmarshal` at"
	*/
	Name  string
	Email string `json:"email" xml:"email" yaml:"email" toml:"email" mapstructure:"email" db:"email" custom:"email"`
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

func testSQLX() {
	var user User
	sqlx.Get(nil, &user, "")
	sqlx.GetContext(nil, nil, &user, "")
	sqlx.Select(nil, &user, "")
	sqlx.SelectContext(nil, nil, &user, "")
	sqlx.StructScan(nil, &user)
	new(sqlx.Conn).GetContext(nil, &user, "")
	new(sqlx.Conn).SelectContext(nil, &user, "")
	new(sqlx.DB).Get(&user, "")
	new(sqlx.DB).GetContext(nil, &user, "")
	new(sqlx.DB).Select(&user, "")
	new(sqlx.DB).SelectContext(nil, &user, "")
	new(sqlx.NamedStmt).Get(&user, nil)
	new(sqlx.NamedStmt).GetContext(nil, &user, nil)
	new(sqlx.NamedStmt).Select(&user, nil)
	new(sqlx.NamedStmt).SelectContext(nil, &user, nil)
	new(sqlx.Row).StructScan(&user)
	new(sqlx.Rows).StructScan(&user)
	new(sqlx.Stmt).Get(&user)
	new(sqlx.Stmt).GetContext(nil, &user)
	new(sqlx.Stmt).Select(&user)
	new(sqlx.Stmt).SelectContext(nil, &user)
	new(sqlx.Tx).Get(&user, "")
	new(sqlx.Tx).GetContext(nil, &user, "")
	new(sqlx.Tx).Select(&user, "")
	new(sqlx.Tx).SelectContext(nil, &user, "")
}

func testCustom() {
	var user User
	custom.Marshal(user)
	custom.Unmarshal(nil, &user)
}
