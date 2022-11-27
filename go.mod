module github.com/junk1tm/musttag

go 1.18

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/mitchellh/mapstructure v1.5.0
	golang.org/x/tools v0.3.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
)

replace (
	github.com/BurntSushi/toml => ./testdata/src/github.com/BurntSushi/toml
	github.com/mitchellh/mapstructure => ./testdata/src/github.com/mitchellh/mapstructure
	gopkg.in/yaml.v3 => ./testdata/src/gopkg.in/yaml.v3
)
