module workspace.dev/shared/go/repositories

go 1.24.4

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/google/uuid v1.6.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.30.0
	workspace.dev/shared/go/models v0.0.0-00010101000000-000000000000
	workspace.dev/shared/go/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace workspace.dev/shared/go/models => ../models

replace workspace.dev/shared/go/utils => ../utils
