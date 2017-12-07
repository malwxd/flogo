package withinbusinesshours

import (
	"time"	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
        "log"
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

	initialhour, ok := context.GetInput("initialhour").(int)
	if !ok {
		panic(ok)}
	initialminute, ok := context.GetInput("initialminute").(int)
	if !ok {
		panic(ok)}
	finalhour, ok := context.GetInput("finalhour").(int)
	if !ok {
		panic(ok)}
	finalminute, ok := context.GetInput("finalminute").(int)
	if !ok {
		panic(ok)}
	
	// do eval
        log.Println("Now the time is: ", time.Now().Local())
	//Get current local time        
	cTime := time.Now().Local()
	zone, ok := time.LoadLocation("Local")  

	if !ok {
		log.Println("Wrong zone")
		panic(ok)}

	

        startTime := time.Date(cTime.Year(), cTime.Month(), cTime.Day(), initialhour, initialminute, 00, 000000000, zone)
        endTime := time.Date(cTime.Year(), cTime.Month(), cTime.Day(), finalhour, finalminute, 00, 000000000, zone)

	log.Println("The business hours are between: ",startTime)
	log.Println("and ", endTime)

	timestr := cTime.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	//log.Println("Time: ", timestr)

	if cTime.After(startTime) && cTime.Before(endTime) {
		context.SetOutput("result", "true")
		log.Println("You came within business hours")	
	} else {
		context.SetOutput("result", "false")
		log.Println("You are trying to access out of business hours. Your attempt will be reported.")
	}
	context.SetOutput("currentTime", timestr)
	return true, nil
}
