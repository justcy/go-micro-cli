package main

import (
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/micro/micro/v3/cmd/protoc-gen-micro/generator"
	_ "github.com/micro/micro/v3/cmd/protoc-gen-micro/plugin/micro"
)

func main() {
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	g := generator.New()
}