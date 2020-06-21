# sapi
a simple router to build **S**erverless **API**s on AWS application load balancer and lambda

# defining handlers
all handlers must have the signature 
```
func(context.Context, sapi.Payload) (interface{}, int, error)
```
where the `interface{}` should be json.Marshal-able, the `int` is an http status code

# example
```
package main

import (
	"context"
	"github.com/jackhwolf/sapi"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

// a test endpoint to retrieve some sample data
// all endpoints have to have this function signature
func sampleData(ctx context.Context, payload sapi.Payload) (interface{}, int, error) {
	sample := struct {
		Message string
		Time    int
	}{"Hello", 123}
	return sample, 200, nil
}

func main() {
	// create new router with url prefix `/prefix`
	rtr := sapi.NewRouter("/prefix/")

	// assign sampleData function to GET /prefix/sample
	rtr.AddRoute(sampleData, "/sample", http.MethodPost)

	// tell lambda to use this router
	lambda.Start(rtr.HandleLambda)
}

```