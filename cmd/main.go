package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// import (
// 	"github.com/jcaroba/golang-k8s-context/internal/myMenu"
// )

var Config struct {
	Contexts []struct {
		Name string `yaml:"name"`
	} `yaml:"contexts"`
}

func main() {
	// myMenu.MyMenu()
	MyMenu()
}

func MyMenu() {
	fmt.Println("         / \\__")
	fmt.Println("        (    @\\___")
	fmt.Println("        /         O")
	fmt.Println("       /   (_____/")
	fmt.Println("      /_____/   U")
	fmt.Println("")
	fmt.Println(" K8s Context management ")
	fmt.Println("--------------------------")
	data, error := os.ReadFile("/Users/juarez.carpes/.kube/config")
	if error != nil {
		panic(error)
	}

	myContexts := Config

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

	fmt.Printf("Switching to %s", myContexts.Contexts[chosenInt].Name)
}
