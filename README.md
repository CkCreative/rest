# Go REST API example

## Running the project

- Clone this repository
- Make sure you have go installed, together with c compile tools like gcc
- Open the project directory in your terminal and run: `go run main.go`
- Start messing around with the API

## Creating Users

`POST` the following data to `localhost:8080/users`

```json
{
  "name": "Mike Chumba",
  "email": "email@gmail.com",
  "password": "password",
  "gender": "Male"
}
```

## Authenticating Users

`POST` user credentials to `localhost:8080/login`

```json
{
  "email": "email@gmail.com",
  "password": "password"
}
```

You should receive your authentication token.

## Using Auth Token

try to `GET` the `localhost:8080/auth` route and include the following header:

`x-access-token`: `YOUR_TOKEN_HER`
