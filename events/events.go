package events

type Event int

const (
	InspectionSuccess Event = iota
	InspectionError
	BuildError
	SonnarFailed
)

// String() é o método padrão Go para representação textual (como toString())
func (e Event) String() string {
	switch e {
	case InspectionSuccess:
		return "Inspection/Success"
	case InspectionError:
		return "Inspection/Error"
	case BuildError:
		return "Build/Error"
	case SonnarFailed:
		return "Sonnar/Failed"
	default:
		return "UNKNOWN"
	}
}

func (e Event) WithContext(contexts ...ContextValue) *Context {
	ctx := &Context{Event: e}
	return ctx.PutValue(contexts...)
}
