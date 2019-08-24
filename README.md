# A framework less REST Api created using golang

pair-monitor-go is an attempt to learn to develop a REST Api using golang. 
I have used only libraries instead of frameworks like buffalo/revel. 

## libraries used for

- Test - [net/http/httptest](https://golang.org/pkg/net/http/httptest/), [testing](https://golang.org/pkg/testing/)
- REST Api - [net/http](https://golang.org/pkg/net/http) 
- ORM - [jinzhu/gorm](https://github.com/jinzhu/gorm)
- hot reload - [codegangsta/gin](https://github.com/codegangsta/gin)

## Starting the app
`go run main.go` or 
with hot reload `gin -i run main.go`

## Api endpoints
- record a dev pairing session: `curl -X POST http://localhost:5000/sessions -d 'dev1=Superman&dev2=Batman'`
- get all sessions: `curl -X GET http://localhost:5000/sessions`
