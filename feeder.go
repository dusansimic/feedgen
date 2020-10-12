package feedgen

import "time"

type Feeder interface {
	Title() string
	Subtitle() string
	Icon() string
	Links() []Link
	ID() string
	Updated() time.Time
	Entries() []Contenter
}
