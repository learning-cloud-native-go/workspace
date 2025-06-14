module workspace.dev/tools/db/dbmigrate

go 1.24.4

require (
	github.com/jackc/pgx/v5 v5.7.5
	github.com/pressly/goose/v3 v3.24.3
	workspace.dev/shared/go/configs v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace workspace.dev/shared/go/configs => ../../../shared/go/configs
