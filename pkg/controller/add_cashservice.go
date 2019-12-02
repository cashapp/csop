package controller

import (
	"github.com/cashapp/csop/pkg/controller/cashservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cashservice.Add)
}
