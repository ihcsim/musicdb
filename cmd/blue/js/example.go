// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/ihcsim/bluelens/design
// --host=localhost
// --out=$(GOPATH)/src/github.com/ihcsim/bluelens/cmd/blue
// --scheme=http
// --version=v1.1.0
//
// bluelens JavaScript Client Example
//
// The content of this file is auto-generated, DO NOT MODIFY

package js

import (
	"github.com/goadesign/goa"
)

// MountController mounts the JavaScript example controller under "/js".
// This is just an example, not the best way to do this. A better way would be to specify a file
// server using the Files DSL in the design.
// Use --noexample to prevent this file from being generated.
func MountController(service *goa.Service) {
	// Serve static files under js
	service.ServeFiles("/js/*filepath", "/home/isim/workspace/go/src/github.com/ihcsim/bluelens/cmd/blue/js")
	service.LogInfo("mount", "ctrl", "JS", "action", "ServeFiles", "route", "GET /js/*")
}