package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const delay = 5

func main() {

	showNames()
	showIntro()
	showMenu()

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Showing Logs...")
	// } else if command == 0 {
	// 	fmt.Println("Leaving..")
	// } else {
	// 	fmt.Println("Unknown Command")
	// }

	command := readCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		fmt.Println("Leaving..")
		os.Exit(0)
	default:
		fmt.Println("Unknown Command")
		os.Exit(-1)

	}
}

func showIntro() {
	name := "Naruto"
	var age = 17
	var version float32 = 1.1
	fmt.Println("Hello Sr.", name, "you are ", age, "years old")
	fmt.Println("This application is in version: ", version)

	fmt.Println("O tipo da variável é: ", reflect.TypeOf(age))
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("The command is", command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	//sites := []string{"https://www.alura.com.br/", "https://www.caelum.com.br/"}

	sites := readSitesFromFile()

	for i := 0; i < 300; i++ {
		for j, site := range sites {
			response, err := http.Get(site)

			if err != nil {
				fmt.Println("Ocorreu um erro: ", err)
			}

			if response.StatusCode == 200 {
				fmt.Println(i, ".", j, " - Site ", site, "Loaded Successfully Status Code", response.StatusCode)
				writeLogs(site, true)
			} else {
				fmt.Println(i, ".", j, " - Site ", site, "With Problem Status Code", response.StatusCode)
				writeLogs(site, false)
			}
		}
		time.Sleep(delay * time.Millisecond)
	}
}

func showNames() {
	names := []string{"Pedro", "Silva", "Ana"}
	names = append(names, "Cris")
	fmt.Println(names)
	fmt.Println(len(names), cap(names))
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Ocorreu um erro: ", err)
		}
	}

	file.Close()

	fmt.Println(sites)
	return sites
}

func writeLogs(site string, status bool) {

	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("Showing Logs...")
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(file))

}
