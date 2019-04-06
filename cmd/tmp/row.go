package config

type Row struct {
	Name     string
	Value    interface{}
	Init     func()
	Get      func() interface{}
	Put      func(interface{}) bool
	Validate func(*Row, interface{}) bool
	Usage    string
}
type Rows map[string]*Row
