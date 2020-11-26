package embedutil

// Footer ...
type Footer struct {
	text    string
	iconURL string
}

// NewFooter ...
func NewFooter() *Footer {
	return &Footer{}
}

// Text ...
func (f *Footer) Text(t string) *Footer {
	f.text = t
	return f
}

// IconURL ...
func (f *Footer) IconURL(u string) *Footer {
	f.iconURL = u
	return f
}

// Footer ...
func (e *Embed) Footer(f *Footer) *Embed {
	e.footer = f
	return e
}

// SetFooter ...
func (e *Embed) SetFooter(args ...string) *Embed {
	f := NewFooter()
	if len(args) >= 1 {
		f.Text(args[0])
	}
	if len(args) >= 2 {
		f.IconURL(args[1])
	}
	e.footer = f
	return e
}
