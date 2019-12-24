package controller

import (
	"github.com/cheng1li/sdewan-operator/pkg/controller/sdewan"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sdewan.Add)
}
