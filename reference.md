# Reference
<details><summary><code>client.ServeTerminalTerminalGet() -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Serve the terminal HTML page
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ServeTerminalTerminalGet(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sandbox
<details><summary><code>client.Sandbox.GetContext() -> *sandboxsdkgo.SandboxResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sandbox environment information
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sandbox.GetContext(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandbox.GetPythonPackages() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get installed packages by language
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sandbox.GetPythonPackages(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandbox.GetNodejsPackages() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get installed packages by language
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sandbox.GetNodejsPackages(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Shell
<details><summary><code>client.Shell.ExecCommand(request) -> *sandboxsdkgo.ResponseShellCommandResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute command in the specified shell session
Supports SSE streaming if Accept header contains 'text/event-stream'
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.ExecCommand(
        context.TODO(),
        &sandboxsdkgo.ShellExecRequest{
            Command: "command",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` — Unique identifier of the target shell session, if not provided, one will be automatically created
    
</dd>
</dl>

<dl>
<dd>

**execDir:** `*string` — Working directory for command execution (must use absolute path)
    
</dd>
</dl>

<dl>
<dd>

**command:** `string` — Shell command to execute
    
</dd>
</dl>

<dl>
<dd>

**asyncMode:** `*bool` — Whether to execute command asynchronously (default: False for async, False for synchronous execution)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.View(request) -> *sandboxsdkgo.ResponseShellViewResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

View output of the specified shell session
Supports SSE streaming if Accept header contains 'text/event-stream'
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.View(
        context.TODO(),
        &sandboxsdkgo.ShellViewRequest{
            Id: "id",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier of the target shell session
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.WaitForProcess(request) -> *sandboxsdkgo.ResponseShellWaitResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Wait for the process in the specified shell session to return
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.WaitForProcess(
        context.TODO(),
        &sandboxsdkgo.ShellWaitRequest{
            Id: "id",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier of the target shell session
    
</dd>
</dl>

<dl>
<dd>

**seconds:** `*int` — Wait time (seconds)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.WriteToProcess(request) -> *sandboxsdkgo.ResponseShellWriteResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Write input to the process in the specified shell session
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.WriteToProcess(
        context.TODO(),
        &sandboxsdkgo.ShellWriteToProcessRequest{
            Id: "id",
            Input: "input",
            PressEnter: true,
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier of the target shell session
    
</dd>
</dl>

<dl>
<dd>

**input:** `string` — Input content to write to the process
    
</dd>
</dl>

<dl>
<dd>

**pressEnter:** `bool` — Whether to press enter key after input
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.KillProcess(request) -> *sandboxsdkgo.ResponseShellKillResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Terminate the process in the specified shell session
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.KillProcess(
        context.TODO(),
        &sandboxsdkgo.ShellKillProcessRequest{
            Id: "id",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier of the target shell session
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.CreateSession(request) -> *sandboxsdkgo.ResponseShellCreateSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new shell session and return its ID
If id already exists, return the existing session
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.CreateSession(
        context.TODO(),
        &sandboxsdkgo.ShellCreateSessionRequest{},
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `*string` — Unique identifier for the shell session, auto-generated if not provided
    
</dd>
</dl>

<dl>
<dd>

**execDir:** `*string` — Working directory for the new session (must use absolute path)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.GetTerminalUrl() -> *sandboxsdkgo.ResponseStr</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new shell session and return the terminal URL
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.GetTerminalUrl(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.ListSessions() -> *sandboxsdkgo.ResponseActiveShellSessionsResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all active shell sessions
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.ListSessions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.CleanupAllSessions() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cleanup all active shell sessions
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.CleanupAllSessions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.CleanupSession(SessionId) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Manually cleanup a specific shell session
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Shell.CleanupSession(
        context.TODO(),
        "session_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## File
<details><summary><code>client.File.ReadFile(request) -> *sandboxsdkgo.ResponseFileReadResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read file content
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.ReadFile(
        context.TODO(),
        &sandboxsdkgo.FileReadRequest{
            File: "file",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**file:** `string` — Absolute file path
    
</dd>
</dl>

<dl>
<dd>

**startLine:** `*int` — Start line (0-based)
    
</dd>
</dl>

<dl>
<dd>

**endLine:** `*int` — End line (not inclusive)
    
</dd>
</dl>

<dl>
<dd>

**sudo:** `*bool` — Whether to use sudo privileges
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.WriteFile(request) -> *sandboxsdkgo.ResponseFileWriteResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Write file content (supports both text and binary files)

For binary files, set encoding to 'base64' and provide base64-encoded content.
For text files, use default 'utf-8' encoding.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WriteFile(
        context.TODO(),
        &sandboxsdkgo.FileWriteRequest{
            File: "file",
            Content: "content",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**file:** `string` — Absolute file path
    
</dd>
</dl>

<dl>
<dd>

**content:** `string` — Content to write (text or base64 encoded for binary)
    
</dd>
</dl>

<dl>
<dd>

**encoding:** `*sandboxsdkgo.FileContentEncoding` — Content encoding: utf-8 for text, base64 for binary data
    
</dd>
</dl>

<dl>
<dd>

**append:** `*bool` — Whether to use append mode
    
</dd>
</dl>

<dl>
<dd>

**leadingNewline:** `*bool` — Whether to add leading newline (only for text mode)
    
</dd>
</dl>

<dl>
<dd>

**trailingNewline:** `*bool` — Whether to add trailing newline (only for text mode)
    
</dd>
</dl>

<dl>
<dd>

**sudo:** `*bool` — Whether to use sudo privileges
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.ReplaceInFile(request) -> *sandboxsdkgo.ResponseFileReplaceResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace string in file
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.ReplaceInFile(
        context.TODO(),
        &sandboxsdkgo.FileReplaceRequest{
            File: "file",
            OldStr: "old_str",
            NewStr: "new_str",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**file:** `string` — Absolute file path
    
</dd>
</dl>

<dl>
<dd>

**oldStr:** `string` — Original string to replace
    
</dd>
</dl>

<dl>
<dd>

**newStr:** `string` — New string to replace with
    
</dd>
</dl>

<dl>
<dd>

**sudo:** `*bool` — Whether to use sudo privileges
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.SearchInFile(request) -> *sandboxsdkgo.ResponseFileSearchResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search in file content
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.SearchInFile(
        context.TODO(),
        &sandboxsdkgo.FileSearchRequest{
            File: "file",
            Regex: "regex",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**file:** `string` — Absolute file path
    
</dd>
</dl>

<dl>
<dd>

**regex:** `string` — Regular expression pattern
    
</dd>
</dl>

<dl>
<dd>

**sudo:** `*bool` — Whether to use sudo privileges
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.FindFiles(request) -> *sandboxsdkgo.ResponseFileFindResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Find files by name pattern
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.FindFiles(
        context.TODO(),
        &sandboxsdkgo.FileFindRequest{
            Path: "path",
            Glob: "glob",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**path:** `string` — Directory path to search
    
</dd>
</dl>

<dl>
<dd>

**glob:** `string` — Filename pattern (glob syntax)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.UploadFile(request) -> *sandboxsdkgo.ResponseFileUploadResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload file using streaming
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.UploadFile(
        context.TODO(),
        strings.NewReader(
            "",
        ),
        &sandboxsdkgo.BodyUploadFile{},
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.DownloadFile() -> any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Download file using FileResponse
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.DownloadFile(
        context.TODO(),
        &sandboxsdkgo.FileDownloadFileRequest{
            Path: "path",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**path:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.ListPath(request) -> *sandboxsdkgo.ResponseFileListResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List path contents with flexible options
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.ListPath(
        context.TODO(),
        &sandboxsdkgo.FileListRequest{
            Path: "path",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**path:** `string` — Directory path to list
    
</dd>
</dl>

<dl>
<dd>

**recursive:** `*bool` — Whether to list recursively
    
</dd>
</dl>

<dl>
<dd>

**showHidden:** `*bool` — Whether to show hidden files
    
</dd>
</dl>

<dl>
<dd>

**fileTypes:** `[]string` — Filter by file extensions (e.g., ['.py', '.txt'])
    
</dd>
</dl>

<dl>
<dd>

**maxDepth:** `*int` — Maximum depth for recursive listing
    
</dd>
</dl>

<dl>
<dd>

**includeSize:** `*bool` — Whether to include file size information
    
</dd>
</dl>

<dl>
<dd>

**includePermissions:** `*bool` — Whether to include file permissions
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*string` — Sort by: name, size, modified, type
    
</dd>
</dl>

<dl>
<dd>

**sortDesc:** `*bool` — Sort in descending order
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.StrReplaceEditor(request) -> *sandboxsdkgo.ResponseStrReplaceEditorResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

An filesystem editor tool that allows the agent to
- view
- create
- navigate
- edit files
The tool parameters are defined by Anthropic and are not editable.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.StrReplaceEditor(
        context.TODO(),
        &sandboxsdkgo.StrReplaceEditorRequest{
            Command: sandboxsdkgo.CommandView,
            Path: "path",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**command:** `*sandboxsdkgo.Command` — The commands to run. Allowed options are: `view`, `create`, `str_replace`, `insert`, `undo_edit`.
    
</dd>
</dl>

<dl>
<dd>

**path:** `string` — Absolute path to file or directory, e.g. `/workspace/file.py` or `/workspace`.
    
</dd>
</dl>

<dl>
<dd>

**fileText:** `*string` — Required parameter of `create` command, with the content of the file to be created.
    
</dd>
</dl>

<dl>
<dd>

**oldStr:** `*string` — Required parameter of `str_replace` command containing the string in `path` to replace.
    
</dd>
</dl>

<dl>
<dd>

**newStr:** `*string` — Optional parameter of `str_replace` command containing the new string (if not given, no string will be added). Required parameter of `insert` command containing the string to insert.
    
</dd>
</dl>

<dl>
<dd>

**insertLine:** `*int` — Required parameter of `insert` command. The `new_str` will be inserted AFTER the line `insert_line` of `path`.
    
</dd>
</dl>

<dl>
<dd>

**viewRange:** `[]int` — Optional parameter of `view` command when `path` points to a file. If none is given, the full file is shown. If provided, the file will be shown in the indicated line number range, e.g. [11, 12] will show lines 11 and 12. Indexing at 1 to start. Setting `[start_line, -1]` shows all lines from `start_line` to the end of the file.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jupyter
<details><summary><code>client.Jupyter.ExecuteCode(request) -> *sandboxsdkgo.ResponseJupyterExecuteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute Python code using Jupyter kernel with session persistence

This endpoint allows you to execute Python code and get results back.
You can optionally specify a kernel_name (defaults to 'python3').
Use session_id to maintain variable state across multiple requests.
Sessions automatically expire after 30 minutes of inactivity.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jupyter.ExecuteCode(
        context.TODO(),
        &sandboxsdkgo.JupyterExecuteRequest{
            Code: "code",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` — Python code to execute
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*int` — Execution timeout in seconds
    
</dd>
</dl>

<dl>
<dd>

**kernelName:** `*string` — Kernel name to use (e.g., 'python3', 'python3.11'). Defaults to 'python3'
    
</dd>
</dl>

<dl>
<dd>

**sessionId:** `*string` — Session ID to maintain kernel state across requests
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jupyter.GetInfo() -> *sandboxsdkgo.ResponseJupyterInfoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about available Jupyter kernels
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jupyter.GetInfo(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jupyter.ListSessions() -> *sandboxsdkgo.ResponseActiveSessionsResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all active Jupyter sessions
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jupyter.ListSessions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jupyter.DeleteSessions() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cleanup all active sessions
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jupyter.DeleteSessions(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Jupyter.DeleteSession(SessionId) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Manually cleanup a specific session
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jupyter.DeleteSession(
        context.TODO(),
        "session_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sessionId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Nodejs
<details><summary><code>client.Nodejs.ExecuteCode(request) -> *sandboxsdkgo.ResponseNodeJsExecuteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute JavaScript code using Node.js

This endpoint allows you to execute JavaScript code and get results back.
Each request creates a fresh execution environment that's cleaned up automatically.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Nodejs.ExecuteCode(
        context.TODO(),
        &sandboxsdkgo.NodeJsExecuteRequest{
            Code: "code",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `string` — JavaScript code to execute
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*int` — Execution timeout in seconds
    
</dd>
</dl>

<dl>
<dd>

**stdin:** `*string` — Standard input for the process
    
</dd>
</dl>

<dl>
<dd>

**files:** `map[string]*string` — Additional files to create in execution directory
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Nodejs.GetInfo() -> *sandboxsdkgo.ResponseNodeJsRuntimeInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about Node.js runtime and available languages
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Nodejs.GetInfo(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Mcp
<details><summary><code>client.Mcp.ListMcpTools(ServerName) -> *sandboxsdkgo.ResponseListToolsResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all available tools from the specified MCP server

Args:
    server_name: The name of the MCP server as defined in mcp-servers.json

Returns:
    Response containing the list of available tools with their descriptions and parameters
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Mcp.ListMcpTools(
        context.TODO(),
        "server_name",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serverName:** `string` — Name of the MCP server
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mcp.ExecuteMcpTool(ServerName, ToolName, request) -> *sandboxsdkgo.ResponseCallToolResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a specific tool on the specified MCP server

Args:
    server_name: The name of the MCP server as defined in mcp-servers.json
    tool_name: The name of the tool to execute
    arguments: Tool arguments dictionary

Returns:
    Response containing the tool execution results
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Mcp.ExecuteMcpTool(
        context.TODO(),
        "server_name",
        "tool_name",
        map[string]any{
            "key": "value",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serverName:** `string` — Name of the MCP server
    
</dd>
</dl>

<dl>
<dd>

**toolName:** `string` — Name of the tool to execute
    
</dd>
</dl>

<dl>
<dd>

**request:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Mcp.ListMcpServers() -> *sandboxsdkgo.ResponseListStr</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all configured MCP servers

Returns:
    Response containing the list of configured and filtered MCP servers
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Mcp.ListMcpServers(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Browser
<details><summary><code>client.Browser.GetInfo() -> *sandboxsdkgo.ResponseBrowserInfoResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about browser, like cdp url, viewport size, etc.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Browser.GetInfo(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Browser.ExecuteAction(request) -> *sandboxsdkgo.ActionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a validated action on the current display.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Browser.ExecuteAction(
        context.TODO(),
        &sandboxsdkgo.Action{
            MoveTo: &sandboxsdkgo.MoveToAction{
                X: 1.1,
                Y: 1.1,
            },
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*sandboxsdkgo.Action` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Code
<details><summary><code>client.Code.ExecuteCode(request) -> *sandboxsdkgo.ResponseCodeExecuteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Run code through the unified runtime, dispatching to Python, Node.js, or future language executors
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Code.ExecuteCode(
        context.TODO(),
        &sandboxsdkgo.CodeExecuteRequest{
            Language: sandboxsdkgo.LanguagePython,
            Code: "code",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**language:** `*sandboxsdkgo.Language` — Target runtime language
    
</dd>
</dl>

<dl>
<dd>

**code:** `string` — Source code to execute
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*int` — Execution timeout in seconds
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Code.GetInfo() -> *sandboxsdkgo.ResponseCodeInfoResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return metadata about supported code runtimes
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Code.GetInfo(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Util
<details><summary><code>client.Util.ConvertToMarkdown(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Convert a given URI to Markdown format
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Util.ConvertToMarkdown(
        context.TODO(),
        &sandboxsdkgo.UtilConvertToMarkdownRequest{
            Uri: "uri",
        },
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**uri:** `string` — The URI of the resource to convert
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
