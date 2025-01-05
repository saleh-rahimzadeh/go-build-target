go-build-target
===============

The **go-build-target** is a build target provider for Golang projects.

---

Install:

```sh
go get -u github.com/saleh-rahimzadeh/go-build-target
```

Import:

```go
import (
  "github.com/saleh-rahimzadeh/go-build-target"
)
```

Use:

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

Conditional target building:

- `<file>.develop.go`:

```go
//go:build !release

package myapp

const Address string = "http://localhost:8080"
```

- `<file>.release.go`:

```go
//go:build release

package myapp

const Address string = "http://www.myapp.com"
```

- `main.go`:

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
