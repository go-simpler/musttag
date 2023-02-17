module github.com/junk1tm/musttag

go 1.19

require (
	github.com/junk1tm/musttag/testdata/src v0.0.0
	golang.org/x/tools v0.3.0
)

require (
	example.com/custom v0.0.0 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	example.com/custom => ./testdata/src/example.com/custom
	example.com/examples => ./testdata/src/examples
	github.com/BurntSushi/toml => ./testdata/src/github.com/BurntSushi/toml
	github.com/junk1tm/musttag/testdata/src => ./testdata/src
	github.com/mitchellh/mapstructure => ./testdata/src/github.com/mitchellh/mapstructure
	gopkg.in/yaml.v3 => ./testdata/src/gopkg.in/yaml.v3
)
