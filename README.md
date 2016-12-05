# macpack
[![Go Report Card](https://goreportcard.com/badge/github.com/murlokswarm/macpack)](https://goreportcard.com/report/github.com/murlokswarm/macpack)
[![Coverage Status](https://coveralls.io/repos/github/murlokswarm/macpack/badge.svg?branch=master)](https://coveralls.io/github/murlokswarm/macpack?branch=master)

Program to package MacOS apps. 
Will create a ```.app```.

## Install
```
go get -u github.com/murlokswarm/macpack
```

## Usage
In a directory where the package name is main,
```go
package main
```

launch:

```
macpack
```

## Customize build
Build can be customized by providing a JSON named mac.json.

```js
{
    "name": "Hello",            // Name displayed in menu and dock
    "version": "1.0.0.0",       // x.x.x.x where x is non negative number
    "icon": "",                 // Name of .png relative to the resources dir
    "id": "murlok.Hello",       // UTI in reverse-DNS format with (A-Za-z0-9), (-) and (.) eg com.murlok.Hello-World
    "os-min-version": "10.12",  // >= 10.12
    "role": "None",             // Editor | Viewer | Shell | None
    "sandbox": true,            // Sandbox mode
    "supported-files": []       // Slice of UTI representing types of the supported files
}
```

This JSON is auto generated if nonexistent.
