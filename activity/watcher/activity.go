package watcher

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
        "log"
	"time"
	"strconv"
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
        log.Println("Hi ", name)
        cTime := time.Now().Local()

	zone, ok := time.LoadLocation("Local")  
	if ok!=nil {
		log.Println("Wrong zone")
		panic(ok)} 
 
        onTime := time.Date(cTime.Year(), cTime.Month(), cTime.Day(), 10, 30, 00, 000000000, zone)
        late := cTime.After(onTime)
        str := strconv.FormatBool(late)
        context.SetOutput("result", str)
	return true, nil
}
