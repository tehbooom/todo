# Todo app

<p align="center">
  <img src="https://goreportcard.com/badge/github.com/tehbooom/todo" title="ReportCard">
  <img src="https://raw.githubusercontent.com/tehbooom/todo/badges/.badges/main/coverage.svg" alt="Coverage">
</p>

Simple todo CLI application to add, edit, and remove tasks and also group them together by projects. 

I wanted a simple todo application that didnt have a lot of options or functionality. Todo apps in the browser or desktop dont work since I spend much of time in the terminal. 

Other tools like [Taskwarrior](https://taskwarrior.org/) are great however they released a breaking change and of a lot options that I simply do not need nor want.

This app outputs a table wtih the [Gruvbox](https://github.com/morhetz/gruvbox) color scheme which is the same theme I use everywhere else. May add an option in the future to support providing your own theme but at the moment it defaults to Gruvbox

## Install

```bash
go install github.com/tehbooom/todo
```

## Commands

The app supports specifying your own data file where tasks are stored. Default location is `$HOME/.td.json`

### `todo add`

Add a task to your todo list. Just add the task description and optionally specify the group that it should be in.

If the group does not exist it will create the group for you

### `todo edit`

Edit a task by ID number. You can edit the description of the task or the group in which the task is in.

Editing a task will update the timestamp of when it was "created"

### `todo list`

List tasks by group or all of them

### `todo rm`

Remove a task by ID. 

This is sort of saying its complete. I don't really care about what I have completed a seeing them that is why they get deleted.

### `todo group`

List available groups 

### `todo group add`

Add a group

### `todo group rm`

Remove a group and tasks in that group


## Man page

```bash
todo app thats simple and not unique allowing you to add, list and remove tasks

Usage:
  todo [command]

Available Commands:
  add         add a task
  completion  Generate the autocompletion script for the specified shell
  edit        Edit existing task
  group       Lists groups
  help        Help about any command
  list        list all tasks
  rm          Remove a task by ID number

Flags:
  -h, --help   help for todo

Use "todo [command] --help" for more information about a command.
```
