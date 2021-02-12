# Welcome to Test

To run this project:
```
go run main.go
```
Methods:
### POST 
localhost:8080/api/transaction
input body: 
{
	"type": <"c" or "d">,
	"amount": < float >
}

example:
{
	"type": "d",
	"amount":1500.00
}

### GET
localhost:8080/api/getBalance

localhost:8080/api/getHistory

To generate a binary:
```
go build main.go
```

To test /api/transaction inside api folder run :
```
go test
```
