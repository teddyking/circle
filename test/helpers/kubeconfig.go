package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Cluster struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

type User struct {
	Name string `yaml:"name"`
}

type Context struct {
	Name        string `yaml:"name"`
	ClusterName string `yaml:"cluster"`
	UserName    string `yaml:"user"`
}

type Kubeconfig struct {
	Contexts []Context `yaml:"contexts"`
	Users    []User    `yaml:"users"`
	Clusters []Cluster `yaml:"clusters"`
}

var DefaultKubeconfigFilepath = filepath.Join(os.Getenv("HOME"), ".kube", "config")

func LoadDefaultKubeconfig() (*Kubeconfig, error) {
	bytes, err := ioutil.ReadFile(DefaultKubeconfigFilepath)
	if err != nil {
		return &Kubeconfig{}, err
	}

	kubeconfig := &Kubeconfig{}
	if err := yaml.Unmarshal(bytes, kubeconfig); err != nil {
		return &Kubeconfig{}, err
	}

	return kubeconfig, nil
}
