# Installation Details

- If you want to upgrade from an older version, first remove previous version

## Delete an Existing version

- To delete an existing version delete the go directory : 
    - For Linux & Mac : 
        - `sudo rm -rf /usr/local/go`
    
## Install Go

- To install Go run these commands : 
    - `go_version=1.16.6`
    - `cd ~/Downloads`
    - `sudo apt-get update`
    - `sudo apt-get install -y build-essential git curl wget`
    - `wget https://dl.google.com/go/go${go_version}.linux-amd64.tar.gz`
    - `sudo tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz`
    - `sudo chown -R $(id -u):$(id -g) /usr/local/go`
    - `rm go${go_version}.linux-amd64.tar.gz`
    
## Adding Go in `$PATH` variable

- `mkdir $HOME/go`
- `nano ~/.bashrc`
- ```
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
  ```
- `source ~/.bashrc`


## Check Installation

- `go version` : To check whether you correctly installed Go

## Goland IDE

- `https://www.jetbrains.com/go/` : Install Goland IDE from here 