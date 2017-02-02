# macpack
[![Build Status](https://travis-ci.org/murlokswarm/macpack.svg?branch=master)](https://travis-ci.org/murlokswarm/macpack)
[![Go Report Card](https://goreportcard.com/badge/github.com/murlokswarm/macpack)](https://goreportcard.com/report/github.com/murlokswarm/macpack)
[![Coverage Status](https://coveralls.io/repos/github/murlokswarm/macpack/badge.svg?branch=master)](https://coveralls.io/github/murlokswarm/macpack?branch=master)

Program to package MacOS apps. 
Will create a ```.app```.

## Install
```
go get -u github.com/murlokswarm/macpack
```

## Usage
In a directory where the package name is ```main```,
```
macpack build
```
**macpack** will automatically create these directories and files:
- resources
- resources/css
- mac.json

## Customize build
Build can be customized by modifying ```mac.json```.

eg:
```json
{
  "name": "Jubiz",
  "id": "maxence.Jubiz",
  "version": "1.0",
  "build-number": 4243,
  "icon": "jubiz.png",
  "dev-region": "en",
  "deployment-target": "10.12",
  "copyright": "Copyright © 2017 maxence. All rights reserved",
  "role": "None",
  "category": "public.app-category.photography",
  "sandbox": true,
  "capabilities": {
    "network": {
      "in": false,
      "out": true
    },
    "hardware": {
      "camera": false,
      "microphone": false,
      "usb": false,
      "printing": false,
      "bluetooth": false
    },
    "app-data": {
      "contacts": false,
      "location": false,
      "calendar": false
    },
    "file-access": {
      "user-selected": "",
      "downloads": "",
      "pictures": "",
      "music": "",
      "movies": ""
    }
  },
  "app-store": false,
  "supported-files": []
}
```

## Sass
[Sass](http://sass-lang.com/guide) can be used with:
```
macpack sass
```

This launches sass with watch mode.
**.scss** from ```resources/scss``` will be converted to **.css** in ```resources/css```.

Require sass to be installed:
```
sudo gem install sass
```
