module github.com/kosatnkn/catalyst

go 1.15

require (
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.0
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/iancoleman/strcase v0.1.2
	github.com/kosatnkn/db v0.0.0-20210210053948-356124926aaf
	github.com/kosatnkn/log v0.0.0-20210210103509-a1ccb63bd2c1
	github.com/kr/pretty v0.2.1 // indirect
	github.com/prometheus/client_golang v1.8.0
	golang.org/x/sys v0.0.0-20201017003518-b09fb700fbb7 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/kosatnkn/db => /home/kosala/Development/go/github.com/kosatnkn/db

replace github.com/kosatnkn/log => /home/kosala/Development/go/github.com/kosatnkn/log
