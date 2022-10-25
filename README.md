# Upsilon Serializer

Objective here is to store all "communcation" related structs, ways to serialize and unserialize them. This module is expected to be included in all upsilon servers. 

This may be overkill.

## Structure Usage

Let's take a struct for example:

```go
    type SampleStruct struct {
        ID uuid.UUID
        Name string
        SomeObscureData float64
    }
```

This is a data struct that will be transited over the network. To be used inside some application, there are three way to go:

Either by embedding
```go
    type InsideStruct struct {
        SampleStruct
        SomeOtherData string // that won't be serialized. 
    }
```

Or by composition
```go
    type InsideStruct struct {
        Outside SampleStruct
        SomeOtherData string // that won't be serialized. 
    }
```

Both embedding and composition will allow you to add functions to this struct. Whereas using it directly will prevent you from doing so.

## Message layout and data expected to be found here

Expect the message to be mostly compatible with (upsilontools' actor message)[github.com/ecumeurs/upsilontools])

Which means: 

* It will have a unique identifier (as uuid.UUID) 
* It will provide a Target struct ( with a string and a custom struct used as input parameters). see (method)[method/README.md] for more details.
* A Content struct (same). see (data)[data/README.md] for more details.
  



