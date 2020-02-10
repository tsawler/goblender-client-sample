[![Go Report Card](https://goreportcard.com/badge/github.com/tsawler/goblender-client-sample)](https://goreportcard.com/report/github.com/tsawler/goblender-client-sample)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/tsawler/goblender/master/LICENSE)
[![Version](https://img.shields.io/badge/goversion-1.13.x-blue.svg)](https://golang.org)
<a href="https://golang.org"><img src="https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square" alt="Built with GoLang"></a> 


# GoBlender client

Sample code for client specific code for GoBlender.

## Setup

First, install [goblender](https://github.com/tsawler/goblender).

All client specific code lives in `./client/clienthandlers`, and is its own 
git repository. When working in JetBrains Goland, you must create 
(or clone) a git repository in this location, and then add the directory
in Preferences -> Version Control.

## Updating on server
Change  `update.sh` to enable git pull of client:

```
# uncomment if using custom client code
#cd ./client/clienthandlers
#git pull
#cd ../..
```

After changing, it should look like this:

```
# uncomment if using custom client code
cd ./client/clienthandlers
git pull
cd ../..
```

