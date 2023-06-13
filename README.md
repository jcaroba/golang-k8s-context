# Easy context switching ( k8s )

When you use kubectl to interact with more than one clusters, you have to type some commands to switch between the different contexts. This is a very simple solution to help you on that.


## How it works

This automation code will read the ~/.kube/config and will give the possibility to choose the context you want to switch on. It is important to use alias to define a friendly name on each context.

![screen](./images/screenhis.png)

## Build & Installation

`env GOOS=target-OS GOARCH=target-architecture go build package-import-path`

