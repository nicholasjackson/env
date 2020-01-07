# Env 

Env is a simple package providing configuration as environment variables for 12 factor applications, inspired by the Go `flag` package.

## Basic usage

```go
package main

var bindAddress = env.String("BIND_ADDRESS",true,"","bind address for server, i.e. localhost")
var bindPort = env.Integer("BIND_Port",true,0,"bind port for server, i.e. 9090")

func main() {
  err := env.Parse()
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  fmt.Println("bind address:", *bindAddress)
  fmt.Println("bind port:", *bindPort)
}
```

## Showing help menu

Configuring an application with environment variables can be confusing to the user as often the configuration options are listed in the documentation rather than accessible from the command line like flag based configuration.  Env has a built in help menu which can be accessed using the `--help` command line flag.

```go
var bindAddress = env.String("BIND_ADDR", false, "localhost:9090", "Bind address for the server, i.e. localhost:9090")
var cacheURI = env.String("CACHE_URI", true, "", "URI for the cache server, i.e. localhost:9090")

func main() {
  err := env.Parse()
```

```shell
➜ BIND_PORT=9090 BIND_ADDRESS=localhost go run main.go --help
Configuration values are set using environment variables, for info please see the following list.

Environment variables:
  BIND_ADDRESS  default: no default
       bind address for server, i.e. localhost
  BIND_PORT  default: '0'
       bind port for server, i.e. 9090
```