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

var Config struct {
	Contexts []struct {
		Name string `yaml:"name"`
	} `yaml:"contexts"`
}

func main() {
	MyMenu()
}

func MyMenu() {
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

	dirname, _ := os.UserHomeDir()
	kubeConfig := dirname + "/.kube/config"

	// Read kubeconfig
	data, error := os.ReadFile(kubeConfig)
	if error != nil {
		panic(error)
	}

	myContexts := Config

	// Unmarshal yml
	err := yaml.Unmarshal([]byte(data), &myContexts)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(myContexts.Contexts); i++ {
		fmt.Println(i, "=>", myContexts.Contexts[i].Name)
	}
	fmt.Println("")
	fmt.Println("Please type the context number:")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	chosen := input.Text()
	chosenInt, _ := strconv.Atoi(chosen)

	if chosenInt < 0 || chosenInt >= len(myContexts.Contexts) {
		fmt.Printf("Invalid Option\n\n")
		os.Exit(0)
	}

	fmt.Printf("Switching to %s", myContexts.Contexts[chosenInt].Name)

	cmd := exec.Command("kubectl", "config", "use-context", myContexts.Contexts[chosenInt].Name)
	fmt.Println("\n\n")

	if err := cmd.Run(); err != nil {
		log.Fatal("Invalid option")
	}
}
