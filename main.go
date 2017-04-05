package main

import (
	"flag"
	"log"
	"os/user"
	"strconv"

	"github.com/docker/go-plugins-helpers/authorization"
)

const (
	defaultDockerHost = "unix:///var/run/docker.sock"
	// I must call the socket like this since socket name with length longer than that will fail on Mac
	pluginSocket = "/dap.sock"
)

var (
	flDockerHost = flag.String("host", defaultDockerHost, "Specifies the host where docker daemon is running")
	// Version is version
	Version string
	// Build is build
	Build string
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Plugin Version:", Version, "Build: ", Build)

	// Create image authorization plugin
	plugin, err := newPlugin(*flDockerHost)
	checkError(err)

	handler := authorization.NewHandler(plugin)

	// var err2 = handler.ServeTCP("docker-auth-test", ":8787", nil)
	// checkError(err2)

	// Start service handler on the local sock
	rootUser, err := user.Lookup("root")
	checkError(err)
	gid, err := strconv.Atoi(rootUser.Gid)
	checkError(err)
	if err := handler.ServeUnix(pluginSocket, gid); err != nil {
		log.Fatal(err)
	}
}
