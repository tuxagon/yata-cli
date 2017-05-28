# Yata Roadmap

## General
- `[x]` Add `timestamp` field to task
- `[ ]` ~~Replace an existing task with a new one~~
- `[ ]` Tasks are update-able through additions, changes, or deletions
- `[x]` Apply tags to any task `#sometag`
- `[x]` Allow filtering based on tag
- `[x]` List all tags
- `[ ]` ~~Change task from completed to incomplete?~~
- `[ ]` Allow aliases to be set up to simplify commands
- `[x]` Command line colors
- `[ ]` ~~File-based and SQLite storage options~~
- `[x]` Create a separate backup file if one already exists
- `[ ]` Something regarding subtasks should be figured out
- `[ ]` Add super-friendly messages and possibly allow for customizing messages based on names and configuration (like adding pirate-y messages)
- `[ ]` ~~Allow for messages to be suppressed unless there are errors~~

## Add command
- `[x]` Add new task with only required fields {description}

## Archive command
- `[x]` Allow archiving of tasks
- `[ ]` ~~Should archiving based on a criteria be permitted?~~

# Complete command 
- `[x]` Mark a task as completed

## Config
- `[x]` Handle "autoincrement" for files using `.yataid` file
- `[x]` Config file `.yataconfig`

# Delete command
- `[x]` Delete tasks

# Import/Export commands
- `[ ]` Import tasks from json
- `[ ]` Export tasks to json

## List command
- `[x]` List all tasks unfiltered, unsorted
- `[x]` List all tasks sorted on specified field {created_at,description,priority}
- `[x]` List all tasks that meet criteria {filter,completed,tag}
- `[x]` Change how the list renders {simple,verbose,printf}

## Prune command
- `[x]` Prune completed tasks from database/file

## Push/Fetch command
- `[ ]` Push the tasks data to a server {Google Drive,Dropbox}
- `[ ]` Fetch tasks data from a server {Google Drive,Dropbox}

## Reset command
- `[x]` Create a nuclear option to just restart

## Show command
- `[x]` Show a specific task based on ID

## Tutor command
- `[ ]` Create a tutorial type command, like `yata tutor`


This is not listed in any particular order and some of these are just thoughts and may change or just not get implemented. Roadmap = Braindump right now.