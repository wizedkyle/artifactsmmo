# artifactsmmo

This repo contains a CLI tool and pre-built bots to interact with [Artifacts MMO](https://artifactsmmo.com).
# CLI

`artifacts` CLI allows you to interact with the [Artifacts API](https://api.artifactsmmo.com/docs) and perform all functions from the command line.
The easiest way to authenticate the CLI is to use `artifacts configure -c <charactername> -t <token>`. This command will create a 
file under the current users home directory for example:
* Linux/macOS -> `~/.artifacts/credentials/creds.json`
* Windows -> `%USERPROFILE%\.artifacts\credentails\creds.json`

The token within the credentials file is encrypted using the machine ID as the key meaning the file cannot be copied to another machine and used. 
This downside is you will need to re-run `artifacts configure` on each machine you want to use the CLI on.

# Bots

The Bots below are built to automate certain tasks within Artifacts. They are designed to be run 24x7 and can be deployed on a virtual machine or using Docker via a container orchestration platform.

Unlike the authentication with the CLI bots can be authenticated using environment variables. The following environment variables are needed to authenticate successfully:
* CHARACTER_NAME
* TOKEN

## Miner


