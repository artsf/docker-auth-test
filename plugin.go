package main

import (
	"log"
	"net/url"
	"regexp"

	dockerapi "github.com/docker/docker/api"
	dockerclient "github.com/docker/docker/client"
	"github.com/docker/go-plugins-helpers/authorization"
)

// Image Authorization Plugin struct definition
type ImgAuthZPlugin struct {
	// Docker client
	client *dockerclient.Client
}

// Create a new image authorization plugin
func newPlugin(dockerHost string) (*ImgAuthZPlugin, error) {
	client, err := dockerclient.NewClient(dockerHost, dockerapi.DefaultVersion, nil, nil)

	if err != nil {
		return nil, err
	}

	return &ImgAuthZPlugin{
		client: client,
	}, nil
}

// AuthZReq Authorizes the docker client command.
// Non registry related commands are allowed by default.
// If the command uses a registry, the command is allowed only if the registry is authorized.
// Otherwise, the request is denied!
func (plugin *ImgAuthZPlugin) AuthZReq(req authorization.Request) authorization.Response {
	// Parse request and the request body
	reqURI, err := url.QueryUnescape(req.RequestURI)
	checkError(err)
	reqURL, err := url.ParseRequestURI(reqURI)
	checkError(err)

	urlPath := reqURL.Path
	log.Println(urlPath)

	// Catch container start in the format of /<version>/containers/{id}/start
	startURL, err := regexp.Compile("containers/\\w+/start$")
	checkError(err)
	if startURL.MatchString(urlPath) {
		return authorization.Response{Allow: false, Msg: "¡No pasarán!"}
	}

	// Allowed by default.
	return authorization.Response{Allow: true}
}

// AuthZRes Authorizes the docker client response.
// All responses are allowed by default.
func (plugin *ImgAuthZPlugin) AuthZRes(req authorization.Request) authorization.Response {
	// Allowed by default.
	return authorization.Response{Allow: true}
}
