module github.com/Abdulrahman-Tayara/notes-app/notes-service

go 1.19

require (
	github.com/Abdulrahman-Tayara/notes-app/pkg v0.0.0-unpublished
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.0
)

replace github.com/Abdulrahman-Tayara/notes-app/pkg v0.0.0-unpublished => ../pkg

replace github.com/Abdulrahman-Tayara/notes-app/proto/authentication v0.0.0-unpublished => ../proto/authentication

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
