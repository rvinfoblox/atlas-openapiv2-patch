package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/infobloxopen/atlas-openapiv2-patch/internal/descriptor"
	genopenapi "github.com/infobloxopen/atlas-openapiv2-patch/pkg/svc"
	"github.com/spf13/pflag"
)

func run(reg *descriptor.Registry, swaggerFiles []string) error {
	for _, file := range swaggerFiles {
		fileName := file
		var f []byte
		var err error
		f, err = ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		responseCodesMap := map[string]int{
			"GET":    reg.GetResponse(),
			"POST":   reg.PostResponse(),
			"PUT":    reg.PutResponse(),
			"PATCH":  reg.PatchResponse(),
			"DELETE": reg.DeleteResponse(),
		}

		fmt.Printf("file = %s, responseCodesMap = %+v\n", fileName, responseCodesMap)
		resp := genopenapi.AtlasSwagger(f, reg.IsWithPrivateOperations(), reg.IsWithCustomAnnotations(),
			responseCodesMap)

		if reg.IsWithPrivateOperations() {
			err = os.Remove(file)
			if err != nil {
				return err
			}
			fileName = strings.Replace(file, ".swagger.json", ".private.swagger.json", -1)
		}
		err = ioutil.WriteFile(fileName, []byte(resp), os.FileMode(0644))
		if err != nil {
			return fmt.Errorf("unable to generate swagger definition")
		}
		glog.V(1).Infof("New OpenAPI file will emit")
	}
	return nil
}

func main() {
	pflag.Parse()
	reg := descriptor.NewRegistry()

	reg.SetPrivateOperations(*withPrivate)
	reg.SetCustomAnnotations(*withCustomAnnotations)
	reg.SetPostResponse(*withPostResponse)
	reg.SetPutResponse(*withPutResponse)
	reg.SetPatchResponse(*withPatchResponse)
	reg.SetDeleteResponse(*withDeleteResponse)
	fmt.Printf("input args: %+v\n", reg)
	glog.V(1).Info("Processing code generator request")

	if len(*swaggerFiles) == 0 {
		glog.Fatal("invalid swagger input files provided")
	}
	err := run(reg, *swaggerFiles)
	if err != nil {
		glog.Fatal(err)
	}
	return
}
