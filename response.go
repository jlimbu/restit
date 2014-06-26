package restit

// Response needed to fulfill this interface
// in order to be tested by RESTit
type Response interface {
	Count() int
	NthValid(int) error
	GetNth(int) (interface{}, error)
	Match(interface{}, interface{}) error
}