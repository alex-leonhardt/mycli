module github.com/alex-leonhardt/mycli/cobra_cli/v2

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	mycli v0.0.0-00010101000000-000000000000
)

replace mycli => ./
