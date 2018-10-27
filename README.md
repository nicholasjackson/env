# Env 
Env is a simple package providing configuration as environment variables for 12 factor applications, inspired by the Go `flag` package.

## Basic usage
```go
package main

var bindAddress = env.String("BIND_ADDRESS",true,"","bind address for server, i.e. localhost")
var bindPort = env.Integer("BIND_Port",true,0,"bind port for server, i.e. 9090")

func main() {
  err := env.Parse
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  fmt.Println("bind address:", *bindAddress)
  fmt.Println("bind port:", *bindPort)
}
```

## Showing help menu
Configuring an application with environment variables can be confusing to the user as often the configuration options are listed in the documentation rather than accessible from the command line like flag based configuration.  Env has a built in help menu which can be used like the following example:
```go

var version = "v0.0.0"

var bindAddress = env.String("BIND_ADDR", false, "localhost:9090", "Bind address for the server, i.e. localhost:9090")
var cacheURI = env.String("CACHE_URI", true, "", "URI for the cache server, i.e. localhost:9090")

var help = flag.Bool("help", false, "--help to show help")

func main() {
	//logger := hclog.Default()
	flag.Parse()

	// if the help flag is passed show configuration options
	if *help == true {
		fmt.Println("My service version:", version)
		fmt.Println("Configuration values are set using environment variables, for info please see the following list")
		fmt.Println("")
		fmt.Println(env.Help())
	}

```

```bash
$ service --help
My service version: v0.0.0
Configuration values are set using environment variables, for info please see the following list

Environment variables:
  BIND_ADDR  default: 'localhost:9090'
       Bind address for the server, i.e. localhost:9090
  CACHE_URI  default: no default
       URI for the cache server, i.e. localhost:9090
```
