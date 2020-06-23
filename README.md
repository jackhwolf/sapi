# sapi
a simple router to build **S**erverless golang **API**s on AWS application load balancer and lambda

# defining handlers
all handlers must have the signature 
```
func(context.Context, sapi.Payload) *HandlerReturn
```
where `HandlerReturn` is defined as 
```
type HandlerReturn struct {
	Body       interface{}
	StatusCode int
	Err        error	
}
```

# example
Check out the (examples repo)[https://github.com/jackhwolf/sapi-examples]
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
	rtr := sapi.NewRouter("/prefix/")

	// assign sampleData function to GET /prefix/sample
	rtr.AddRoute(func(ctx context.Context, payload Payload) *sapi.HandlerReturn {
		sample := struct {
			Message string
			Time    int
		}{"Hello", 123}
		return &sapi.HandlerReturn{sample, 200, nil}
	}, "/sample", http.MethodGet)

	lambda.start(rtr.HandleLambda)
}

```