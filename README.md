# Log Hunter (WIP.)

This programs aim is to to allow you to collect a series of logs from remote host and bring them to your local machine in a timely manner.
Currently only supports connections over SSH. Plans to implement WinRM in the future.

## Prequisites

* Go 1.12.1

## Installing 

```
git clone https://github.com/Jammicus/log-hunter.git
go get gopkg.in/yaml.v2 github.com/sirupsen/logrus gopkg.in/Luzifer/go-openssl.v3 github.com/pkg/sftp golang.org/x/crypto/ssh
```
 
## Running the binary

* Download Binary
* Make sure binary is executable
* If you are using encrypted passwords
* Ensure node config is stored in a file called "hosts.yml"

```
./log-hunter -hostsFile <pathTohost.yml> 

// If using Encrypted passwords in hosts.yml

./log-hunter -hostsFile <pathTohost.yml>  -passphrase <hash>
```

## Encrypting passwords which you want to store in hosts.yml

```
./log-hunter -encrypt <password> -passphrase <hash>

// Example

./log-hunter -encrypt examplepassword -passphrase z4yH36a6zerhfE5427ZV
```

## Running the Tests

```
// Unit tests
go test ./encryption ./parser
// Benchmarks
cd parser
go test -bench=<testname>
// Example:
go test -bench=BenchmarkDefault100
```

### Testing Locally using Vagrant

```
// Ensure you have done the installation steps above

vagrant up
go run main.go
```
