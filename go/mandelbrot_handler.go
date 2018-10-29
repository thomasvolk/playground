package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"

	mandelbrot "./lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m := mandelbrot.Mandelbrot{
		Xmin:       -2.0,
		Ymin:       -2.0,
		Step:       0.01,
		Iterations: 100,
		Width:      400,
		Height:     400,
	}
	image := m.Draw()
	buf := new(bytes.Buffer)
	png.Encode(buf, image)
	headers := make(map[string]string)
	headers["Content-Type"] = "text/html"
	body := fmt.Sprintf(`<!DOCTYPE html>
	<html>
  		<head>
    		<title>Mandelbrot</title>
  		</head>
		<body>
		  <img src="data:image/png;base64, %s" alt="Mandelbrot" />
		</body>
	</html>`, base64.StdEncoding.EncodeToString(buf.Bytes()))
	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            body,
		Headers:         headers,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
