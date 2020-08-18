package observer

import "time"

type Observable struct {
	Observers []Observer
	Notifies  []time.Time
}

func (o *Observable) AddObserver(os Observer) {
	o.Observers = append(o.Observers, os)
}

func (o *Observable) HasObservers() bool {
	return len(o.Observers) > 0
}

func (o *Observable) NotifyAll(value string) {
	o.Notifies = append(o.Notifies, time.Now())

	for _, ob := range o.Observers {
		ob.Update(value)
	}
}
