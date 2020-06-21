package sapi

//https://medium.com/chewy-innovation/building-a-serverless-api-with-aws-lambda-5ba30b7e1830

// Payload defines the request from ALB-->Lambda
type Payload struct {
	RequestContext struct {
		Elb struct {
			TargetGroupArn string `json:"targetGroupArn"`
		} `json:"elb"`
	} `json:"requestContext"`
	HTTPMethod            string            `json:"httpMethod"`
	Path                  string            `json:"path"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	Body                  string            `json:"body"`
	IsBase64Encoded       bool              `json:"isBase64Encoded"`
}
