# Log-replay

`log-replay` is a tool for replaying a given log at a given speed.  The tool is designed to be CLI friendly, allowing pipes where possible.

## Installation instruction
## dependency

Updare the system then install the goland dependency
```bash
$ sudo apt-get update && sudo apt-get upgrade -y

```bash
sudo apt-get install goland


As of now I have notice that the test portion of the Makefile is broken so if you try to build using ```bash $ sudo make all or ```bash $ sudo make  the compilation will fail. Instead use the commands below.
```bash
$ sudo make build && sudo make release 

The release will give the compress version of the files while the build will compile the code. If you just want to compile the code then
```bash 
$ sudo make build

 
## Usage

```bash
$ log-replay -h
Usage of log-replay:
  -i string
    	Input file (if not using stdin)
  -s int
        Lines per second (default 100)
```

`log-replay` accepts input in two ways, you can either pass a file path with the `-i` flag, or you can pipe input into stdin.  All output is written to stdout.

## Examples

Replaying a log file from stdin at 2 lines/sec:

```bash
$ log-replay -s 2 < input.txt
```

Replaying a log file from stdin to syslog:

```bash
$ log-replay -s 2 < input.txt | logger -t test
```

Replaying a log file from stdin to a remote logger

```bash
$ log-replay -s 2 < input.txt | logger -t test -n loghost.example.com
```
