module workspace.dev/services/apis/bookapi

go 1.24.4

require (
	github.com/go-chi/chi/v5 v5.2.1
	github.com/go-playground/validator/v10 v10.26.0
	github.com/google/uuid v1.6.0
	github.com/rs/xid v1.6.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.30.0
	workspace.dev/shared/go/configs v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/errors v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/logger v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/models v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/repositories v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/utils v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/validator v0.0.0-00010101000000-000000000000
)

require (
	github.com/gabriel-vasile/mimetype v1.4.9 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace workspace.dev/shared/go/configs => ../../../shared/go/configs

replace workspace.dev/shared/go/errors => ../../../shared/go/errors

replace workspace.dev/shared/go/logger => ../../../shared/go/logger

replace workspace.dev/shared/go/models => ../../../shared/go/models

replace workspace.dev/shared/go/repositories => ../../../shared/go/repositories

replace workspace.dev/shared/go/utils => ../../../shared/go/utils

replace workspace.dev/shared/go/validator => ../../../shared/go/validator
