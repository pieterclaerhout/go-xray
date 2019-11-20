## go-xray

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-xray)](https://goreportcard.com/report/github.com/pieterclaerhout/go-xray)
[![Documentation](https://godoc.org/github.com/pieterclaerhout/go-xray?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-xray)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-xray/raw/master/LICENSE)
[![GitHub version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-xray.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-xray)
[![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-xray.svg)](https://github.com/pieterclaerhout/go-xray/issues)

This is a [Golang](https://golang.org) library with reflection related functions which I use in my different projects.

## `KeyValue`

This type is used to construct a key-value pair. You can use it as follows:

```go
package main

import (
    "fmt"

    "https://github.com/pieterclaerhout/go-xray"
)

func main() {

    pair := &xray.KeyValue{
        Key:   "key",
        Value: "value",
    }

    fmt.Println(pair.String())
    // Ouptut: key=value

}
```

If both `Key` and `Value` are empty, an empty string is returned.

## `Name`

Name allows you to get the name of an object.

```go
package main

import (
    "fmt"

    "https://github.com/pieterclaerhout/go-xray"
)

func main() {

    v := &test{}

    fmt.Println(xray.Name(v))
    // Ouptut: test

    n := nil

    fmt.Println(xray.Name(n))
    // Ouptut: <nil>
    
}
```

## `Properties`

The function `xray.Properties` returns the a list with the property names defined in a struct:

```go
package main

import (
    "fmt"

    "https://github.com/pieterclaerhout/go-xray"
)

func main() {

    type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

    v, _ := xray.Properties(sampleStruct{})

    // v now contains:
    // []string{
    //     "Name",
    //     "Title",
    // }

}
```

The function `xray.PropertiesAsMap` does the same, but also includes the values and returns a `map[string]interface{}` instance:


```go
package main

import (
    "fmt"

    "https://github.com/pieterclaerhout/go-xray"
)

func main() {

    type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

    v, _ := xray.PropertiesAsMap(sampleStruct{})

    // v now contains:
    // map[string]interface{}{
    //     "Name": "name",
    //     "Title": "title",
    // }

}
```

The function `xray.Property` retrieves a single value by it's name:

```go
package main

import (
    "fmt"

    "https://github.com/pieterclaerhout/go-xray"
)

func main() {

    type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

    v, _ := xray.Property(sampleStruct{}, "Name")

    // v now contains:
    // "name"

}
```


## `Tags`

The `xray.Tags` function allows you to easily extract tags from a struct:

```go
package main

import (
    "fmt"

    "https://github.com/pieterclaerhout/go-xray"
)

func main() {

    type sampleStruct struct {
		Name  string `form:"name" json:"name"`
		Title string `form:"title" json:"title"`
	}

    v, _ := xray.Tags(sampleStruct{}, "form")

    // v now contains:
    // map[string]string{
    //     "Name": "name",
    //     "Title": "title",
    // }

}
```
