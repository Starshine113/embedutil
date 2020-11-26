package embedutil

// Field is a single embed field
type Field struct {
	name   string
	value  string
	inline bool
}

// NewField ...
func NewField() *Field {
	return &Field{}
}

// Name ...
func (f *Field) Name(n string) *Field {
	f.name = n
	return f
}

// Value ...
func (f *Field) Value(v string) *Field {
	f.value = v
	return f
}

// Inline toggles the field being inline
func (f *Field) Inline() *Field {
	if f.inline {
		f.inline = false
	} else {
		f.inline = true
	}
	return f
}

// Field adds a Field object
func (e *Embed) Field(f *Field) *Embed {
	e.fields = append(e.fields, f)
	return e
}

// AddField adds a field
func (e *Embed) AddField(name, value string) *Embed {
	e.fields = append(e.fields, &Field{
		name:  name,
		value: value,
	})
	return e
}

// Inline toggles the inline status of the last-added field
func (e *Embed) Inline() *Embed {
	if len(e.fields) == 0 {
		return e
	}
	if e.fields[len(e.fields)-1].inline {
		e.fields[len(e.fields)-1].inline = false
	} else {
		e.fields[len(e.fields)-1].inline = true
	}
	return e
}
