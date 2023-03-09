# Invoice Manager

Small API example to handle invoices for a freelancer.

## Technical Improvements

  - Improve the unit test with many transactions to test concurrency

  - Missing versionning of the endpoints
  
  - Missing logging / tracing handler (otel)

  - Missing continous deployment files with Kubes files (helm ?)

  - 

## App Improvements

  - add more features to manipulate invoices and transactions. As example, add the possibility to get many transactions to pay an invoice.

  - User's login via token with an email handler to get forgotten password
  
  - add front-end part (users friendly)
## Librairies Choices

  - [golang-migrate](https://github.com/golang-migrate/migrate) : to run migration properly
  
  - [sqlc](https://github.com/kyleconroy/sqlc) : not an ORM, but a library to implement a type-safe CRUD in Go. It generate Go code interfaces to queries. Faster than the ORM as GORM and got more options than the native sql library.

  - [gin](https://github.com/gin-gonic/gin) : one of the best framework for my needs in this project and i am familiar with. It is using 'httprouter' as router.
