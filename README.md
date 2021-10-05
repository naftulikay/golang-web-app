# golang-webapp ![Build Status][status.svg]

An example Golang service with batteries included to demonstrate building a full end-to-end web app.

## Libraries

 - Config and CLI
   - [github.com/spf13/cobra][dep-cobra]: Command-line argument parsing.
   - [github.com/spf13/viper][dep-viper]: Configuration file and environment variable parsing.
 - Database
   - [github.com/go-gorm/gorm][dep-gorm]: SQL database object-relational mapper (ORM).
 - Dependency Injection
   - [github.com/google/wire][dep-wire]: Google's dependency injection code generation framework.
 - HTTP
   - [github.com/gorilla/mux][dep-mux]: HTTP router/dispatcher.
   - [github.com/gorilla/handlers][dep-gorilla-handlers]: Common handlers, it is used here for CORS.
 - Observability
   - [github.com/uber-go/zap][dep-zap]: Structured logging.
 - Security
   - [github.com/golang-jwt/jwt][dep-jwt]: JWT token generation/validation with custom claims objects.
   - [github.com/howeyc/gopass][dep-gopass]: Read passwords from standard input.
   - [golang.org/x/crypto/argon2][dep-argon]: Argon 2 key-derivation function for password hashing and validation.
 - Swagger:
   - [github.com/swaggo/http-swagger][dep-http-swagger]: Generic HTTP handlers for serving generated Swagger docs.
   - [github.com/swaggo/swag][dep-swag]: CLI utility for generating Swagger API schema from Go source code annotations.
 - Testing:
    - [github.com/stretchr/testify][dep-testify]: Assertion framework for intuitive test cases.
 - Utilities:
   - [github.com/cenkalti/backoff][dep-backoff]: Retry with a configurable backoff and an optional limit of attempts.
   - [github.com/go-playground/validator][dep-validator]: Tag-based struct validator library.
   - [github.com/hashicorp/go-multierror][dep-multierror]: Return multiple errors within a single error value.
   - [gopkg.in/guregu/null.v4][dep-null]: Null types for representing optional data.

## License

Licensed at your discretion under either

 - [Apache Software License, Version 2.0](./LICENSE-APACHE)
 - [MIT License](./LICENSE-MIT)

 [dep-argon]:            https://pkg.go.dev/golang.org/x/crypto/argon2
 [dep-backoff]:          https://github.com/cenkalti/backoff
 [dep-cobra]:            https://github.com/spf13/cobra
 [dep-gopass]:           https://github.com/howeyc/gopass
 [dep-gorilla-handlers]: https://github.com/gorilla/handlers
 [dep-gorm]:             https://github.com/go-gorm/gorm
 [dep-http-swagger]:     https://github.com/swaggo/http-swagger
 [dep-jwt]:              https://github.com/golang-jwt/jwt
 [dep-multierror]:       https://github.com/hashicorp/go-multierror
 [dep-mux]:              https://github.com/gorilla/mux
 [dep-null]:             https://github.com/guregu/null
 [dep-swag]:             https://github.com/swaggo/swag
 [dep-testify]:          https://github.com/stretchr/testify
 [dep-validator]:        https://github.com/go-playground/validator
 [dep-viper]:            https://github.com/spf13/viper
 [dep-wire]:             https://github.com/google/wire
 [dep-zap]:              https://github.com/uber-go/zap

 [status.svg]: https://github.com/naftulikay/golang-webapp/actions/workflows/golang.yml/badge.svg
