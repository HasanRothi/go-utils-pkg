package main

import (
	"fmt"
	httputils "github.com/HasanRothi/go-utils-pkg/utils/http"
)

func main() {
	fmt.Println("Go Utils pkg")
	postman := httputils.Postman{
		Url: "https://dummyjson.com/todos/1",
	}
	postman.Get()
	fmt.Println(postman.StatusCode)
	fmt.Println(postman.Response)
}
