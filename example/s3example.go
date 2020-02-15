package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	port := getenv("APP_PORT", "8081")
	listen := fmt.Sprintf("0.0.0.0:%s", port)

	http.HandleFunc("/", helloWorldHandler)

	log.Print(fmt.Sprintf("Listening on %s", listen))
	log.Fatal(http.ListenAndServe(listen, nil))
}

func getobject() string {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := s3.New(sess)

	object, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("bucketgump"),
		Key:    aws.String("test.json"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	body, _ := ioutil.ReadAll(object.Body)

	defer object.Body.Close()

	return string(body)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, getobject())
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
