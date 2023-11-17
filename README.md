# AniCat CLI


## About

AniCat-Cli is a command-line client used to control AniCat

## Usage
add cli path to environment variables
```shell
windows:
setx PATH "%PATH%;{anicat-cli path}" /M
linux:
vim ~/.bash_profile 
# append to last line
export PATH=$PATH:{anicat-cli path}
```


```shell
linux:
go build -o anicat
windows:
go build -o anicat.exe
```



use "anicat [command] --help" for more information about a command:

```shell
anicat --help
```
information about add command 
```shell
anicat add --help
```

