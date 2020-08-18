package observer

type Observer interface {
	Update(value string)
}

func CreateWatcher(name string) *Watcher {
	return &Watcher{Name: name}
}

type Watcher struct {
	Name string
}

func (c *Watcher) Update(value string) {
	println("An event occurred: ", value)
}
