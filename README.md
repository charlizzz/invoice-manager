# Invoice Manager

Small API example to handle invoices for a freelancer.

## Endpoints

  - "/users" of type GET is returning a list of the first 100 users from the db
  
  - "/invoice" of type POST is returning a HTTP Status code 204, otherwise an HTTP Status error code (500, 400)
    - accept a JSON as input such as : 

    ```json
    {
    "user_id": 21,
    "amount": 113,
    "label": "Work for April"
    }
    ```

  - "/transaction" of type POST is returning a HTTP Status code 204, otherwise an HTTP Status error code (404, 400, 422, 500)
      - accept a JSON as input such as : 

    ```json
    {
    "invoice_id": 42,
    "amount": 956,
    "reference": "JMPINV200220117"
    }
    ```

## Note

For now, on the endpoints "/invoice" and "/transaction", I made the choice to refuse float type number as input. Consistency of the values' type.

## Technical Improvements

  - Have to mock the db to test HTTP API (aim 100% coverage)()

  - Improve the unit test with many transactions to test concurrency

  - Missing versionning of the endpoints
  
  - Missing logging / tracing handler (otel)

  - Missing continous deployment files with Kubes files (helm ?)

## App Improvements

  - add more features to manipulate invoices and transactions. As example, add the possibility to get many transactions to pay an invoice.

  - User's login via token with an email handler to get forgotten password
  
  - add front-end part (users friendly)
## Librairies Choices

  - [golang-migrate](https://github.com/golang-migrate/migrate) : to run migration properly
  
  - [sqlc](https://github.com/kyleconroy/sqlc) : not an ORM, but a library to implement a type-safe CRUD in Go. It generate Go code interfaces to queries. Faster than the ORM as GORM and got more options than the native sql library.

  - [gin](https://github.com/gin-gonic/gin) : one of the best framework for my needs in this project and i am familiar with. It is using 'httprouter' as router.

  - [viper](https://github.com/spf13/viper) : to handle the env config easily.
