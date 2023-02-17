module github.com/junk1tm/musttag/testdata/src

go 1.19

require (
	example.com/custom v0.0.0
	github.com/BurntSushi/toml v1.2.1
	github.com/mitchellh/mapstructure v1.5.0
	gopkg.in/yaml.v3 v3.0.1
)

replace (
	example.com/custom => ./example.com/custom
	github.com/BurntSushi/toml => ./github.com/BurntSushi/toml
	github.com/mitchellh/mapstructure => ./github.com/mitchellh/mapstructure
	gopkg.in/yaml.v3 => ./gopkg.in/yaml.v3
)
