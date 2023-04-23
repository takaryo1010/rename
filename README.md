# rename
This was designed to be used as a command in git bash.
## Usage
```
go run . -p FromName,ToName ./your/dir
```
```
go run . -a hoge ./your/dir
```
## -h
```
Usage: C:\Users\user\mycommand\rename.exe -p or -a [ARGUMENTS] [PATH]
This is my command.
Options:
  -a string
        Renames all files in the specified directory to the argument + number.
        Usage:'-a NewName
  -p string
        Changes the specified part of a file name in the specified directory to the argument.
        Usage:'-p FromName,ToName
```
## article
