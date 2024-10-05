package raketa

import "time"

type DeliveryTrace struct {
	Step []Step
}

type Step struct {
	Status string
	Date   time.Time
}
