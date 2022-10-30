# LLat

LLat, is short for Legal LF Access Tool.

The `llat` app is to provide access to LF VPN **only** for China team when working from home.

`llat` is tested on MacOS only.

A detailed explanation may refer to the [page](https://lotusflare.atlassian.net/wiki/spaces/~87917946/pages/4073783773/How+LLat+works).

## Install

### Dependencies

`llat` need **Homebrew** to install dependencies:

```
brew install wireguard-tools
```

The `wireguard-tools` needs bash 4+, so it should install a bash as well.

The installed dir of the `bash` should be **like** `/usr/local/Cellar/bash/5.1.16`. Inside the installed dir, there is an executable `bash` located under `bin` folder. So the path of executable of `bash` would be **like** `/usr/local/Cellar/bash/5.1.16/bin/bash`.

Take a note of it.

### Install Binary

There are 2 ways to install `llat` binary.

1. Put the `llat` binary in your laptop's PATH. For example, put it in the folder `/usr/local/bin`. Make sure the binary file has executable permission

    ```
    sudo chmod a+x `which llat`
    ```
2. Uncompress llat.tar with command: `tar -xvf llat.tar && sudo mv llat /usr/local/bin/llat`

Both the binary and tar file are available in [Github releases](https://github.com/xinzhang-lotusflare/llat/releases).

You can execute `llat` directly in console to confirm it is installed correctly.

> If the MacOS alerts the binary cannot be verified, please set it as trusted in the `Security & Privacy` of `System Preferences`.


### Prepare workspace for `llat`

Once the `llat` is installed, we need to prepare workspace before using it.

```
llat install --bash <bash executable path> --config <path of config>
```
* The bash executable path is from above steps. It is `/usr/local/Cellar/bash/5.1.16/bin/bash` by the example.
* The config is provided by LLat admin.

It will create a folder `~/.llat` and file `bash`, `config` inside it.

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

`llat` is implemented in Go. Make sure you have Go developement environment prepared. Required Go version: > `1.19.0`.

Prepare required variables claimed in file `variables.sh`.

[IMPORTANT: Please avoid submitting these credentials in any commit]()

Then execute script `pre-compile.sh`. It should generate a `wg_config.go` which contains the config for WireGuard client.

The last step is to execute command at the root of the repo:

```
go build
```

The executale `llat` should be generated then.
