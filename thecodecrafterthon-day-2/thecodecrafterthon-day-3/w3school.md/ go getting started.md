Go Get Started

To start using Go, you need two things:

1: a text editor like vs code or terminal.
2:a compiler that will translate the language into what the computer can understand.

==GO INSTALL
you can find very relevant info about installation on https://golang.org/dl/.
you can check if go was installed properly by running  (go version.)

==Go Install IDE 
An IDE is a smart workspace where you can write, fix, and run your code all in one place.  sugusing VS Code is better because it is the best and easiest choice for beginners. 

== Go Quickstart

to create a GO programm 

    Launch the VS Code editor
    Open the extension manager or alternatively, press Ctrl + Shift + x
    In the search box, type "go" and hit enter
    Find the Go extension by the GO team at Google and install the extension
    After the installation is complete, open the command palette by pressing Ctrl + Shift + p
    Run the Go: Install/Update Tools command
    Select all the provided tools and click OK

VS Code is now configured to use Go.

Open up a terminal window and type:
go mod init example.com/hello
 then create a new file . in it type hello word and save it.
helloworld.go

package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
} 
then run it.
to run a code (type go run helloworld,go)
