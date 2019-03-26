# Sawari
CLI to automate fuel receipt 

Example Output can be viewed [here](./output-examples/hello.pdf) :rocket:

## Install

### if you have your `$GOPATH` configured
```go 
go get -u github.com/minhajuddinkhan/sawari/bin/sawari
```
### if you don't

### (linux)

```bash
wget https://github.com/minhajuddinkhan/sawari/releases/download/v1.0/sawari-linux  
chmod +x sawari-linux
sudo mv sawari-linux /bin/sawari
```
### (macOS)
``` bash
curl https://github.com/minhajuddinkhan/sawari/releases/download/v1.0/sawari-darwin -Lo sawari
chmod +x sawari
cp sawari usr/local/bin/
```

## Usage
```golang
sawari form --env [path-to-env-file]
```

#### Your env file should look like this 
```env
NAME="Minhaj Uddin Khan"
EMPLOYEE_CODE="877"
CEILING="5000"
DESIGNATION="Senior Software Developer (Grade -1)"
```
####  NOTE: If env not provided, the cli will prompt for information


