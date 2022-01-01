package types

import (
	"fmt"
	"io"
)

// Exec has the extra options for exec command
type Exec struct {
	Dir    string
	Env    Envs
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Pipe   bool
}

// Envs stores a mapping of keys and values to set multiple environments
type Envs struct {
	fields map[string]string
}

// KeyValuePair are used to initialize a new Envs
type KeyValuePair struct {
	Key   string
	Value string
}

// Env creates a new KeyValuePair for initializing Envs
func Env(key, value string) KeyValuePair {
	return KeyValuePair{Key: key, Value: value}
}

// NewEnvs returns a new Envs populated with initilial envs
func NewEnvs(initialEnvs ...KeyValuePair) Envs {
	envs := Envs{fields: map[string]string{}}
	for _, env := range initialEnvs {
		envs.Add(env.Key, env.Value)
	}
	return envs
}

// Add a new value to the set of values
func (envs Envs) Add(key, value string) {
	if len(envs.fields) == 0 {
		envs.fields = make(map[string]string)
	}
	envs.fields[key] = value
}

// ToString converts the mapping of key and value pair to string
func (envs Envs) ToString() []string {
	var strEnv []string
	for key, value := range envs.fields {
		strEnv = append(strEnv, fmt.Sprintf("%s=%s", key, value))
	}
	return strEnv
}
