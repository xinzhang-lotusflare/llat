# bulldozer

The `bulldozer` app is to provide access to LF VPN for China team when working from home.

## Install

### Dependencies

`Bulldozer` need **Homebrew** to install dependencies:

```
brew install wireguard-tools
```

The `wireguard-tools` needs bash 4+, so it should install a bash as well.

The installed dir of the `bash` should be **like** `/usr/local/Cellar/bash/5.1.16`. Inside the installed dir, there is an executable `bash` located under `bin` folder. So the path of executable of `bash` would be **like** `/usr/local/Cellar/bash/5.1.16/bin/bash`.

Take a note of it.

### Prepare workspace for `bulldozer`

Place the `bulldozer` executable in your laptop's PATH. For example, put it in the folder `/usr/local/bin`.

Then execute command

```
bulldozer install --bash <bash executable path>
```
The bash executable path is from above steps. It is `/usr/local/Cellar/bash/5.1.16/bin/bash` by the example.

It will create a folder `~/.bdz` and a file `bash` inside it. The content of `bash` is the path of executable.

## How to use

`Bulldozer` supports `-h` for help information at every level.

### Start bulldozer

```
sudo bulldozer run
```

### Shutdown bulldozer

```
sudo bulldozer stop
```

For `run` and `stop`, `bulldozer` needs **root** permission for setting up / tearing down virtual network interface.

*HINT:*

Although the `bulldozer` only transmits LF VPN related packets, we have found some apps may hang up when establishing network connections while the `bulldozer` is running. If you met any network issue, please shut `bulldozer` down at first.

## How to build

Prepare required variables claimed in file `variables.sh`.

[IMPORTANT: Please avoid submitting these credentials in any commit]()

Then execute script `pre-compile.sh`. It should generate a `wg_config.go` which contains the config for WireGuard client.

The last step is to execute command at the root of the repo:

```
go build
```

The executale `bulldozer` should be generated then.
