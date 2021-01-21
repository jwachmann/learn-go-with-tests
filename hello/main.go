package main

import (
	"os"
	"time"

	"james.com/hello2/mocking"
)

func main() {
	// fmt.Println(hello("World", ""))
	// fmt.Println(rscHello())

	sleeper := mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
