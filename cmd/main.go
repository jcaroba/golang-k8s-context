package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Struct to unmarshal yml ~/.kube/config
var Config struct {
	Contexts []struct {
		Name string `yaml:"name"`
	} `yaml:"contexts"`
}

func main() {
	switchK8sContext()
}

func switchK8sContext() {
	fmt.Println("\n")
	fmt.Println("         / \\__")
	fmt.Println("        (    @\\___")
	fmt.Println("        /         O")
	fmt.Println("       /   (_____/")
	fmt.Println("      /_____/   U")
	fmt.Println("	Thor ")
	fmt.Println("")
	fmt.Println(" K8s Context Switching ")
	fmt.Println("--------------------------")

	// Get the profile .kube/config
	dirname, _ := os.UserHomeDir()
	kubeConfig := dirname + "/.kube/config"

	// Read .kube/config
	data, error := os.ReadFile(kubeConfig)
	if error != nil {
		panic(error)
	}

	// Creating object struct Config type
	myContexts := Config

	// Unmarshal yml
	err := yaml.Unmarshal([]byte(data), &myContexts)
	if err != nil {
		panic(err)
	}

	// Print each context that exist on the kubeconfig
	for i := 0; i < len(myContexts.Contexts); i++ {
		fmt.Println(i, "=>", myContexts.Contexts[i].Name)
	}
	fmt.Println("")
	fmt.Println("Please type the context number:")

	// Get the number of the chosen context
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	chosen := input.Text()
	chosenInt, _ := strconv.Atoi(chosen)

	// check if were chosen a existent context
	if chosenInt < 0 || chosenInt >= len(myContexts.Contexts) {
		fmt.Printf("Invalid Option\n\n")
		os.Exit(0)
	}

	// Change to the context you have chosen
	cmd := exec.Command("kubectl", "config", "use-context", myContexts.Contexts[chosenInt].Name)
	if err := cmd.Run(); err != nil {
		log.Fatal("Invalid option")
	}

	// Show the current context
	cmdGetCurrentContext := exec.Command("kubectl", "config", "current-context")
	output, _ := cmdGetCurrentContext.Output()
	fmt.Println("\nNow your Current Context is: ", string(output))

}
