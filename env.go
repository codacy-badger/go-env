//
// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
//
// MIT Licence. See http://opensource.org/licenses/MIT
//
// Created on 2018-01-27
//

package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	sysEnv *envReader
)

func init() {
	sysEnv = &envReader{&realEnv{}}
}

// Env is the datasource for bindings and lookup. It is an optional
// parameter to Bind(). By specifying a custom Env, it's possible
// to populate a struct from an alternative source.
//
// The demo program in examples/docopt implements a custom Env
// to populate a struct from docopt options via Bind().
type Env interface {
	// Lookup retrieves the value of the variable named by key.
	//
	// It follows the same semantics as os.LookupEnv(). If a variable
	// is unset, the boolean will be false. If a variable is set, the
	// boolean will be true, but the variable may still be an empty
	// string.
	Lookup(key string) (string, bool)
}

// Get returns the value for envvar "key".
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or an empty string.
func Get(key string, fallback ...string) string {
	return sysEnv.Get(key, fallback...)
}

// GetString is a synonym for Get.
func GetString(key string, fallback ...string) string {
	return sysEnv.GetString(key, fallback...)
}

// GetInt returns the value for envvar "key" as an int.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with strconv.ParseInt(). If strconv.ParseInt()
// fails, tries to parse the number with strconv.ParseFloat() and
// truncate it to an int.
func GetInt(key string, fallback ...int) int {
	return sysEnv.GetInt(key, fallback...)
}

// GetFloat returns the value for envvar "key" as a float.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.0.
//
// Values are parsed with strconv.ParseFloat().
func GetFloat(key string, fallback ...float64) float64 {
	return sysEnv.GetFloat(key, fallback...)
}

// GetDuration returns the value for envvar "key" as a time.Duration.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with time.ParseDuration().
func GetDuration(key string, fallback ...time.Duration) time.Duration {
	return sysEnv.GetDuration(key, fallback...)
}

// GetBool returns the value for envvar "key" as a boolean.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or false.
//
// Values are parsed with strconv.ParseBool().
func GetBool(key string, fallback ...bool) bool {
	return sysEnv.GetBool(key, fallback...)
}

// realEnv reads values from the real environment
type realEnv struct{}

func (env *realEnv) Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

// envReader implements the conversion of strings to other types.
type envReader struct {
	env Env
}

func (r *envReader) Get(key string, fallback ...string) string {

	var fb string

	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}
	return s
}

func (r *envReader) GetString(key string, fallback ...string) string {
	return r.Get(key, fallback...)
}

func (r *envReader) GetInt(key string, fallback ...int) int {

	var fb int

	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	// log.Printf("[env] %s=%s", key, s)

	i, err := parseInt(s)
	if err != nil {
		return fb
	}

	return int(i)
}

func (r *envReader) GetFloat(key string, fallback ...float64) float64 {

	var fb float64

	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	// log.Printf("[env] %s=%s", key, s)

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fb
	}

	return n
}

func (r *envReader) GetDuration(key string, fallback ...time.Duration) time.Duration {

	var fb time.Duration

	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	// log.Printf("[env] %s=%s", key, s)

	d, err := time.ParseDuration(s)
	if err != nil {
		return fb
	}

	return d
}

func (r *envReader) GetBool(key string, fallback ...bool) bool {

	var fb bool

	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	// log.Printf("[env] %s=%s", key, s)

	b, err := strconv.ParseBool(s)
	if err != nil {
		return fb
	}

	return b
}

// parse an int, falling back to parsing it as a float
func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return int(i), nil
	}

	// Try to parse as float, then convert
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int: %v", s)
	}
	return int(n), nil
}
