# Golang SDK for BMM

## Status

This SDK implements only the things needed by BCC Media at the moment.
It is primarily meant for use in the backend.

PRs expanding it are welcome.

## Usage

```go
import (
    "github.com/bccmedia/bmm-sdk-golang"
)

func main() {
    token, err:= bmm.Token(...)
    if err != nil {
        panic(err)
    }
    
    client := bmm.NewApiClient("http://bmm.base.url", token)
}
```

## Logging

Internally the SDK uses the [slog](https://pkg.go.dev/log/slog) package for logging.
A custom logger can be set using the `SetLogger(logger *slog.Logger)` method.
