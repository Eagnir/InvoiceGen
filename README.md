InvoiceGen is a simple invoice generator written in Go Language (golang).
**(Project Under Development)**
- [Introduction](#Introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Version History](#version-history)
- [Building from Source](#building-from-source)
- [Tasks](#tasks)
- [Troubleshoot](#troubleshoot)
- [Contribute](#contribute)

# Introduction
This is a simple Invoice Generator (InvoiceGen) project, developed in Go Language (golang) + VUE.js with the following feature set.

1. Maintain records of Clients for a company with their basic details
2. Maintain Company details
3. Record Invoice details with individual InvoiceItem entries
4. CURD Tax information, in the form of Tax and TaxGroup
5. Ability to Tag an Invoice or Client with custom Tags
6. Ability to take HTML file (Template) for printing PDF Invoice
7. Ability to CURD Admin User for system access

# Installation
Simply download the binary for your platform and run the appropriate interface file as given in the below interface section.

- Windows:  (release coming soon)
- MacOS: (release coming soon)

# Usage
You can access the system in one of two ways. Either by a Web interface or via the command line (coming soon)

## Web Interface
To access the web interface, run the downloaded binary with the '-web' in the name, it will start a web server listening on the localhost on port 8080.

You can than access it via http://localhost:8080 from your favourite browser.

To change the port it listens on use `-port <port_number>` as a start up arguement on the binary

## Command Line Interface
Coming soon...

# Version History
- v1.0 - (Current)

# Building from Source
This section is for developers interested in building from source or simply want to take a look under the hood.

## Technology Used

### Frontend Progressive Web App (development pending, this can change)
- VUE.js - @vue/cli v4.5.12 [Office Site](https://vuejs.org)
- Typescript - [Office Site](https://www.typescriptlang.org)

### Backend
- Go (golang) - go1.16.3 for darwin/amd64 [Office Site](https://golang.org)
- GORM (ORM) - [Office Site](https://gorm.io)
- Echo (Web Framework) - [Office Site](https://echo.labstack.com)
- SQLite Database - [Office Site](https://www.sqlite.org)
- Make - v3.81 for i386-apple-darwin11.3.0 [Office Site](https://www.gnu.org/software/make/)

Full list of dependencies in [go.mod file](./go.mod)

## Build Process on MacOS
---
To build the project certain dependencies are required.

### Dependencies
Installing the Go Language (golang) from their official page
```
https://golang.org/dl/
```

To ensure you can cross-compile with CGO enabled from MacOS to Windows
```
brew install mingw-w64
```

Installing Make and GCC via XCode from Terminal
```
xcode-select --install
```

### Building from Source
Once you have all the dependencies installed building the project is a simple as a single line of code.

All binary are compiled into their respective platform folder under `./bin/` folder.
***Build all platforms***
```
make build
```
***Build only for MacOS***
```
make build-macos
```
***Build only for Windows***
```
make build-windows
```
The above commands will build the binaries for both the interfaces of the selected platform, if you like to build only web or command line interface, you can use the following.

**Web Interface Only**
```
make build-web-macos
```
or
```
make build-web-windows
```
**Command Line Interface Only**
```
make build-cli-macos
```
or
```
make build-cli-windows
```
## Run from Source Code
---
If you are tinkering with the source code or simply want to modify it. It is best to run the project from source rather then constantly compile.

**Run Project's Web Interface**
```
go run ./interface/web/main.go
```
which will start the web server and listen on the port 8080 on localhost, you can than visit the web interface at http://localhost:8080 from your favourite browser

**Run Project's Command Line Interface**
```
go run ./interface/cli/main.go
```
Remember the above command directly starts the execution so pass in any parameters you want at the end of the run command, for example.
```
go run ./interface/cli/main.go list -entity=adminusers
```
You can add a `json` flag at the end to output in json format (which exports more data), but it has to be the absolute last flag in the command, if it's not passed, output is in text format.


## Unit Tests
Will implement the unit test files soon, this would allow the developer to execute test before build and automatically identity if something fails.

It is a tedious task and will get it done, but in due time.

## Licensing

This code is released under the [GNU General Public License v3.0](./LICENSE). For more information, see the [LICENSE file](./LICENSE).

# Tasks
## Tasks for Core Functionality
- [ ] HTML to pdf library like wkhtmltopdf via go-wkhtmltopdf from https://github.com/SebastiaanKlippert/go-wkhtmltopdf
- [ ] List all the availale **|property-tags|** for html templates
- [ ] Upload and manage HTML templates
- [ ] Backup & restore functionality for all HTML templates + database + settings (if any)

## Tasks for Web Interface
- [ ] Implement web interface via VUE.js
- [ ] Implement http handler for all API calls
- [ ] Document API calls in README.md file

## Tasks for Command Line Interface
- [ ] Implement a basic listing command for all major entities
- [ ] Implement a creating an invoice command
- [ ] Implement a generating an invoice pdf command
- [ ] Document all commands in README.md file

# Troubleshoot
## Windows build error for cross-compile from MacOS
```
# runtime/cgo
gcc_libinit_windows.c:8:10: fatal error: 'windows.h' file not found
```
If you receive the following error, it is most likely due to wrong cross-compiler being used. Kindly use mingw-w64 gcc compiler with the appropriate `CC` flag while issuing the `build` command

## Linux build error for cross-compile from MacOS
---

```
# runtime/cgo
linux_syscall.c:67:13: error: implicit declaration of function 'setresgid' is invalid in C99 [-Werror,-Wimplicit-function-declaration]
linux_syscall.c:67:13: note: did you mean 'setregid'?
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/usr/include/unistd.h:593:6: note: 'setregid' declared here
```
It is clumbersum to say the least to cross-compile for different platforms when the libraries being used requires CGO to be enabled `CGO_ENABLED=1`.

For this reason for now, there is no linux platform support for this project, but hopefully I will add that in the future via a build process through a Docker container.

Some interesting reads on Cross Compilers
- https://blog.filippo.io/easy-windows-and-linux-cross-compilers-for-macos/
- https://github.com/mitchellh/gox
- https://github.com/karalabe/xgo


# Contribute

- If you want to contribute, simply fork it, branch it, change it, pull it.
- If you want feedback or suggestions or feature requests, send them across.
- Report issues on Github.
- Send pull requests for bugfixes and improvements.
- Send proposals on Github issues.