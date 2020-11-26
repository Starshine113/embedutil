package embedutil

// Author ...
type Author struct {
	name      string
	avatarURL string
	url       string
}

// NewAuthor ...
func NewAuthor() *Author {
	return &Author{}
}

// Name ...
func (a *Author) Name(n string) *Author {
	a.name = n
	return a
}

// AvatarURL ...
func (a *Author) AvatarURL(u string) *Author {
	a.avatarURL = u
	return a
}

// URL ...
func (a *Author) URL(u string) *Author {
	a.url = u
	return a
}

// Author sets the author field to an Author object
func (e *Embed) Author(a *Author) *Embed {
	e.author = a
	return e
}

// SetAuthor sets the author fields automatically
func (e *Embed) SetAuthor(args ...string) *Embed {
	switch len(args) {
	case 1:
		e.author = &Author{
			name: args[0],
		}
	case 2:
		e.author = &Author{
			name:      args[0],
			avatarURL: args[1],
		}
	case 3:
		e.author = &Author{
			name:      args[0],
			avatarURL: args[1],
			url:       args[2],
		}
	}

	return e
}
