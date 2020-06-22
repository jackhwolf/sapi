# sapi
a simple router to build **S**erverless golang **API**s on AWS application load balancer and lambda

# defining handlers
all handlers must have the signature 
```
func(context.Context, sapi.Payload) *HandlerReturn
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

func main() {
	// create new router with url prefix `/prefix`
	rtr := NewRouter("/prefix/")

	// assign sampleData function to GET /prefix/sample
	rtr.AddRoute(func(ctx context.Context, payload Payload) *HandlerReturn {
		sample := struct {
			Message string
			Time    int
		}{"Hello", 123}
		return &HandlerReturn{sample, 200, nil}
	}, "/sample", http.MethodGet)

	lambda.start(rtr.HandleLambda)
}

```