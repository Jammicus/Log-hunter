package connection

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/ssh"
)

// GetLog establishes the connection to the remote server and copys the log file to the local machine
func GetLog(host, user, password, logLocation, downloadDirectory, fileName, port, deleteLog string) {

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

	copyFile(logLocation, downloadDirectory, fileName, deleteLog, client)

}

func copyFile(logLocation, downloadDirectory, filename, deleteLog string, client *ssh.Client) {
	address := client.RemoteAddr()

	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Unable to start session")
	}
	defer sftp.Close()

	remoteLog, err := sftp.Open(logLocation + filename)
	if err != nil {
		log.Fatal("Unable to open log file on the remote host")
	}

	defer remoteLog.Close()

	localLog, err := os.Create(makeDownloadDirectory(downloadDirectory, address.String()) + filename)
	if err != nil {
		log.Fatal("Unable to close log file on the remote host")

	}

	defer localLog.Close()
	log.Info("Writing log to ", localLog.Name())
	remoteLog.WriteTo(localLog)

	rStat, rErr := remoteLog.Stat()
	lStat, lErr := localLog.Stat()
	if rErr != nil {
		log.Fatal("Error getting information on remote file", rErr)
	}

	if lErr != nil {
		log.Fatal("Error getting information on local file ", localLog)
	}

	if rStat.Size() != lStat.Size() {
		log.Fatal("File sizes do not match after transfer")
	}

	if deleteLog == "true" {
		err := removeRemoteLog(remoteLog.Name(), sftp)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Info("Writing log to ", localLog.Name(), " complete")
}

func makeDownloadDirectory(downloadDirectory, address string) string {

	downloadDirectory = downloadDirectory + strings.Split(address, ":")[0] + "/"
	if _, err := os.Stat(downloadDirectory); !os.IsNotExist(err) {
		log.Info("Download directory " + downloadDirectory + " already exists, doing nothing")
		return downloadDirectory
	}

	log.Info("Making download directory :" + downloadDirectory + " directory ")

	os.MkdirAll(downloadDirectory, os.ModePerm)

	return downloadDirectory
}

func removeRemoteLog(logLocation string, client *sftp.Client) error {
	err := client.Remove(logLocation)
	if err != nil {
		return errors.New("Unable to delete remote file " + logLocation)
	}
	log.Info("Deleted remote log" + logLocation)
	return nil
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
