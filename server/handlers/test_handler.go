package handlers

import "ComputerWorld_API/server"

// TODO: Learn about how handlers are implemented in the mdm-api on GitHub. Ask Cameron, or Matt on what these are and how I would go about implementing them.
// TODO: It seems to be similar to the repository files in terms of layout.

type TestHandler struct {
	Server server.Server
}

func NewTestHandler(server server.Server) *TestHandler {
	return &TestHandler{Server: server}
}
