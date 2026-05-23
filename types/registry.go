package types

import "errors"

var (
	registry       = map[uint16]*Model{}
	ErrNoSuchModel = errors.New("no such model")
)

// RegisterModel adds a model definition to the global registry, keyed
// on its numeric id. Generated model packages call this from their
// init() so that registering is automatic when the package is imported.
func RegisterModel(m *Model) {
	registry[m.Id] = m
}

// GetModel returns the model definition with the given id, or nil if
// none is registered.
func GetModel(id uint16) *Model {
	return registry[id]
}

// DoModels calls fn once for every registered model. Iteration stops
// and the error is returned if fn returns a non-nil error.
func DoModels(fn func(*Model) error) error {
	for _, m := range registry {
		if err := fn(m); err != nil {
			return err
		}
	}
	return nil
}
