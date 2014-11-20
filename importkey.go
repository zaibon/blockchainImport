package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type input struct {
	Keys []struct {
		Label string `json:"label"`
		Priv  string `json:"Priv"`
	} `json:"keys"`
}

var (
	keysFile = flag.String("keys", "export.json", "the export file given by blockchain.org under bitcoin-qt format")
	confFile = flag.String("conf", homeDir(), "the export file given by blockchain.org under bitcoin-qt format")
	bin      = flag.String("bin", "/usr/local/bin/bitcoin-cli", "path to the bitcoin-cli binary")
)

func homeDir() string {
	return path.Join(os.Getenv("HOME"), ".bitcoin", "bitcoin.conf")
}

func main() {
	var err error
	flag.Parse()
	if *keysFile == "" {
		fmt.Println("you have to provide the file where the keys are stored")
		return
	}

	f, err := os.Open(*keysFile)
	if err != nil {
		log.Fatalln(err)
	}

	input := input{}
	if err = json.NewDecoder(f).Decode(&input); err != nil {
		log.Fatalln(err)
	}

	for _, v := range input.Keys {
		fmt.Printf("importing private key: %s (%s)...\n", v.Label, v.Priv)

		confArg := fmt.Sprintf("-conf=%s", *confFile)
		cmd := exec.Command(*bin, confArg, "importprivkey", v.Priv)
		if v.Label != "" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("\"%s\"", v.Label))
		}

		cmdStderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("run :", cmd.Path, strings.Join(cmd.Args, " "))
		err = cmd.Start()
		if err != nil {
			log.Fatalln(err)
		}

		r := bufio.NewReader(cmdStderr)
		for line, err := r.ReadString('\n'); err == nil; {
			fmt.Println(line)
		}

		err = cmd.Wait()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("done")
}
