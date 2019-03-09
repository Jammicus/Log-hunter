package connection

import (
	"errors"
	"os"
	"time"

	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/ssh"
)

func GetLog(host, user, password, logLocation, downloadDirectory, fileName, port string) {

	log.Info("######Connection Information ##############", "\n",
		"Host = "+host, "\n",
		"User = "+user, "\n",
		"Log Location =  "+logLocation, "\n",
		"Log Name = "+fileName, "\n",
		"Download Directory = "+downloadDirectory, "\n",
		"Connection Port = "+port, "\n",
		"####################")

	client, err := connect(host, user, password, port)
	if err != nil {
		log.Fatal(err)

	}
	defer client.Close()

	copyFile(logLocation, downloadDirectory, fileName, client)

}

func copyFile(srcPath, destPath, filename string, client *ssh.Client) {
	address := client.RemoteAddr()

	destPath = destPath + address.String() + "/"

	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		os.MkdirAll(destPath, os.ModePerm)

		log.Info("Making Logging Directory ", destPath)

	}

	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Unable to start session")
	}
	defer sftp.Close()

	srcFile, err := sftp.Open(srcPath + filename)
	if err != nil {
		log.Fatal("Unable to open log file on the remote host")
	}

	defer srcFile.Close()

	dstFile, err := os.Create(destPath + filename)
	if err != nil {
		log.Fatal("Unable to close log file on the remote host")

	}

	defer dstFile.Close()

	srcFile.WriteTo(dstFile)
	log.Info("Writing log to", destPath, "complete")
}

func connect(host, user, password, port string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
		Timeout: 30 * time.Second,
	}

	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return nil, errors.New("Unable to establish SSH Connection to " + host)
	}

	return client, nil
}
