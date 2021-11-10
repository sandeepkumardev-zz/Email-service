# Email-service
A service where users can register and send an email &amp; do live chat.

###  How To get this working:
- You need to GoLang installed.
- You will need to run go mod download to install dependencies for this project. Do that in the root directory of the project.

## Starting - Local
- Copy .env.example to .env in the same directory
- Remember to create your own .env based on the template provided.
- Update environment variables with credentials for mysql

```sh
go run main.go
```

## Documentation of API's
> Run the app, and browse to http://localhost:3000/swagger/index.html. Find the Swagger 2.0 Api documents