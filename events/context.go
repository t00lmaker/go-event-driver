package events

type ContextValue interface {
	ContextKey() string
	ContextValue() map[string]string
}

type Context struct {
	Event  Event
	Params []ContextValue
}

func (c *Context) NewContext(event Event) *Context {
	return &Context{
		Event:  event,
		Params: make([]ContextValue, 0),
	}
}

func (c *Context) GetByKey(key string) ContextValue {
	for _, value := range c.Params {
		if value.ContextKey() == key {
			return value
		}
	}
	return nil
}

func (c *Context) PutValue(contextValue ...ContextValue) *Context {
	if c.Params == nil {
		c.Params = make([]ContextValue, 0)
	}
	c.Params = append(c.Params, contextValue...)

	return c
}
