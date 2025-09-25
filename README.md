# Agent Sandbox Go SDK

A Go SDK for the All-in-One Sandbox SDK, giving you typed clients for the sandbox, shell, file, Jupyter, Node.js, and MCP services.

## Installation

```bash
go get github.com/agent-infra/sandbox-sdk-go
```

## Quickstart

```go
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

    ctx := context.Background()
    request := &api.ShellExecRequest{Command: "ls -lat"}

    response, err := c.Shell.ExecCommand(ctx, request)
    if err != nil {
        log.Fatalf("execute command: %v", err)
    }

    fmt.Println(response)
}
```

Run the example with:

```bash
go run examples/basic.go
```

## Cloud Providers

### Volcengine

Pass additional `option.With...` values to `client.NewClient` to set a Volcengine gateway URL, inject custom headers, or tweak transport behaviour. See `examples/basic.go` for a working template.

```go
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
	// For sandbox on volcengine
	c := client.NewClient(
		option.WithBaseURL("https://sd39tjsa51edqmpfaked.apigateway-cn-beijing.volceapi.com"),
		option.WithHTTPHeader(http.Header{
			"x-faas-instance-name": []string{"xxxxx-5jfnyyyyy3il1i3-aaa-faked"},
		}),
	)

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

```

```json
{
  "success": true,
  "message": "Command executed",
  "data": {
    "session_id": "3accc500-1e47-45cb-bebc-b529e462145a",
    "command": "ls -lat",
    "status": "completed",
    "output": "total 56\ndrwxr-xr-x  4 gem  gem  4096 Sep 25 15:38 .local\ndrwxr-x--- 10 gem  gem  4096 Sep 25 15:38 .\ndrwxr-xr-x  3 gem  gem  4096 Sep 25 15:38 .npm\ndrwxr-xr-x  2 gem  gem  4096 Sep 25 15:38 .ipython\ndrwxr-xr-x  1 gem  gem  4096 Sep 25 15:38 .jupyter\ndrwx------  3 gem  gem  4096 Sep 25 15:38 .pki\ndrwxr-xr-x  5 gem  gem  4096 Sep 25 15:38 .cache\ndrwxr-xr-x  6 gem  gem  4096 Sep 25 15:38 .config\n-rw-rw-r--  1 gem  gem     0 Sep 25 15:37 .Xauthority\ndrwxr-xr-x  3 gem  gem  4096 Sep 25 15:37 .npm-global\ndrwxr-xr-x  1 root root 4096 Sep 25 15:37 ..\n-rwxr-xr-x  1 gem  gem    42 Sep 22 23:52 .bashrc\n-rw-r--r--  1 gem  gem   220 Jan  7  2022 .bash_logout\n-rw-r--r--  1 gem  gem   807 Jan  7  2022 .profile",
    "console": [
      {
        "ps1": "$ ",
        "command": "ls -lat",
        "output": "total 56\ndrwxr-xr-x  4 gem  gem  4096 Sep 25 15:38 .local\ndrwxr-x--- 10 gem  gem  4096 Sep 25 15:38 .\ndrwxr-xr-x  3 gem  gem  4096 Sep 25 15:38 .npm\ndrwxr-xr-x  2 gem  gem  4096 Sep 25 15:38 .ipython\ndrwxr-xr-x  1 gem  gem  4096 Sep 25 15:38 .jupyter\ndrwx------  3 gem  gem  4096 Sep 25 15:38 .pki\ndrwxr-xr-x  5 gem  gem  4096 Sep 25 15:38 .cache\ndrwxr-xr-x  6 gem  gem  4096 Sep 25 15:38 .config\n-rw-rw-r--  1 gem  gem     0 Sep 25 15:37 .Xauthority\ndrwxr-xr-x  3 gem  gem  4096 Sep 25 15:37 .npm-global\ndrwxr-xr-x  1 root root 4096 Sep 25 15:37 ..\n-rwxr-xr-x  1 gem  gem    42 Sep 22 23:52 .bashrc\n-rw-r--r--  1 gem  gem   220 Jan  7  2022 .bash_logout\n-rw-r--r--  1 gem  gem   807 Jan  7  2022 .profile"
      }
    ],
    "exit_code": 0
  }
}
```

## Features

- Sandbox metadata: inspect runtime information and installed packages
- Shell access: execute commands with per-client session tracking
- File APIs: read, write, search, and manage files in the sandbox
- Jupyter kernels: run Python statements and notebooks remotely
- Node.js runtime: execute JavaScript in a managed environment
- MCP integration: connect to Model Context Protocol servers

## Requirements

- Go 1.24 or higher
- Network access to an Agent Sandbox API endpoint
