# Easy context switching ( k8s )

When you use kubectl to interact with more than one clusters, you have to type some commands to switch between the different contexts. This is a very simple solution to help you on that.


## How it works

This automation code will read the ~/.kube/config and will give the possibility to choose the context you want to switch on. It is important to use alias to define a friendly name on each context.

![screen](./images/screen.png)

## Build & Installation

PLease change the target-OS and target-architecture on the command bellow.

`env GOOS=target-OS GOARCH=target-architecture go build -o thor cmd/main.go`

After you build you can move the output file called thor to a bin directory that is on your path.