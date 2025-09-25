package main

import (
	"context"
	"fmt"
	"log"

	api "github.com/agent-infra/sandbox-sdk-go"
	"github.com/agent-infra/sandbox-sdk-go/client"
	"github.com/agent-infra/sandbox-sdk-go/option"
)

func main() {
	c := client.NewClient(option.WithBaseURL("http://localhost:8091"))

	// For sandbox on volcengine
	// c := client.NewClient(
	// 	option.WithBaseURL("https://sd39tjsa51edqmp4igb00.apigateway-cn-beijing.volceapi.com"),
	// 	option.WithHTTPHeader(http.Header{
	// 		"x-faas-instance-name": []string{"dftck686-5jfn3il1i3-reserved-58b4cb7658-427q6"},
	// 	}),
	// )

	ctx := context.Background()

	request := &api.ShellExecRequest{
		Command: "ls -lat",
	}
	response, err := c.Shell.ExecCommand(ctx, request)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
	fmt.Println(response)
}
