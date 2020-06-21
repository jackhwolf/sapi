package sapi

//https://medium.com/chewy-innovation/building-a-serverless-api-with-aws-lambda-5ba30b7e1830

// Response defines the response from Lambda-->ALB
type Response struct {
	IsBase64Encoded   bool   `json:"isBase64Encoded"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	Headers           struct {
		SetCookie   string `json:"Set-cookie"`
		ContentType string `json:"Content-Type"`
	} `json:"headers"`
	Body string `json:"body"`
}
