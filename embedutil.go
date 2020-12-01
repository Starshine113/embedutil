package embedutil

import (
	"strconv"
	"time"
)

// Embed limits
const (
	EmbedLimit       = 6000
	TitleLimit       = 256
	DescriptionLimit = 2048
	FieldValueLimit  = 1024
	FooterTextLimit  = 2048
	FieldLimit       = 25
)

// Embed holds the information for the embeds
type Embed struct {
	content     string
	title       string
	description string
	author      *Author
	fields      []*Field
	footer      *Footer

	url          string
	imageURL     string
	thumbnailURL string
	color        int
	timestamp    *time.Time
}

// NewEmbed ...
func NewEmbed() *Embed {
	return &Embed{}
}

// Description ...
func (e *Embed) Description(d string) *Embed {
	e.description = d
	return e
}

// Title ...
func (e *Embed) Title(t string) *Embed {
	e.title = t
	return e
}

// Content ...
func (e *Embed) Content(c string) *Embed {
	e.content = c
	return e
}

// URL ...
func (e *Embed) URL(u string) *Embed {
	e.url = u
	return e
}

// ImageURL ...
func (e *Embed) ImageURL(u string) *Embed {
	e.imageURL = u
	return e
}

// ThumbnailURL ...
func (e *Embed) ThumbnailURL(u string) *Embed {
	e.thumbnailURL = u
	return e
}

// Color ...
func (e *Embed) Color(c int) *Embed {
	e.color = c
	return e
}

// ColorString ...
func (e *Embed) ColorString(c string) *Embed {
	color, err := strconv.ParseInt(c, 0, 0)
	if err != nil {
		return e
	}
	e.color = int(color)
	return e
}

// Timestamp ...
func (e *Embed) Timestamp(t time.Time) *Embed {
	t = t.UTC()
	e.timestamp = &t
	return e
}
