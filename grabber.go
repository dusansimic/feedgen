package feedgen

// Grabber is a type that implements procedures from grabbing contnent from sources.
type Grabber interface {
	Grab() Feeder
}
