package restgopher

import (
	"testing"

	"github.com/agoradevjsj/RESTGopher/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func TestNewServerGopher(t *testing.T) {
	server := NewServerGopher(CreateConfigApp(":4545", "localhost", "Hello World!!"))

	assert.NotEmpty(t, server.PortServer, "It should have :4545")
	assert.NotEmpty(t, server.message, "It should have being different to empty.")
	assert.NotEmpty(t, server.DomainServer, "It should have being different to empty.")

	server.Use(middleware.JSONMiddleware)

	assert.NotNil(t, server, "The server should be Not Nil.")
	go server.StartServerGopher()
}

func TestNewServerGopherWhitOutParams(t *testing.T) {
	server := NewServerGopher()

	assert.NotEqual(t, ":8000", server.PortServer, "It should have :8000 and have the port "+server.PortServer)
	assert.Equal(t, "", server.message, "It should have being empty and it has a message: "+server.message)
	assert.Equal(t, "", server.DomainServer, "It should have being empty and it has a domanin: "+server.DomainServer)

	server.Use(middleware.JSONMiddleware)

	assert.NotNil(t, server, "The server should be Not Nil.")
	go server.StartServerGopher()
}
