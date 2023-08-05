package main

import (
	"bufio"
	"fmt"
	"github.com/lnstchtped/OpenseaScraper/pkg"
	"log"
	"os"
	"strings"
)

const (
	ColorCyan  = "\033[36m"
	ColorGreen = "\033[32m"
	ColorReset = "\033[0m"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic: %v", r)
		}
	}()

	cli()
}

func cli() {
	collection := askQuestion("Input collection slug (e.g. 'milady' for https://opensea.io/collection/milady)")
	proxy := askQuestion("Input proxy (in the format 'http://user:pass@host:port', can be left blank)")

	err := pkg.RunDownloader(collection, proxy)
	if err != nil {
		log.Printf("Error downloading: %v", err)
	}

	askQuestion("Press enter to exit")
}

func askQuestion(question string) string {
	PrintQuestion := fmt.Sprintf("%s?%s %s: %s", ColorGreen, ColorReset, question, ColorCyan)
	fmt.Print(PrintQuestion)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print(ColorReset)

	return strings.TrimSpace(input)
}
