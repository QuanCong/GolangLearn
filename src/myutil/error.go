package myutil

import "log"

//use Log.Fatal to log the error
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//panic the error
func PanicError(err error) {
	if err != nil {
		panic(err)
	}

}