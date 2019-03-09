package main

import (
	"flag"
	"fmt"
	"log-hunter/connection"
	"log-hunter/encryption"
	"log-hunter/parser"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var defaultPassphase = encryption.Passphrase
var hostsFileFlag *string
var encryptPassFlag *string

func main() {
	var waitGroup sync.WaitGroup
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	if *encryptPassFlag != "" {
		s := strings.Split(*encryptPassFlag, ":")
		encryption.Passphrase = s[0]
		fmt.Println(encryption.Encrypt(s[0]))
		return
	}

	nodes := parser.Parse(*hostsFileFlag)
	if encryption.Passphrase == defaultPassphase {
		log.Warn("Warning - Using default passphase. It is advised that you use your own custom passphrase")
	}

	log.Info("Getting information on hosts from:", *hostsFileFlag)

	waitGroup.Add(len(nodes))
	// index, node
	for e, _ := range nodes {

		go func(e int) {
			connection.GetLog(nodes[e].Host, nodes[e].Username, nodes[e].Password, nodes[e].LogLocation, nodes[e].DownloadDirectory, nodes[e].LogName, nodes[e].Port)
			waitGroup.Done()
		}(e)
	}
	waitGroup.Wait()
}

func usage() {
	log.Errorln("Invalid flag, please use one of the following:")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	passPhraseFlag := flag.String("passphrase", encryption.Passphrase, "PassPhrase used to decrypt passwords")
	hostsFileFlag = flag.String("hostsFile", "hosts.yml", "Path to host file")
	encryptPassFlag = flag.String("encrypt", "", "Takes <passphase:password> and returns the encrypted password")
	flag.Usage = usage
	flag.Parse()
	encryption.Passphrase = *passPhraseFlag
}
