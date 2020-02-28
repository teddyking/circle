package repositories

import "fmt"

type ErrRuntimeNotFound struct {
	Name    string
	Runtime string
}

type ErrCloudFoundryRuntimeNotFound struct {
	Name string
}

func (e ErrRuntimeNotFound) Error() string {
	return fmt.Sprintf("%s runtime '%s' not found", e.Runtime, e.Name)
}

func (e ErrCloudFoundryRuntimeNotFound) Error() string {
	return ErrRuntimeNotFound{Name: e.Name, Runtime: "CloudFoundry"}.Error()
}

func IgnoreNotFound(err error) error {
	switch err.(type) {
	case ErrRuntimeNotFound:
		return nil
	case ErrCloudFoundryRuntimeNotFound:
		return nil
	}

	return err
}
