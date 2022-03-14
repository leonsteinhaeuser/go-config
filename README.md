# go-config

[![unit-tests](https://github.com/leonsteinhaeuser/go-config/actions/workflows/tests.yml/badge.svg)](https://github.com/leonsteinhaeuser/go-config/actions/workflows/tests.yml)

This repository provides a library that automatically detects the file extension and parses the data into the desired struct. If these steps were successful, the library will overload the parsed configuration with the desired environment variables.

## Example

```go
type Config struct {
    IsAlive bool
    Server struct {
        Address string
        Port int
    }
}

func main() {
    os.Setenv("CFG_ISALIVE", "true")
    os.Setenv("CFG_SERVER_ADDRESS", "0.0.0.0")
    os.Setenv("CFG_SERVER_PORT", "8491")

    cfg := Config{}
    AutoloadAndEnrichConfig("config.yml", &cfg)

    // custom env prefix
    os.Setenv("MYPREFIX_ISALIVE", "true")
    os.Setenv("MYPREFIX_SERVER_ADDRESS", "0.0.0.0")
    os.Setenv("MYPREFIX_SERVER_PORT", "8491")
    cfg2 := Config{}
    AutoloadAndEnrichConfigWithEnvPrefix("config.yml", "myprefix", &cfg2)
}
```
