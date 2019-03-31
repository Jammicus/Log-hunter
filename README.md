# Log Hunter 

This programs aim is to to allow you to collect a series of logs from remote host and bring them to your local machine in a timely manner.
Currently only supports connections over SSH.

## Prequisites

* Go 1.12.1

## Installing 

```
git clone https://github.com/Jammicus/log-hunter.git
go build .
```
 
## Running the binary

* Download Binary
* Make sure binary is executable
* If you are using encrypted passwords, make sure to pass in the passphrase 
* Ensure node config is stored in a file called "hosts.yml"

```
./log-hunter -hostsFile <pathTohost.yml> 

// If using Encrypted passwords in hosts.yml

./log-hunter -hostsFile <pathTohost.yml>  -passphrase <hash>
```

## Encrypting passwords which you want to store in hosts.yml

* Ensure you create a passphrase which is 20 characters long

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
## The Future

In the future we have plans to allow WinRM connections.
