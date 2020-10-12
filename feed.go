package feedgen

import "encoding/xml"

type Feed struct {
	XMLName  xml.Name `xml:"feed"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	// Icon     string   `xml:"icon,omitempty"`
	Image   Image   `xml:"image,omitempty"`
	Links   []Link  `xml:"link"`
	ID      string  `xml:"id"`
	Updated string  `xml:"updated"`
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Title   string  `xml:"title"`
	Links   []Link  `xml:"link"`
	ID      string  `xml:"id"`
	Updated string  `xml:"updated"`
	Summary string  `xml:"summary"`
	Content Content `xml:"content"`
	Author  Author  `xml:"author"`
}

type Content struct {
	Type string `xml:"type,attr"`
	Body string `xml:",chardata"`
}

type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email,omitempty"`
}

type Image struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Link  string `xml:"link"`
}
