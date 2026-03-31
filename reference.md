# Reference
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

Get installed Python packages
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

Get installed Node.js packages
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

<details><summary><code>client.Sandbox.ListHooks() -> *sandboxsdkgo.ResponseListSandboxHook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List registered lifecycle hooks, optionally filtered by event.
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
client.Sandbox.ListHooks(
        context.TODO(),
        &sandboxsdkgo.SandboxListHooksRequest{
            Event: sandboxsdkgo.String(
                "event",
            ),
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

**event:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandbox.RegisterHook(request) -> *sandboxsdkgo.ResponseSandboxHook</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Register a lifecycle hook. Currently supported events: shutdown.
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
client.Sandbox.RegisterHook(
        context.TODO(),
        &sandboxsdkgo.RegisterHookRequest{
            Name: "name",
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

**name:** `string` — Unique name for this hook
    
</dd>
</dl>

<dl>
<dd>

**event:** `*string` — Lifecycle event: "shutdown"
    
</dd>
</dl>

<dl>
<dd>

**command:** `string` — Shell command to execute
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*float64` — Per-hook timeout in seconds
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*int` — Execution priority (lower = earlier). Same priority hooks run in parallel
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandbox.RemoveHook(Name) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a hook by name. ENV hooks cannot be removed.
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
client.Sandbox.RemoveHook(
        context.TODO(),
        "name",
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

**name:** `string` 
    
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

<dl>
<dd>

**timeout:** `*float64` — Maximum time (seconds) to wait for command completion before returning running status
    
</dd>
</dl>

<dl>
<dd>

**strict:** `*bool` — Strict mode for working directory validation. If True, returns error when working directory does not exist. If False or None, silently falls back to session working directory.
    
</dd>
</dl>

<dl>
<dd>

**noChangeTimeout:** `*int` — Timeout (seconds) for detecting no new output from a command. If no output change is detected within this time, command returns with NO_CHANGE_TIMEOUT status. Overrides session-level setting for this command only.
    
</dd>
</dl>

<dl>
<dd>

**hardTimeout:** `*float64` — Hard timeout (seconds) for command execution. When reached, the command is forcefully stopped and current console output is returned with HARD_TIMEOUT status. Unlike timeout (which only affects HTTP response timing), this actually terminates the command.
    
</dd>
</dl>

<dl>
<dd>

**preserveSymlinks:** `*bool` — If True, preserve symlinks in working directory path (pwd shows symlink path). If False, symlinks are resolved to physical paths. Defaults to False for backward compatibility.
    
</dd>
</dl>

<dl>
<dd>

**truncate:** `*bool` — If True, truncate output when it exceeds 30000 characters (default: True)
    
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

<dl>
<dd>

**maxWaitSeconds:** `*int` — Maximum wait time (seconds) for the command to complete
    
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

<dl>
<dd>

**noChangeTimeout:** `*int` — Timeout (seconds) for detecting no new output from commands in this session. Default is 120 seconds. If no output change is detected within this time, command returns with NO_CHANGE_TIMEOUT status.
    
</dd>
</dl>

<dl>
<dd>

**preserveSymlinks:** `*bool` — If True, preserve symlinks in working directory path (pwd shows symlink path). If False, symlinks are resolved to physical paths. Defaults to False for backward compatibility.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Shell.UpdateSession(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update shell session configuration (e.g., no_change_timeout)
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
client.Shell.UpdateSession(
        context.TODO(),
        &sandboxsdkgo.ShellUpdateSessionRequest{
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

**noChangeTimeout:** `*int` — New timeout (seconds) for detecting no new output from commands. If no output change is detected within this time, command returns with NO_CHANGE_TIMEOUT status.
    
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

<details><summary><code>client.Shell.GetSessionStats() -> *sandboxsdkgo.ResponseShellSessionStats</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return aggregate statistics for shell sessions.
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
client.Shell.GetSessionStats(
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

## Bash
<details><summary><code>client.Bash.Exec(request) -> *sandboxsdkgo.ResponseBashExecResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute a bash command.

- async_mode=true: returns immediately with RUNNING status
- async_mode=false + timeout: waits up to timeout, then returns RUNNING if not done
- async_mode=false + no timeout: waits until command completes
- hard_timeout: kills the process after this many seconds
- env: extra environment variables for this command only
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
client.Bash.Exec(
        context.TODO(),
        &sandboxsdkgo.BashExecRequest{
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

**sessionId:** `*string` — Target session ID. If not provided, a new session is created automatically. Reuse the same session_id to maintain state (env vars, cwd, etc.) across commands.
    
</dd>
</dl>

<dl>
<dd>

**command:** `string` — Shell command to execute
    
</dd>
</dl>

<dl>
<dd>

**execDir:** `*string` — Working directory (absolute path). Takes effect on every call — if the session already exists, the working directory is updated persistently.
    
</dd>
</dl>

<dl>
<dd>

**env:** `map[string]*string` — Extra environment variables to inject for this command only.
    
</dd>
</dl>

<dl>
<dd>

**asyncMode:** `*bool` — If true, return immediately with running status. Use /output to poll results.
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*float64` — HTTP timeout (seconds). Only effective when async_mode=false. If the command does not complete within this time, HTTP returns running status and the command continues in the background. Use /output to get results.
    
</dd>
</dl>

<dl>
<dd>

**hardTimeout:** `*float64` — Hard execution timeout (seconds). When reached, the process is killed and status becomes timed_out. None means no limit.
    
</dd>
</dl>

<dl>
<dd>

**maxOutputLength:** `*int` — Maximum character length for stdout/stderr in the response. When output exceeds this limit, middle truncation is applied (head and tail preserved, middle replaced with a marker). Only effective in sync mode (async_mode=false). Set to 0 to disable truncation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bash.Output(request) -> *sandboxsdkgo.ResponseBashOutputResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read output from a bash session using offset-based streaming.

- offset/stderr_offset: byte offsets to read from (use values from previous response)
- wait=true: long-poll until new output arrives or wait_timeout
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
client.Bash.Output(
        context.TODO(),
        &sandboxsdkgo.BashOutputRequest{
            SessionId: "session_id",
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

**sessionId:** `string` — Target session ID
    
</dd>
</dl>

<dl>
<dd>

**commandId:** `*string` — Target a specific async command. If not set, uses session-level output.
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Stdout byte offset to read from
    
</dd>
</dl>

<dl>
<dd>

**stderrOffset:** `*int` — Stderr byte offset to read from
    
</dd>
</dl>

<dl>
<dd>

**wait:** `*bool` — If true, long-poll until new output is available or wait_timeout is reached.
    
</dd>
</dl>

<dl>
<dd>

**waitTimeout:** `*float64` — Max seconds to wait for new output when wait=true.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bash.Write(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Write input to a bash session's stdin.
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
client.Bash.Write(
        context.TODO(),
        &sandboxsdkgo.BashWriteRequest{
            SessionId: "session_id",
            Input: "input",
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

**sessionId:** `string` — Target session ID
    
</dd>
</dl>

<dl>
<dd>

**commandId:** `*string` — Target a specific async command. If not set, writes to current command.
    
</dd>
</dl>

<dl>
<dd>

**input:** `string` — Content to write to the process stdin
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bash.Kill(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send a signal to a bash session's process.
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
client.Bash.Kill(
        context.TODO(),
        &sandboxsdkgo.BashKillRequest{
            SessionId: "session_id",
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

**sessionId:** `string` — Target session ID
    
</dd>
</dl>

<dl>
<dd>

**signal:** `*string` — Signal to send: SIGTERM, SIGKILL, or SIGINT
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bash.Sessions() -> *sandboxsdkgo.ResponseListBashSessionInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all active bash sessions.
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
client.Bash.Sessions(
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

<details><summary><code>client.Bash.CreateSession(request) -> *sandboxsdkgo.ResponseBashSessionInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new bash session.
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
client.Bash.CreateSession(
        context.TODO(),
        &sandboxsdkgo.BashSessionCreateRequest{},
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

**sessionId:** `*string` — Session ID. Auto-generated if not provided.
    
</dd>
</dl>

<dl>
<dd>

**execDir:** `*string` — Working directory for the new session (absolute path)
    
</dd>
</dl>

<dl>
<dd>

**snapshotPath:** `*string` — Path to a shell snapshot script to source on session init
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Bash.CloseSession(SessionId) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Close a bash session.
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
client.Bash.CloseSession(
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

<details><summary><code>client.File.GrepFiles(request) -> *sandboxsdkgo.ResponseFileGrepResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Multi-file content search (grep) with regex or fixed string support
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
client.File.GrepFiles(
        context.TODO(),
        &sandboxsdkgo.FileGrepRequest{
            Path: "path",
            Pattern: "pattern",
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

**path:** `string` — File or directory path to search
    
</dd>
</dl>

<dl>
<dd>

**pattern:** `string` — Search pattern (regex or fixed string)
    
</dd>
</dl>

<dl>
<dd>

**include:** `[]string` — File glob filters to include (e.g., ["*.py", "*.ts"])
    
</dd>
</dl>

<dl>
<dd>

**exclude:** `[]string` — Glob patterns to exclude (e.g., ["node_modules", "*.min.js"])
    
</dd>
</dl>

<dl>
<dd>

**caseInsensitive:** `*bool` — Case insensitive search
    
</dd>
</dl>

<dl>
<dd>

**fixedStrings:** `*bool` — Treat pattern as literal string, not regex
    
</dd>
</dl>

<dl>
<dd>

**contextBefore:** `*int` — Number of lines before each match (-B)
    
</dd>
</dl>

<dl>
<dd>

**contextAfter:** `*int` — Number of lines after each match (-A)
    
</dd>
</dl>

<dl>
<dd>

**maxResults:** `*int` — Maximum number of matches to return
    
</dd>
</dl>

<dl>
<dd>

**maxFileSize:** `*string` — Skip files larger than this size (e.g., 1M, 500K)
    
</dd>
</dl>

<dl>
<dd>

**multiline:** `*bool` — Enable multiline matching where . matches newlines and patterns can span lines (rg -U --multiline-dotall)
    
</dd>
</dl>

<dl>
<dd>

**offset:** `*int` — Skip first N matches before returning results (for pagination)
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — File type filter using ripgrep type aliases (e.g., "py", "js", "rust", "go"). Maps to rg --type.
    
</dd>
</dl>

<dl>
<dd>

**recursive:** `*bool` — Search recursively
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.GlobFiles(request) -> *sandboxsdkgo.ResponseFileGlobResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enhanced file glob matching with optional metadata
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
client.File.GlobFiles(
        context.TODO(),
        &sandboxsdkgo.FileGlobRequest{
            Path: "path",
            Pattern: "pattern",
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

**path:** `string` — Base directory path
    
</dd>
</dl>

<dl>
<dd>

**pattern:** `string` — Glob pattern (**, *, ?, [...])
    
</dd>
</dl>

<dl>
<dd>

**exclude:** `[]string` — Glob patterns to exclude
    
</dd>
</dl>

<dl>
<dd>

**includeHidden:** `*bool` — Whether to include hidden files
    
</dd>
</dl>

<dl>
<dd>

**filesOnly:** `*bool` — Only return files (not directories)
    
</dd>
</dl>

<dl>
<dd>

**includeMetadata:** `*bool` — Whether to include size and modified time
    
</dd>
</dl>

<dl>
<dd>

**maxResults:** `*int` — Maximum number of results
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*string` — Sort by: path, name, size, modified
    
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

<details><summary><code>client.File.DownloadFile() -> string</code></summary>
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

<dl>
<dd>

**replaceMode:** `*sandboxsdkgo.StrReplaceEditorRequestReplaceMode` — Optional parameter of `str_replace` command. When specified, controls how multiple occurrences are handled: 'ALL' replaces all occurrences, 'FIRST' replaces only the first, 'LAST' replaces only the last. If not specified, requires unique match (original behavior).
    
</dd>
</dl>

<dl>
<dd>

**pageRange:** `[]int` — Optional parameter for `view` command on PDF files. Specifies page range [start, end] (1-indexed). E.g., [1, 5] reads pages 1-5.
    
</dd>
</dl>

<dl>
<dd>

**sheetName:** `*string` — Optional parameter for `view` command on Excel files. Specifies which sheet to read. If not provided, all sheets are returned.
    
</dd>
</dl>

<dl>
<dd>

**rowRange:** `[]int` — Optional parameter for `view` command on Excel files. Specifies row range [start, end] (1-indexed). E.g., [1, 100] reads rows 1-100.
    
</dd>
</dl>

<dl>
<dd>

**slideRange:** `[]int` — Optional parameter for `view` command on PPTX files. Specifies slide range [start, end] (1-indexed). E.g., [1, 5] reads slides 1-5.
    
</dd>
</dl>

<dl>
<dd>

**enableMetadata:** `*bool` — Optional parameter for `view` command. If true, returns file metadata (total pages, sheets, slides, etc.) in the response.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.WatchList() -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WatchList(
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

<details><summary><code>client.File.WatchCreate(request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WatchCreate(
        context.TODO(),
        &sandboxsdkgo.CreateWatchRequest{
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

**path:** `string` — 监听目录或文件路径
    
</dd>
</dl>

<dl>
<dd>

**recursive:** `*bool` — 是否递归子目录
    
</dd>
</dl>

<dl>
<dd>

**exclude:** `[]string` — 排除的目录/glob 模式
    
</dd>
</dl>

<dl>
<dd>

**debounce:** `*int` — 去抖动窗口(ms)
    
</dd>
</dl>

<dl>
<dd>

**includePatterns:** `[]string` — glob 过滤，空=全部通过
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.WatchEvents(WatcherId) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WatchEvents(
        context.TODO(),
        "watcher_id",
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

**watcherId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.WatchPoll(WatcherId, request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WatchPoll(
        context.TODO(),
        "watcher_id",
        &sandboxsdkgo.PollRequest{},
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

**watcherId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*int` — 上次返回的游标值，只返回 seq > cursor 的事件
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — 最多返回条数
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*int` — 长轮询等待秒数，0=立即返回
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.WatchWait(request) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WatchWait(
        context.TODO(),
        &sandboxsdkgo.FileWatchWaitRequest{
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

**path:** `string` — 等待的文件路径（精确匹配）
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*int` — 最大等待秒数
    
</dd>
</dl>

<dl>
<dd>

**eventTypes:** `[]*sandboxsdkgo.AppSchemasFileWatchWaitRequestEventTypesItem` — 关注的事件类型
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.File.WatchStop(WatcherId) -> any</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.File.WatchStop(
        context.TODO(),
        "watcher_id",
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

**watcherId:** `string` 
    
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

**kernelName:** `*string` — Kernel name: 'python3', 'python3.10', 'python3.11', 'python3.12'. Defaults to 'python3'.
    
</dd>
</dl>

<dl>
<dd>

**sessionId:** `*string` — Session ID to maintain kernel state across requests
    
</dd>
</dl>

<dl>
<dd>

**cwd:** `*string` — Current working directory for the kernel
    
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

<details><summary><code>client.Jupyter.CreateSession(request) -> *sandboxsdkgo.ResponseJupyterCreateSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Jupyter session
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
client.Jupyter.CreateSession(
        context.TODO(),
        &sandboxsdkgo.JupyterCreateSessionRequest{},
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

**sessionId:** `*string` — Unique identifier for the session, auto-generated if not provided
    
</dd>
</dl>

<dl>
<dd>

**kernelName:** `*string` — Kernel name: 'python3', 'python3.10', 'python3.11', 'python3.12'. Defaults to 'python3'.
    
</dd>
</dl>

<dl>
<dd>

**cwd:** `*string` — Current working directory for the session
    
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

For stateless execution (default):
- Each request creates a fresh execution environment
- Environment is cleaned up automatically after execution

For stateful execution (stateful=True):
- Uses persistent REPL session that maintains state between requests
- Variables, functions, and imports persist across calls
- Returns session_id to continue the session in subsequent requests
- Supports async/await at top level
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

<dl>
<dd>

**stateful:** `*bool` — Enable stateful execution with persistent REPL session
    
</dd>
</dl>

<dl>
<dd>

**sessionId:** `*string` — Session ID for stateful execution (reuse existing session)
    
</dd>
</dl>

<dl>
<dd>

**cwd:** `*string` — Working directory for code execution
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — Node.js version to use: "node20", "node22", "node24", or aliases "20", "22", "24"
    
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

Get information about Node.js REPL runtime, including installed packages

Returns Node.js version, npm version, and lists of installed packages
from both the runtime directory and global npm directory.
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

<details><summary><code>client.Nodejs.ListSessions() -> *sandboxsdkgo.ResponseNodeJsSessionListResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all active Node.js REPL sessions

Returns information about all active sessions including their state,
working directory, and idle time.
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
client.Nodejs.ListSessions(
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

<details><summary><code>client.Nodejs.CreateSession(request) -> *sandboxsdkgo.ResponseNodeJsCreateSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Node.js REPL session

Creates a new persistent REPL session with configurable working directory
and idle timeout. Use the returned session_id in subsequent execute requests.
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
client.Nodejs.CreateSession(
        context.TODO(),
        &sandboxsdkgo.NodeJsCreateSessionRequest{},
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

**sessionId:** `*string` — Custom session ID (auto-generated if not provided)
    
</dd>
</dl>

<dl>
<dd>

**cwd:** `*string` — Working directory for the session
    
</dd>
</dl>

<dl>
<dd>

**maxIdleTime:** `*int` — Maximum idle time in seconds (default 24 hours)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Nodejs.GetSession(SessionId) -> *sandboxsdkgo.ResponseNodeJsSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get information about a specific Node.js REPL session

Returns detailed information about a session including its state,
working directory, creation time, and idle time.
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
client.Nodejs.GetSession(
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

<details><summary><code>client.Nodejs.DeleteSession(SessionId) -> *sandboxsdkgo.ResponseNodeJsDeleteSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a Node.js REPL session

Terminates the session and releases all associated resources.
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
client.Nodejs.DeleteSession(
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

<details><summary><code>client.Nodejs.UpdateSession(SessionId, request) -> *sandboxsdkgo.ResponseNodeJsUpdateSessionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a Node.js REPL session configuration

Updates session properties like maximum idle time or working directory.
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
client.Nodejs.UpdateSession(
        context.TODO(),
        "session_id",
        &sandboxsdkgo.NodeJsUpdateSessionRequest{},
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

<dl>
<dd>

**maxIdleTime:** `*int` — New maximum idle time in seconds
    
</dd>
</dl>

<dl>
<dd>

**cwd:** `*string` — New working directory
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Mcp
<details><summary><code>client.Mcp.ListMcpTools(ServerName) -> *sandboxsdkgo.ResponseListToolsResultModel</code></summary>
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

<details><summary><code>client.Mcp.ExecuteMcpTool(ServerName, ToolName, request) -> *sandboxsdkgo.ResponseCallToolResultModel</code></summary>
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

<details><summary><code>client.Browser.SetConfig(request) -> *sandboxsdkgo.Response</code></summary>
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
client.Browser.SetConfig(
        context.TODO(),
        &sandboxsdkgo.BrowserConfigRequest{},
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

**resolution:** `*sandboxsdkgo.Resolution` — The desired screen resolution, allowed values are: 1920x1080, 640x480, 1360x768, 1280x720, 800x600, 1024x768, 1280x800, 1920x1200, 1280x960, 1400x1050, 1680x1050, 1280x1024, 1600x1200.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Browser.Restart(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Restart the browser session.

- **soft** (default): Reconnect the Playwright session only.
- **hard**: Restart the browser process via supervisorctl, optionally
  updating URL blocklist/allowlist policies before restart.
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
client.Browser.Restart(
        context.TODO(),
        &sandboxsdkgo.RestartRequest{},
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

**request:** `*sandboxsdkgo.RestartRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Browser.GetProxyPac() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the current PAC file for proxy split-routing (if configured).
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
client.Browser.GetProxyPac(
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

## BrowserPage
<details><summary><code>client.BrowserPage.Navigate(request) -> *sandboxsdkgo.ResponseDict</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Navigate the current page to a URL.
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
client.BrowserPage.Navigate(
        context.TODO(),
        &sandboxsdkgo.NavigateRequest{
            Url: "url",
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

**url:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**waitUntil:** `*sandboxsdkgo.NavigateRequestWaitUntil` 
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Back() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Go back in browser history.
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
client.BrowserPage.Back(
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

<details><summary><code>client.BrowserPage.Forward() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Go forward in browser history.
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
client.BrowserPage.Forward(
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

<details><summary><code>client.BrowserPage.Reload() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reload the current page.
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
client.BrowserPage.Reload(
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

<details><summary><code>client.BrowserPage.Click(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Click an element by selector, index, or coordinates.
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
client.BrowserPage.Click(
        context.TODO(),
        &sandboxsdkgo.ClickRequest{},
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

**selector:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**index:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**x:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**y:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**button:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clickCount:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Fill(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Fill an input field.
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
client.BrowserPage.Fill(
        context.TODO(),
        &sandboxsdkgo.FillRequest{
            Text: "text",
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

**selector:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**index:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**text:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.TypeText(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Type text with optional delay between keystrokes.
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
client.BrowserPage.TypeText(
        context.TODO(),
        &sandboxsdkgo.TypeTextRequest{
            Text: "text",
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

**text:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**delay:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.PressKey(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Press a single key.
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
client.BrowserPage.PressKey(
        context.TODO(),
        &sandboxsdkgo.KeyRequest{
            Key: "key",
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

**key:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.HotKey(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Press a key combination.
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
client.BrowserPage.HotKey(
        context.TODO(),
        &sandboxsdkgo.HotKeyRequest{
            Keys: []string{
                "keys",
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

**keys:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Hover(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Hover over an element by selector or coordinates.
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
client.BrowserPage.Hover(
        context.TODO(),
        &sandboxsdkgo.HoverRequest{},
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

**selector:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**x:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**y:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.SelectOption(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Select an option in a dropdown.
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
client.BrowserPage.SelectOption(
        context.TODO(),
        &sandboxsdkgo.SelectOptionRequest{
            Selector: "selector",
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

**selector:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**value:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**label:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**index:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Check(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check a checkbox.
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
client.BrowserPage.Check(
        context.TODO(),
        &sandboxsdkgo.CheckRequest{
            Selector: "selector",
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

**request:** `*sandboxsdkgo.CheckRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Uncheck(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Uncheck a checkbox.
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
client.BrowserPage.Uncheck(
        context.TODO(),
        &sandboxsdkgo.CheckRequest{
            Selector: "selector",
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

**request:** `*sandboxsdkgo.CheckRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.UploadFile(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload files to a file input element.
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
client.BrowserPage.UploadFile(
        context.TODO(),
        &sandboxsdkgo.UploadFileRequest{
            Selector: "selector",
            Files: []string{
                "files",
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

**selector:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**files:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.FillForm(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Batch fill multiple form fields.
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
client.BrowserPage.FillForm(
        context.TODO(),
        &sandboxsdkgo.FormFillRequest{
            Items: []map[string]any{
                map[string]any{
                    "key": "value",
                },
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

**items:** `[]map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Scroll(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Scroll the page in a direction.
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
client.BrowserPage.Scroll(
        context.TODO(),
        &sandboxsdkgo.ScrollRequest{},
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

**direction:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.ScrollTo(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Scroll to an absolute position.
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
client.BrowserPage.ScrollTo(
        context.TODO(),
        &sandboxsdkgo.ScrollToRequest{},
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

**x:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**y:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.ScrollToElement(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Scroll an element into view.
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
client.BrowserPage.ScrollToElement(
        context.TODO(),
        &sandboxsdkgo.ScrollToElementRequest{
            Selector: "selector",
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

**selector:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Screenshot() -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Capture a page screenshot.
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
client.BrowserPage.Screenshot(
        context.TODO(),
        &sandboxsdkgo.BrowserPageScreenshotRequest{},
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

**fullPage:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**format:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**quality:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Record(request) -> *sandboxsdkgo.ResponseDict</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record page screencast (once/start/pause/resume/stop/status).
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
client.BrowserPage.Record(
        context.TODO(),
        &sandboxsdkgo.RecordRequest{},
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

**action:** `*sandboxsdkgo.RecordRequestAction` 
    
</dd>
</dl>

<dl>
<dd>

**savePath:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**duration:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**fps:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**quality:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.GetHtml() -> *sandboxsdkgo.ResponseStr</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the page HTML content.
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
client.BrowserPage.GetHtml(
        context.TODO(),
        &sandboxsdkgo.BrowserPageGetHtmlRequest{
            Outer: sandboxsdkgo.Bool(
                true,
            ),
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

**outer:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.GetText() -> *sandboxsdkgo.ResponseStr</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all visible text from the page.
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
client.BrowserPage.GetText(
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

<details><summary><code>client.BrowserPage.GetMarkdown() -> *sandboxsdkgo.ResponseDict</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the page content as Markdown using Readability and Turndown.
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
client.BrowserPage.GetMarkdown(
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

<details><summary><code>client.BrowserPage.GetElements() -> *sandboxsdkgo.ResponseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all interactive elements on the page.
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
client.BrowserPage.GetElements(
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

<details><summary><code>client.BrowserPage.GetConsole() -> *sandboxsdkgo.ResponseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get captured browser console log messages.
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
client.BrowserPage.GetConsole(
        context.TODO(),
        &sandboxsdkgo.BrowserPageGetConsoleRequest{
            Clear: sandboxsdkgo.Bool(
                true,
            ),
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

**clear:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.ExportConsole(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export console logs to a JSON file.
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
client.BrowserPage.ExportConsole(
        context.TODO(),
        &sandboxsdkgo.ExportConsoleLogsRequest{
            SavePath: "save_path",
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

**savePath:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**clear:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Evaluate(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Execute JavaScript and return the result.
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
client.BrowserPage.Evaluate(
        context.TODO(),
        &sandboxsdkgo.EvaluateRequest{
            Expression: "expression",
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

**expression:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.FindText(request) -> *sandboxsdkgo.ResponseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Find text occurrences on the page.
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
client.BrowserPage.FindText(
        context.TODO(),
        &sandboxsdkgo.FindTextRequest{
            Keyword: "keyword",
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

**keyword:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserPage.Wait(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unified wait: selector, load, url, network_idle, download, function, response, request, or timeout.
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
client.BrowserPage.Wait(
        context.TODO(),
        &sandboxsdkgo.WaitRequest{
            Type: sandboxsdkgo.TypeSelector,
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

**type_:** `*sandboxsdkgo.Type` 
    
</dd>
</dl>

<dl>
<dd>

**selector:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**savePath:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**timeout:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**expression:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**polling:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**urlPattern:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BrowserTabs
<details><summary><code>client.BrowserTabs.List() -> *sandboxsdkgo.ResponseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all open browser tabs.
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
client.BrowserTabs.List(
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

<details><summary><code>client.BrowserTabs.Create(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new browser tab.
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
client.BrowserTabs.Create(
        context.TODO(),
        &sandboxsdkgo.CreatePageRequest{},
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

**url:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserTabs.Close(Index) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Close a browser tab by index.
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
client.BrowserTabs.Close(
        context.TODO(),
        1,
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

**index:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserTabs.Activate(Index) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Activate (bring to front) a browser tab by index.
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
client.BrowserTabs.Activate(
        context.TODO(),
        1,
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

**index:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BrowserCookies
<details><summary><code>client.BrowserCookies.GetCookies() -> *sandboxsdkgo.ResponseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get browser cookies, optionally filtered by URLs (comma-separated).
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
client.BrowserCookies.GetCookies(
        context.TODO(),
        &sandboxsdkgo.BrowserCookiesGetCookiesRequest{
            Urls: sandboxsdkgo.String(
                "urls",
            ),
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

**urls:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserCookies.SetCookies(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set browser cookies.
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
client.BrowserCookies.SetCookies(
        context.TODO(),
        &sandboxsdkgo.CookieSetRequest{
            Cookies: []map[string]any{
                map[string]any{
                    "key": "value",
                },
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

**cookies:** `[]map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserCookies.ClearCookies() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Clear all browser cookies.
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
client.BrowserCookies.ClearCookies(
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

## BrowserState
<details><summary><code>client.BrowserState.Save(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Save browser state (cookies, localStorage, etc.) to a file.
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
client.BrowserState.Save(
        context.TODO(),
        &sandboxsdkgo.StateSaveRequest{
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

<details><summary><code>client.BrowserState.Load(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Load browser state from a previously saved file.
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
client.BrowserState.Load(
        context.TODO(),
        &sandboxsdkgo.StateLoadRequest{
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

## BrowserNetwork
<details><summary><code>client.BrowserNetwork.SetHeaders(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set extra HTTP headers for all subsequent requests.
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
client.BrowserNetwork.SetHeaders(
        context.TODO(),
        &sandboxsdkgo.HeadersRequest{
            Headers: map[string]string{
                "key": "value",
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

**headers:** `map[string]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserNetwork.SetScopedHeaders(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set HTTP headers scoped to a specific origin.
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
client.BrowserNetwork.SetScopedHeaders(
        context.TODO(),
        &sandboxsdkgo.ScopedHeadersRequest{
            Origin: "origin",
            Headers: map[string]string{
                "key": "value",
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

**origin:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**headers:** `map[string]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserNetwork.AddRoute(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a network route to mock or abort matching requests.
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
client.BrowserNetwork.AddRoute(
        context.TODO(),
        &sandboxsdkgo.NetworkRouteRequest{
            UrlPattern: "url_pattern",
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

**urlPattern:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**response:** `*sandboxsdkgo.RouteResponseModel` 
    
</dd>
</dl>

<dl>
<dd>

**abort:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserNetwork.RemoveRoute(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a previously added network route.
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
client.BrowserNetwork.RemoveRoute(
        context.TODO(),
        &sandboxsdkgo.NetworkRouteRemoveRequest{
            UrlPattern: "url_pattern",
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

**urlPattern:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserNetwork.GetRequests() -> *sandboxsdkgo.ResponseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get tracked network requests.
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
client.BrowserNetwork.GetRequests(
        context.TODO(),
        &sandboxsdkgo.BrowserNetworkGetRequestsRequest{
            Filter: sandboxsdkgo.String(
                "filter",
            ),
            Limit: sandboxsdkgo.Int(
                1,
            ),
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

**filter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BrowserNetwork.ExportHar(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export tracked network requests as a HAR file.
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
client.BrowserNetwork.ExportHar(
        context.TODO(),
        &sandboxsdkgo.ExportHarRequest{
            SavePath: "save_path",
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

**savePath:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BrowserCaptcha
<details><summary><code>client.BrowserCaptcha.Detect() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Detect CAPTCHA on the current page.
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
client.BrowserCaptcha.Detect(
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

<details><summary><code>client.BrowserCaptcha.Wait(request) -> *sandboxsdkgo.ResponseCaptchaWaitResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Wait for CAPTCHA to be resolved.
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
client.BrowserCaptcha.Wait(
        context.TODO(),
        &sandboxsdkgo.CaptchaWaitRequest{},
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

**timeout:** `*float64` 
    
</dd>
</dl>

<dl>
<dd>

**pollInterval:** `*float64` 
    
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

<dl>
<dd>

**cwd:** `*string` — Current working directory for code execution
    
</dd>
</dl>

<dl>
<dd>

**stateful:** `*bool` — Enable stateful execution using Jupyter kernel. When True, variables and state persist across requests with the same session_id.
    
</dd>
</dl>

<dl>
<dd>

**sessionId:** `*string` — Session ID for stateful execution. Required when stateful=True to maintain state across requests. Auto-generated if not provided.
    
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

Note: Version info is cached at service level (first call only runs subprocess).
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

## Skills
<details><summary><code>client.Skills.RegisterSkills(request) -> *sandboxsdkgo.ResponseSkillRegistrationResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Skills.RegisterSkills(
        context.TODO(),
        strings.NewReader(
            "",
        ),
        &sandboxsdkgo.BodyRegisterSkills{},
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

<details><summary><code>client.Skills.ListMetadata() -> *sandboxsdkgo.ResponseSkillMetadataCollection</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Skills.ListMetadata(
        context.TODO(),
        &sandboxsdkgo.SkillsListMetadataRequest{
            Names: sandboxsdkgo.String(
                "names",
            ),
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

**names:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.ClearSkills() -> *sandboxsdkgo.ResponseDict</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Skills.ClearSkills(
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

<details><summary><code>client.Skills.DeleteSkill(Name) -> *sandboxsdkgo.ResponseSkillMetadata</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Skills.DeleteSkill(
        context.TODO(),
        "name",
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

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Skills.GetContent(Name) -> *sandboxsdkgo.ResponseSkillContentResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Skills.GetContent(
        context.TODO(),
        "name",
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

**name:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Proxy
<details><summary><code>client.Proxy.ListMappings() -> *sandboxsdkgo.ResponseListProxyMappingRoute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all proxy domain-to-port mappings.
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
client.Proxy.ListMappings(
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

<details><summary><code>client.Proxy.AddMapping(request) -> *sandboxsdkgo.ResponseProxyMappingRoute</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add or update a proxy mapping. Supports wildcards in host (*.example.com).
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
client.Proxy.AddMapping(
        context.TODO(),
        &sandboxsdkgo.ProxyMappingAddRequest{
            Source: "myapp.com",
            Target: ":3000",
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

**source:** `string` — Source pattern: [protocol://]host[:port][/path], supports wildcard *
    
</dd>
</dl>

<dl>
<dd>

**target:** `string` — Target address: [host:]port[/path]. Host defaults to 127.0.0.1
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxy.RemoveMapping(Source) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a proxy mapping by source pattern.
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
client.Proxy.RemoveMapping(
        context.TODO(),
        "source",
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

**source:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxy.ListExcludes() -> *sandboxsdkgo.ResponseListStr</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all proxy bypass/exclude patterns.
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
client.Proxy.ListExcludes(
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

<details><summary><code>client.Proxy.AddExclude(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add a bypass/exclude pattern. Matched domains connect directly.
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
client.Proxy.AddExclude(
        context.TODO(),
        &sandboxsdkgo.ProxyBypassRequest{
            Pattern: "*.internal.company.com",
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

**request:** `*sandboxsdkgo.ProxyBypassRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxy.RemoveExclude(request) -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a bypass/exclude pattern.
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
client.Proxy.RemoveExclude(
        context.TODO(),
        &sandboxsdkgo.ProxyBypassRequest{
            Pattern: "*.internal.company.com",
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

**request:** `*sandboxsdkgo.ProxyBypassRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxy.Diagnose() -> *sandboxsdkgo.ResponseProxyDiagnoseResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Diagnose how a URL would be routed through the proxy.
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
client.Proxy.Diagnose(
        context.TODO(),
        &sandboxsdkgo.ProxyDiagnoseRequest{
            Url: "url",
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

**url:** `string` — URL to diagnose routing for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxy.Health() -> *sandboxsdkgo.ResponseProxyHealthCheck</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check proxy subsystem health: GOST alive, nginx alive, config consistency.
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
client.Proxy.Health(
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

<details><summary><code>client.Proxy.GetUpstream() -> *sandboxsdkgo.ResponseUnionProxyUpstreamInfoNoneType</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the current upstream proxy configuration.
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
client.Proxy.GetUpstream(
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

<details><summary><code>client.Proxy.SetUpstream(request) -> *sandboxsdkgo.ResponseProxyUpstreamInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set or update the upstream proxy. Supports user:pass@host:port format.

Takes effect immediately — no browser restart needed.
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
client.Proxy.SetUpstream(
        context.TODO(),
        &sandboxsdkgo.ProxyUpstreamUpdateRequest{
            Server: "proxy.example.com:3128",
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

**server:** `string` — Upstream proxy server. Supports plain host:port or user:pass@host:port
    
</dd>
</dl>

<dl>
<dd>

**authCmd:** `*string` — Optional shell command to obtain proxy credentials. The command stdout should be "username:password". When set, the result is injected into the server URL, replacing any inline credentials.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Proxy.RemoveUpstream() -> *sandboxsdkgo.Response</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove upstream proxy (switch to direct mode).
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
client.Proxy.RemoveUpstream(
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

## Display
<details><summary><code>client.Display.Record(request) -> *sandboxsdkgo.ResponseDisplayRecordResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Control display screen recording (start/stop/status).

Records the entire X11 desktop using ffmpeg x11grab,
including browser UI, multiple tabs, popups, terminal, etc.
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
client.Display.Record(
        context.TODO(),
        &sandboxsdkgo.DisplayRecordRequest{
            Action: sandboxsdkgo.DisplayRecordRequestActionStart,
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

**action:** `*sandboxsdkgo.DisplayRecordRequestAction` — Recording action: start, stop, or status
    
</dd>
</dl>

<dl>
<dd>

**savePath:** `*string` — Output file path (default: /tmp/recordings/recording_{timestamp}.mp4)
    
</dd>
</dl>

<dl>
<dd>

**fps:** `*int` — Frames per second
    
</dd>
</dl>

<dl>
<dd>

**crf:** `*int` — H.264 CRF quality (0=lossless, 51=worst)
    
</dd>
</dl>

<dl>
<dd>

**maxDuration:** `*float64` — Max recording duration in seconds
    
</dd>
</dl>

<dl>
<dd>

**width:** `*int` — Video width in pixels (auto-detected from X11 if omitted)
    
</dd>
</dl>

<dl>
<dd>

**height:** `*int` — Video height in pixels (auto-detected from X11 if omitted)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Auth
<details><summary><code>client.Auth.CreateTicket() -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create and return a short-lived authentication ticket.

This is a non-idempotent action; each call creates a new, unique ticket.
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
client.Auth.CreateTicket(
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

<details><summary><code>client.Auth.Authenticate() -> map[string]string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Authenticate a request using ticket or JWT.

This endpoint receives authentication subrequests (e.g., from Nginx auth_request).
It validates the request based on either a ticket in the 'x-original-uri'
header or a JWT in the 'Authorization' header.
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
client.Auth.Authenticate(
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
