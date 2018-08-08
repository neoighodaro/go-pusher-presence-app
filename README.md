# Create a Go application with online presence using Pusher

This is a demo on how to integrate realtime presence into web applications using Go, JavaScript and Pusher.

[View tutorial](https://pusher.com/tutorials/go-vue-online-presence)

## Prerequisites

* An IDE of your choice e.g. [Visual Studio Code.](https://code.visualstudio.com/)
* [Go](https://golang.org/doc/install) installed on your computer.
* Basic knowledge of GoLang.
* Basic knowledge of JavaScript (ES6 syntax) and Vue.
* Basic knowledge of using a CLI tool or terminal.
* Pusher application. Create one [here.](http://pusher.com/)

## Getting Started

To get started with the project, make sure you have all the prequiisites above.

1. Clone the project to your machine.
2. Update the Pusher keys in the chat.go, app.js and support.js files.
3. Run the command: $ go get github.com/pusher/pusher-http-go
4. Run the command: $ go run presence.go
4. Visit http://localhost:8090 to see application in action.

## Built With

* [Go](https://golang.org/doc/install) - Modern programming language.
* [Pusher](https://pusher.com/) - build realtime applications easily.
* [Echo](https://echo.labstack.com/) - Go web framework.
