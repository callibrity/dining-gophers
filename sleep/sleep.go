package sleep

import (
	"math/rand"
	"time"
)

func RandomSleep() {
	sleepTime := rand.Intn(250) + 250
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
}
