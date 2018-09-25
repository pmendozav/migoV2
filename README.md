# migoV2

Simple Golang project based on clean-architecture approach

## Getting Started

### Prerequisites

* The project uses govendor. To install run in console:
```
go get github.com/kardianos/govendor
```
* Change database parameters on: config/database.yaml (currently this uses bongo db)

### Installing

Download dependencies (ignore errors)
```
govendor sync
```

Build and run
```
go build

./migoV2
```

### Examples

* To create a new user:
```
curl -k -X POST https://0.0.0.0:8080/signup -H "Content-Type: application/json" -d '{"username":"pavel","password":"12345"}'
```

* Login (to generate a jwt token):
```
curl -k -X POST -F 'username=pavel' -F 'password=12345' https://0.0.0.0:8080/login
```

* Find a partner:
```
curl -k -X GET https://0.0.0.0:8080/restricted/partner/id/ongo -H "Authorization: Bearer <USER_TOKEN_HERE>"
```

## Resources

* [clean-architecture](https://erikcaffrey.github.io/ANDROID-clean-architecture/) - Definitions
* [Template](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/) - Template source tutorial based
