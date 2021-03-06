package konfig

import "github.com/lalamove/nui/ngetter"

// Getter returns a mgetter.Getter for the key k
func Getter(k string) ngetter.GetterTyped {
	return instance().Getter(k)
}

func (c *store) Getter(k string) ngetter.GetterTyped {
	return ngetter.GetterTypedFunc(func() interface{} {
		return c.Get(k)
	})
}
