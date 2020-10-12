package informatikapmf

import "time"

type C struct {
	EntryTitle   string
	URL          string
	EntryContent string
	Category     string
	Time         string
}

func (c C) Title() string {
	return c.EntryTitle
}

func (c C) Links() []string {
	return []string{
		c.URL,
	}
}

func (c C) ID() string {
	return c.URL
}

func (c C) Updated() time.Time {
	t, _ := time.Parse(time.RFC3339, c.Time)
	return t
}

func (c C) Summary() string {
	return c.EntryContent
}

func (c C) Content() string {
	return c.EntryContent
}

func (c C) Author() string {
	return c.Category
}
