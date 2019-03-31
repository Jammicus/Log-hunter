package connection

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/Jammicus/log-hunter/parser"
	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

// GetLog establishes the connection to the remote server and copys the log file to the local machine
func GetLog(node parser.Node) {

	log.Info("###### Connection Information ##############", "\n",
		"Host = ", node.Host, "\n",
		"Username = ", node.Username, "\n",
		"Log Location =  ", node.LogLocation, "\n",
		"Log Name = ", node.LogName, "\n",
		"Download Directory = ", node.DownloadDirectory, "\n",
		"Connection Port = ", node.Port, "\n",
		"Checksum Algorithm = ", node.Checksum, "\n",
		"#############",
	)

	var client *ssh.Client
	var err error
	switch {
	case node.KeyLocation != "":
		client, err = connectSSHKey(node.Host, node.Username, node.KeyLocation, node.Port)

	case node.Password != "":
		client, err = connectSSHPass(node.Host, node.Username, node.Password, node.Port)

	default:
		log.Fatal("No Authentication method found. Terminating")
	}

	if err != nil {
		log.Fatal("Unable to establish connection:", err)

	}
	defer client.Close()

	copyFile(node.LogLocation, node.DownloadDirectory, node.LogName, node.DeleteLog, node.Checksum, client)

}

func copyFile(logLocation, downloadDirectory, filename, deleteLog, checksum string, client *ssh.Client) {
	address := client.RemoteAddr()

	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Unable to start session", err)
	}
	defer sftp.Close()

	remoteLog, err := sftp.Open(logLocation + filename)
	if err != nil {
		log.Fatal("Unable to open log file on the remote host", err)
	}

	defer remoteLog.Close()

	localLog, err := os.Create(makeDownloadDirectory(downloadDirectory, address.String()) + filename)
	if err != nil {
		log.Fatal("Unable to close log file on the remote host", err)
	}

	defer localLog.Close()
	log.Info("Writing log to ", localLog.Name())

	if _, err := io.Copy(localLog, remoteLog); err != nil {
		log.Fatal("Error writing remote file to local file", err)
	}

	if checksum != "" {
		if err := verifyFileIntegrity(remoteLog, localLog, checksum); err != nil {
			log.Fatal(err)
		}
	}

	if deleteLog == "true" {
		if err := removeRemoteLog(remoteLog.Name(), sftp); err != nil {
			log.Fatal(err)
		}
	}

	log.Info("Writing log to ", localLog.Name(), " complete")
}

func verifyFileIntegrity(remote, local io.Reader, algo string) error {
	var remoteHash, localHash hash.Hash

	switch algo {
	case "sha512":
		remoteHash = sha512.New()
		localHash = sha512.New()
	case "sha256":
		remoteHash = sha256.New()
		localHash = sha256.New()
	case "sha1":
		remoteHash = sha1.New()
		localHash = sha1.New()
	case "md5":
		remoteHash = md5.New()
		localHash = md5.New()
	default:
		return errors.New("Invalid Hashing Algo " + algo)
	}

	if _, err := io.Copy(remoteHash, remote); err != nil {
		log.Error("Error calculating remote hash", err)
		return errors.New("Error calculating remote hash")
	}

	if _, err := io.Copy(localHash, local); err != nil {
		log.Error("Error calculating local hash", err)
		return errors.New("Error calculating local hash ")
	}

	if reflect.DeepEqual(localHash.Sum(nil), remoteHash.Sum(nil)) {
		log.Debugf("Remote Hash: %x\n", remoteHash)
		log.Debugf("Local Hash: %x\n", localHash)
		return errors.New("Checksums do not match")
	}
	return nil
}

// Relative to where the binary is ran.
func makeDownloadDirectory(downloadDirectory, address string) string {

	downloadDirectory = downloadDirectory + strings.Split(address, ":")[0] + "/"
	if _, err := os.Stat(downloadDirectory); !os.IsNotExist(err) {
		log.Info("Download directory " + downloadDirectory + " already exists, doing nothing")
		return downloadDirectory
	}

	log.Info("Making download directory: " + downloadDirectory + " directory ")

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

func connectSSHPass(host, user, password, port string) (*ssh.Client, error) {

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

func connectSSHKey(host, user, keyLocation, port string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
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
