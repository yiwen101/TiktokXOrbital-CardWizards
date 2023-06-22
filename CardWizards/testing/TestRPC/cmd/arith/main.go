package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"log"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	"github.com/kitex-contrib/registry-nacos/registry"
	//arithmetic "github.com/yiwen101/CardWizards/TestRPC/kitex_gen/arithmetic"
	//calculator "github.com/yiwen101/CardWizards/TestRPC/kitex_gen/arithmetic/arithmetic"
)

// CalculatorImpl implements the last service interface defined in the IDL.
type CalculatorImpl struct{}

/*
// Add implements the CalculatorImpl interface.
func (s *CalculatorImpl) Add(ctx context.Context, request *arithmetic.Request) (resp *arithmetic.Response, err error) {
	// TODO: Your code here...
	return &arithmetic.Response{FirstArguement: request.FirstArguement, SecondArguement: request.SecondArguement, Result_: request.FirstArguement + request.SecondArguement}, nil
}

// Subtract implements the CalculatorImpl interface.
func (s *CalculatorImpl) Subtract(ctx context.Context, request *arithmetic.Request) (resp *arithmetic.Response, err error) {
	// TODO: Your code here...
	return &arithmetic.Response{FirstArguement: request.FirstArguement, SecondArguement: request.SecondArguement, Result_: request.FirstArguement - request.SecondArguement}, nil
}

// Multiply implements the CalculatorImpl interface.
func (s *CalculatorImpl) Multiply(ctx context.Context, request *arithmetic.Request) (resp *arithmetic.Response, err error) {
	// TODO: Your code here...
	return &arithmetic.Response{FirstArguement: request.FirstArguement, SecondArguement: request.SecondArguement, Result_: request.FirstArguement * request.SecondArguement}, nil
}

// Divide implements the CalculatorImpl interface.
func (s *CalculatorImpl) Divide(ctx context.Context, request *arithmetic.Request) (resp *arithmetic.Response, err error) {
	// TODO: Your code here...
	return &arithmetic.Response{FirstArguement: request.FirstArguement, SecondArguement: request.SecondArguement, Result_: request.FirstArguement / request.SecondArguement}, nil
}

func (s *CalculatorImpl) TestValidator(ctx context.Context, request *arithmetic.TestValidator) (resp *arithmetic.Response, err error) {
	// TODO: Your code here...
	return &arithmetic.Response{FirstArguement: 17, SecondArguement: 17, Result_: 17}, nil
}
*/

type GenericServiceImpl struct {
}

type requestStruct struct {
	FirstArguement  int `json:"firstArguement"`
	SecondArguement int `json:"secondArguement"`
}

type responseStruct struct {
	FirstArguement  int `json:"firstArguement"`
	SecondArguement int `json:"secondArguement"`
	Result          int `json:"result"`
}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	str, err := strconv.Unquote(string(jsonBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var req requestStruct
	err = json.Unmarshal([]byte(str), &req)
	if err != nil {
		return nil, err
	}

	var resp responseStruct
	resp.FirstArguement = req.FirstArguement
	resp.SecondArguement = req.SecondArguement
	resp.Result = req.FirstArguement + req.SecondArguement
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	result := string(respBytes)
	return result, nil

}

func main() {

	p, err := generic.NewThriftFileProvider("arithmetic.thrift", "../../../idl")
	if err != nil {
		panic(err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}

	svc := genericserver.NewServer(
		new(GenericServiceImpl),
		g,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "arithmetic"}),
		server.WithRegistry(r))
	if err != nil {
		panic(err)
	}

	/*
		svr := calculator.NewServer(
			new(CalculatorImpl),

		)
	*/

	err = svc.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
