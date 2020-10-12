package instagram

import (
	"fmt"
	"time"
)

type C struct {
	Caption  string
	MediaID  string
	Medias   []string
	Username string
	Time     int
}

func (c C) Title() string {
	return fmt.Sprintf("%s posted a new photo", c.Username)
}

func (c C) Links() []string {
	return []string{
		fmt.Sprintf("https://instagram.com/p/%s/", c.MediaID),
	}
}

func (c C) ID() string {
	return c.MediaID
}

func (c C) Updated() time.Time {
	t := time.Unix(int64(c.Time), 0)
	return t
}

func (c C) Summary() string {
	return c.Caption
}

func (c C) Content() string {
	var imgs string
	if len(c.Medias) > 0 {
		for _, m := range c.Medias {
			imgs += fmt.Sprintf("<img src=\"https://instagram.com/p/%s/media/?size=l\"/>", m)
		}
	} else {
		imgs = fmt.Sprintf("<img src=\"https://instagram.com/p/%s/media/?size=l\"/>", c.MediaID)
	}
	return fmt.Sprintf("<p>%s</p>%s", c.Caption, imgs)
}

func (c C) Author() string {
	return c.Username
}
