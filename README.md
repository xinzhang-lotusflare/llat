# LLat

LLat, is short for Legal LF Access Tool.

The `llat` app is to provide access to LF VPN **only** for China team when working from home.

`llat` is tested on MacOS only.

## Install

### Dependencies

`llat` need **Homebrew** to install dependencies:

```
brew install wireguard-tools
```

The `wireguard-tools` needs bash 4+, so it should install a bash as well.

The installed dir of the `bash` should be **like** `/usr/local/Cellar/bash/5.1.16`. Inside the installed dir, there is an executable `bash` located under `bin` folder. So the path of executable of `bash` would be **like** `/usr/local/Cellar/bash/5.1.16/bin/bash`.

Take a note of it.

### Prepare workspace for `llat`

Download the latest `llat` executable from [Github releases](https://github.com/xinzhang-lotusflare/llat/releases).

Place the `llat` in your laptop's PATH. For example, put it in the folder `/usr/local/bin`.

Make sure the binary file has executable permission

```
sudo chmod a+x `which llat`
```

> If the MacOS alerts the binary cannot be verified, please set it as trusted in the `Security & Privacy` of `System Preferences`.

Then execute command

```
llat install --bash <bash executable path>
```
The bash executable path is from above steps. It is `/usr/local/Cellar/bash/5.1.16/bin/bash` by the example.

It will create a folder `~/.llat` and a file `bash` inside it. The content of `bash` is the path of executable.

## How to use

`llat` supports `-h` for help information at every level.

### Start llat

```
sudo llat run
```

### Shutdown llat

```
sudo llat stop
```

For `run` and `stop`, `llat` needs **root** permission for setting up / tearing down virtual network interface.

*HINT:*

Although the `llat` only transmits LF VPN related packets, we have found some apps may hang up when establishing network connections while the `llat` is running. If you met any network issue, please shut `llat` down at first.

## How to build

`llat` is implemented in Go. Make sure you have Go developement environment prepared.

Prepare required variables claimed in file `variables.sh`.

[IMPORTANT: Please avoid submitting these credentials in any commit]()

Then execute script `pre-compile.sh`. It should generate a `wg_config.go` which contains the config for WireGuard client.

The last step is to execute command at the root of the repo:

```
go build
```

The executale `llat` should be generated then.
