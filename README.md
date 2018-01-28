
go-env
======

Easily access environment variables from Go, and bind them to structs.


<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
    * [Access environment variables](#access-environment-variables)
    * [Binding](#binding)
* [Installation](#installation)
* [Documentation](#documentation)
* [Licence](#licence)

<!-- vim-markdown-toc -->

Usage
-----

Import path is `git.deanishe/deanishe/go-env`, import name is `env`.

You can directly access environment variables, or populate your structs from them using struct tags and `env.Bind()`.


### Access environment variables

Read `int`, `float64`, `duration` and `string` values from environment variables, with optional fallback values for unset variables.

```go
import "git.deanishe.net/deanishe/go-env"

// Get value for key or return empty string
s := env.Get("SHELL")

// Get value for key or return a specified default
s := env.Get("DOES_NOT_EXIST", "fallback value")

// Get an int (or 0 if SOME_NUMBER is unset or empty)
i := env.GetInt("SOME_NUMBER")

// Int with a fallback
i := env.GetInt("SOME_UNSET_NUMBER", 10)
```


### Binding

You can also populate a struct directly from the environment by appropriately tagging it and calling `env.Bind()`:

```go
// Simple configuration struct
type config struct {
    HostName string `env:"HOSTNAME"` // default would be HOST_NAME
    Port     int    // leave as default (PORT)
    SSL      bool   `env:"USE_SSL"` // default would be SSL
    Online   bool   `env:"-"`       // ignore this field
}

// Set some values in the environment for test purposes
os.Setenv("HOSTNAME", "api.example.com")
os.Setenv("PORT", "443")
os.Setenv("USE_SSL", "1")
os.Setenv("ONLINE", "1") // will be ignored

// Create a config and bind it to the environment
c := &config{}
if err := Bind(c); err != nil {
    // handle error...
}

// config struct now populated from the environment
fmt.Println(c.HostName)
fmt.Printf("%d\n", c.Port)
fmt.Printf("%v\n", c.SSL)
fmt.Printf("%v\n", c.Online)
// Output:
// api.example.com
// 443
// true
// false
```


Installation
------------

```bash
go get git.deanishe.net/deanishe/go-env
```


Documentation
-------------

Read the documentation on [GoDoc][godoc].


Licence
-------

This library is released under the [MIT Licence][mit].

[mit]: ./LICENCE.txt
[godoc]: http://localhost:9090/pkg/git.deanishe.net/deanishe/go-env/
