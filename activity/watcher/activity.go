package watcher

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
        "github.com/TIBCOSoftware/flogo-lib/logger"
	"time"
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

	name := context.GetInput("name").(string)
        logger.Debugf("Hi [%s]", name)
        cTime := time.Now()  
        // logger.Debugf("Now is: [%s]", cTime)
        onTime := time.Date(cTime.Year(), cTime.Month(), cTime.Day(), 10, 30, 00, 000000000, time.UTC)
        // logger.Debugf("Now is: [%s]", onTime) 	
        late := cTime.After(onTime)
        context.SetOutput("result", late)
	return true, nil
}
