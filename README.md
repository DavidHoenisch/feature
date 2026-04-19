# Feature

_simple, pure go feature flag library_


## Usage

Feature flags can be set in a variety of ways with optionally hierarical
evaluation. In its most basic form, flags are set using key=value pairs like any
other environmental variable.

```env
SOMEKEY=SOMEVALUE
```

However, `feature` provides methods for defining these in much great detail.



### The Syntax

Values are be broken down into:

1. Category: [_string_], defines an arbritrary high level feature category.
2. Feature: [_string_], defines an arbritrary feature title/name.
3. State: [_boolean_], defines the state of this feature by a boolan value.
   Any proper boolean supported by `strconv.ParseBool` is supported in this field.
4. Value: [_any_], defines an *optional* value that for the feature. This is useful
   for storing values that the application will use inside of the feature logic if the
   feature flag is enabled as defined by the `State` field.


When a variable is accessed in `feature` it is marshalled into the follow struct

```go
type Flag struct {
	Category string
	Feature  string
	State    bool
	Value    any
}
```

Flags values, then, can be defined in the following way:

```env
FEATURE_FLAG="ops:otel_export:true:0.1"
```

In this example, we have a category of 'ops', and a feature of 'otel_export'.
This feature is enabled and has value of '0.1'. We will use this '0.1' in our
feature logic to defined the sample rate we want our otel exporter to use.

## Basic Example

```go
package main

import (
  "fmt"
  "github.com/DavidHoenisch/feature"
)


func main() {

  flags, err := feature.NewFeature()
  if err != nil {
    // handle err
  }

  // get flag state (true/false) from env var 
  feature1 := flags.GetFlagState("SOME_FEATURE1_FLAG")
  if feature1 {
    fmt.Printf("SOME_FEATURE1_FLAG is set with %s", feature1)
  }

  // GetFlagModel is useful when you want to access the underlying
  // struct that the flag is marshalled into  
  feature2, val := flags.GetFlagModel("SOME_FEATURE2_FLAG")
  if feature2 {
    fmt.Printf("")
  }
  
}
  
```



```go

type basic interface {
  Foo()
}


type basicImpl struct {}


func (b *basicImpl) Bar()



  
```
