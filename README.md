# Linkr

## Table of Contents

- [Linkr](#linkr)
    - [Table of Contents](#table-of-contents)
    - [About](#about-)
    - [Current Caveats](#current-caveats)
    - [Getting Started](#getting-started-)
        - [Prerequisites](#prerequisites)
        - [Installing](#installing)
            - [With `just` on Linux](#with-just-on-linux)
            - [Without `just` on Linux](#without-just-on-linux)
    - [Ideas](#ideas)
    - [TO DO](#to-do)

## About <a name = "about"></a>

I am not a fan of having to use specific browsers full time to get certain web sites to work, or toget the full experience. This tool was build to give me the option to run a default browser of choice, and create "rules" based on a URL to open in a specific browser.

For example:

- I want a lightweight or minimal browser as a default
- I want Firefox for its container technology on social media
- I need Chrome for sites that work best or provide the full feature set on Chome

With this application set as the *default* http handler on my system, I can now write rules where the URL will match and direct to the appropriate browser.

## Current Caveats

- **I am just starting out, so this is a work in progress.** Let's call it a "Proof of Concept". But I am running it on my system daily.
- This document needs some work. It should tighten up once there are no longer hardcoded paths.
- Currently this does not work when clicking a link in browser. The browsers will keep the clicked links withing themselves due to the fact that they are not aware of Linkr. If I can ever figure out how to do this without learning (time constrained) to write browser extensions, I would likely add this feature. Currently Linkr should support any situation where clicking a link will leverage the default http(s) handler.
- Installation is still to be flushed out. Currently the desktop files and associated paths are hardcoded. And I use my justfile to install currently while I develop this.
- The way this `justfile` installs the app, I consider the process of setting Linkr as the default to be a little fragile. It expects certail handlers to exist in `~/.config/mimeapps.list` and simply used `sed` to replacve the desktop file referenced.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [Installing](#installing) for notes on how to deploy the project on a live system.

*NOTE* - I am developing on Linux, and do not have access to a Mac, or Windows. I have tried to put the Mac specific code in there, but it is untested. Windows has not been included at all at this time.

- Setup a Go dev environment setup
- Install `just` (Makefile alternative) if you wish to leverage the `justfile` in the repo
- Install `sed`
- Clone/Fork this repo
- To see the various `just` options for builds, etc, run the `just` command in the root of the repo
    - I use the `install-linux:` and `dev-rebuild:` mostly through dev
- You will want to edit the following for your use:
    - The `install-linux:` step in the `justfile` for your paths
    - The desktop file paths
    - create a `config.yaml` file in the `sample` directory. The file will be used in the `justfile` steps for testing, etc
        - If you are planning to change the behavior of the `config.yaml` please ensure you copy a good example back to `sample/cfg-sample.yaml`

### Prerequisites

- A recent Version of `go`
- The `just` tool installed
- The `sed` tool installed
- Running Linux
    - If you plan on testing Mac, have at it Hoss. I am just not sure how I can support it.
    - If you would like to add Windows support, open to discussion. I am just not sure how I can support it.

### Installing

At this point, the only installation method (early days) is to use the `just` command, or if you prefer, you can run the terminal commands manually from the `install-linux:` step in the `justfile`.

- Clone this repo
- Adjust the files:
    - You will want to edit the following for your use:
    - The `install-linux:` step in the `justfile` for your paths
    - The desktop file paths
    - create a `config.yaml` file in the `sample` directory (if you plan to use the `justfile`), or you can place it manually at `~/.config/linkr/config.yaml` (Linux). For Mac, it will look in `~/Library/Application Support/linkr/config.yaml`. MAc experts feel free to pipe up if there is a better location.

The command samples below are for Linux as I have not tested on a Mac.

#### With `just` on Linux

```shell
just deploy
```

#### Without `just` on Linux

follow the steps in the `justfile`, mod paths for your system/preferences, and run the commands in your terminal.

## Ideas

- [ ] Rules based on networks
- [ ] add wildcard support
- [ ] add a devnull blocker

## TO DO

- [x] create Linux Desktop File
- [x] create a "logo" of some kind
