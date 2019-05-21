package env

import (
    "flag"
)

var host string;
var port int;

func init() {
    flag.StringVar(&host, "host", "localhost", "specify the service listening host")
    flag.IntVar(&port, "port", 8080, "specify the service listening port")
}

func GetCmdlineHostAndPort() ( string, int) {
    flag.Parse()

    return host, port
}