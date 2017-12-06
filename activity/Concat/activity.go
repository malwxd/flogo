package Concat

import (
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	First := context.GetInput("First").(string)
	Second := context.GetInput("Second").(string)

	// Use the log object to log the greeting
	logger.Debugf("The Flogo engine will concat following strings: [%s] and  [%s]", First, Second)

	// Set the result as part of the context
	context.SetOutput("result", First+" "+Second)


	return true, nil
}
