module tests

go 1.20

require (
	example.com/custom v0.1.0
	github.com/BurntSushi/toml v1.3.2
	github.com/jmoiron/sqlx v1.3.5
	github.com/mitchellh/mapstructure v1.5.0
	gopkg.in/yaml.v3 v3.0.1
)

replace example.com/custom => ./example.com/custom
