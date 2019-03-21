package parser

import (
	"log"
	"os"

	"github.com/Jammicus/log-hunter/encryption"
	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Nodes    []nodeInfo    `yaml:"nodes"`
	Defaults []defaultInfo `yaml:"defaults"`
}

type defaultInfo struct {
	Username          string `yaml:"username"`
	Port              string `yaml:"port"`
	Connection        string `yaml:"connection"`
	LogLocation       string `yaml:"logLocation"`
	DownloadDirectory string `yaml:"downloadDirectory"`
	LogName           string `yaml:"logName"`
	DeleteLog         string `yaml:"deleteLog"`
	Checksum          string `yaml:"checksum"`
}

type nodeInfo struct {
	Host              string `yaml:"hostname"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	EncryptedPassword string `yaml:"encryptedPassword"`
	Port              string `yaml:"port"`
	Connection        string `yaml:"connection"`
	LogLocation       string `yaml:"logLocation"`
	DownloadDirectory string `yaml:"downloadDirectory"`
	LogName           string `yaml:"logName"`
	DeleteLog         string `yaml:"deleteLog"`
	Checksum          string `yaml:"checksum"`
}

type Node struct {
	Host              string
	Username          string
	Password          string
	Port              string
	Connection        string
	LogLocation       string
	DownloadDirectory string
	LogName           string
	DeleteLog         string
	Checksum          string
}

func getNode(defaultList defaultInfo, nonDefaultList nodeInfo) Node {
	node := Node{Host: nonDefaultList.Host,
		Username:          isDefault(defaultList.Username, nonDefaultList.Username),
		Password:          passwordHandler(nonDefaultList.EncryptedPassword, nonDefaultList.Password),
		Port:              isDefault(defaultList.Port, nonDefaultList.Port),
		Connection:        isDefault(defaultList.Connection, nonDefaultList.Connection),
		LogLocation:       isDefault(defaultList.LogLocation, nonDefaultList.LogLocation),
		DownloadDirectory: isDefault(defaultList.DownloadDirectory, nonDefaultList.DownloadDirectory),
		LogName:           isDefault(defaultList.LogName, nonDefaultList.LogName),
		DeleteLog:         isDefault(defaultList.DeleteLog, nonDefaultList.DeleteLog),
		Checksum:          isDefault(defaultList.Checksum, nonDefaultList.Checksum),
	}

	return node

}

func isDefault(defaultVal, nonDefaultValue string) string {
	if nonDefaultValue == "" {
		return defaultVal
	}
	return nonDefaultValue
}

func passwordHandler(encrypted, nonEncrypted string) string {
	if encrypted != "" {
		return encryption.Decrypt(encrypted)
	}
	return nonEncrypted
}

// Parse Gathers node information from yaml file
// Takes a path to the config file
// Returns a list of nodes.
func Parse(hostFile string) []Node {
	f, err := os.Open(hostFile)
	if err != nil {
		log.Fatal("Unable to open hosts file", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	var config config
	err = dec.Decode(&config)
	if err != nil {
		log.Fatal("Unable to decode Yaml config", err)
	}
	// fmt.Printf("Decoded YAML dependencies: %#v\n", config.Nodes)
	// fmt.Printf("Decoded YAML dependencies: %#v\n", config.Defaults)

	nodes := make([]Node, len(config.Nodes))

	for j := 0; j < len(config.Nodes); j++ {
		if len(config.Defaults) == 0 {
			nodes[j] = getNode(defaultInfo{}, config.Nodes[j])
			break
		}
		// Assume only ever 1 default block if defaults are present.
		nodes[j] = getNode(config.Defaults[0], config.Nodes[j])
	}

	return nodes
}
