# Scrapper
## Prerequisite
-  go 1.16 or better
-  all the config will be save in `.env` file

## How to run
go to file project directory and run
```
go run main.go
```
open your browser then go to `localhost:3000` (you can change the port number in `.env` file)

## corner cases
- this web app will only check html version from `doctype` declaration
- this web app will return true if there's `input type password` for finding the login form