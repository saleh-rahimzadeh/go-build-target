go-build-target
===============

The **go-build-target** is a build target provider for Golang projects.

---

## Install

Get:

```sh
go get -u github.com/saleh-rahimzadeh/go-build-target
```

Import:

```go
import (
  "github.com/saleh-rahimzadeh/go-build-target"
)
```

---

## Use

Get instance:

```go
var b buildtarget.BuildTarget = buildtarget.DEVELOP
```

Current build status:

```go
fmt.Println(buildtarget.Status)
```

Check build targets:

```go
switch buildtarget.Status {
case DEVELOP:
case RELEASE:
}
```

---

## Conditional target building

Develop file:

Create `<file>.develop.go` file:

```go
//go:build !release

package myapp

const Address string = "http://localhost:8080"
```

Create `<file>.release.go` file:

```go
//go:build release

package myapp

const Address string = "http://www.myapp.com"
```

Use in `main.go`:

```go
fmt.Println("Address: ", Address)
```

Build for develop:

```sh
go build
```

Build for release:

```sh
go build -tags release
```

---

## Contributing

Install `stringer`:

```sh
go get -tool golang.org/x/tools/cmd/stringer
```

Execute:

```sh
go generate
# OR
go tool stringer -type BuildTarget -linecomment -output build-target_string.go
```
