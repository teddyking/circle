package helpers

import (
	"fmt"
	"io"
	"os/exec"
)

type Kubectl struct {
	Context Context

	stdout io.Writer
	stderr io.Writer
}

func NewKubectlForContext(context string, stdout, stderr io.Writer) (*Kubectl, error) {
	kubeconfig, err := LoadDefaultKubeconfig()
	if err != nil {
		return &Kubectl{}, err
	}

	var matchedContext *Context
	for _, ctx := range kubeconfig.Contexts {
		if ctx.Name == context {
			matchedContext = &ctx
		}
	}

	if matchedContext == nil {
		return &Kubectl{}, fmt.Errorf("unable to find context '%s' in kubeconfig", context)
	}

	return &Kubectl{
		Context: *matchedContext,
		stdout:  stdout,
		stderr:  stderr,
	}, nil
}

func (k *Kubectl) Apply(filepath string) error {
	cmd := exec.Command("kubectl", "--context", k.Context.Name, "apply", "-f", filepath)
	return cmd.Run()
}

func (k *Kubectl) Get(resource, resourceName string) error {
	cmd := exec.Command("kubectl", "--context", k.Context.Name, "get", resource, resourceName)
	cmd.Stdout = k.stdout
	cmd.Stderr = k.stderr
	return cmd.Run()
}

func (k *Kubectl) Delete(filepath string) error {
	cmd := exec.Command("kubectl", "--context", k.Context.Name, "delete", "-f", filepath)
	return cmd.Run()
}
