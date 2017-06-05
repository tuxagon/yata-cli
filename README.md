## Yata CLI

[![Build Status](https://travis-ci.org/tuxagon/yata-cli.svg?branch=master)](https://travis-ci.org/tuxagon/yata-cli)

Yata is a command line task manager that works on Windows, Mac, and Linux. The idea for it came because I was frustrated with the other 
alternatives I found and I really only wanted a simple task manager. I really wanted git, but for tasks. This is what yata wants to be.

## Getting Help

Yata comes with 12 commands right now and each can be seen when you run

```
$ yata
$ yata help
$ yata --help
```

## Basic Workflow

A basic workflow with Yata

```shell
$ yata add "Replace the bat signal lightbulb" # Creates a new task
$ yata list # Lists the current open tasks
1 Replace the bat signal lightbulb
$ yata complete 1 # Marks the task with ID 1 as completed
```

## Tagging

Tags are used to essentially label a task

Tags can be added in 2 ways
1. Using the `--tags` flag allows you to specify a comma-delimited list of tags
2. Adding tags directly in the task description, prefixing the tag with a `#`

```shell
$ yata add "Join the #fellowship"
$ yata add "Bring the ring to Mordor" --tags fellowship
$ yata list
1 Join the #fellowship [fellowship]
2 Bring the ring to Mordor [fellowship]
```

## Priority

Tasks can be marked with a single priority
1: Low
2: Normal
3: High

## Synchronization

An important feature for me when developing this was to have the ability to push my tasks to a server and 
have the ability to download the tasks and essentially synchronize all my computers

Currently, Google Drive is the only server implementation

### Configuring Google Drive

Follow the instructions found on "Step 1" [here](https://developers.google.com/drive/v3/web/quickstart/go)

You should have a file called `client_secret.json` at the end, which should be placed in your `$HOME/.yata` directory (For Windows users, it would be `%USERPROFILE%/.yata`)

Once that file is in the `.yata` directory, you can just push and Yata will walk you through the rest of the setup, which involves authorizing the Yata.

**Note**: You will also need to run `yata config googledrive.secretfile client_secret.json`

### Usage

```shell
$ yata push -g # Pushes the tasks to Google Drive
$ yata fetch -g # Downloads the tasks found in Google Drive
$ yata merge # Merges the fetched tasks with the local tasks
```
