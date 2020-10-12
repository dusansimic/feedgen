package feedgen

import "time"

// Contenter is a type that includes as single peice of contnet from a source. Information can be
// returned using getter methods.
type Contenter interface {
	Title() string
	Links() []string
	ID() string
	Updated() time.Time
	Summary() string
	Content() string
	Author() string
}
