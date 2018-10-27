# Env 
Env is a simple package providing configuration as environment variables for 12 factor applications.

## Usage
```
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
