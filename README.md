[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/HugoSmits86/dreamcast-vms-icondata-tool/icondata?status.svg)](https://godoc.org/github.com/HugoSmits86/dreamcast-vms-icondata-tool/icondata)
[![Go Report Card](https://goreportcard.com/badge/github.com/m-zajac/json2go)](https://goreportcard.com/report/github.com/m-zajac/json2go) 
[![Build Status](https://travis-ci.com/HugoSmits86/dreamcast-vms-icondata-tool.svg?branch=master)](https://travis-ci.com/HugoSmits86/dreamcast-vms-icondata-tool) 
[![codecov](https://codecov.io/gh/HugoSmits86/dreamcast-vms-icondata-tool/branch/master/graph/badge.svg)](https://codecov.io/gh/HugoSmits86/dreamcast-vms-icondata-tool)

# Introduction

Icondata is a Go package for decoding and encoding ICONDATA.VMS file format.

I suspect that most users of this project are likely hobbyists instead of full-blown programmers.\
For them I have included a command-line tool in the form of main.go. This is a stand-alone\
program that lets users encode and decode ICONDATA.VMS files from the command-line.

Currently only the black and white icons are supported.

# Install package

The package includes a make file that can install the package for multiple platforms.

```Bash
#compile and install package for Windows
make install-windows
#compile and install package for MacOs
make install-macos
#compile and install package for Linux
make install-linux
```

# Build tool

The command-line tool includes a make file that can build the tool for multiple platforms.

```Bash
#compile and install package for Windows
make install-windows
#compile and install package for MacOs
make install-macos
#compile and install package for Linux
make install-linux
```

# Usage tool

The tool program can be used from the command-line. Here is an example:
```Bash
./icontool_macos -i test.vms -o test.png
```

# Todo

* Support color icons.
* Generate VMI file for VMS file.

# Disclaimer

This project is offered on a "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,\
either express or implied.

