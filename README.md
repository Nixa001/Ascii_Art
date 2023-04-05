## ASCII-ART

*  A small Go program that generates ASCII art from a string input. 


## Installation

 - To use this program simple clone this Git repository on your local machine

```bash
  git clone https://learn.zone01dakar.sn/git/nifaye/ascii-art-output
```
-  Open a Terminal in the repo and install go packages
```bash
  apt install golang
```
- Usage: 
```console
go run . [OPTION] [STRING] [BANNER]
```

- EX: ``` go run . --output=<fileName.txt> something standard ```

```bash
$ go run . --output=banner.txt "Nifaye" shadow
$ cat -e banner.txt
```
```
                                                  $
_|      _| _|     _|_|                            $
_|_|    _|      _|       _|_|_| _|    _|   _|_|   $
_|  _|  _| _| _|_|_|_| _|    _| _|    _| _|_|_|_| $
_|    _|_| _|   _|     _|    _| _|    _| _|       $
_|      _| _|   _|       _|_|_|   _|_|_|   _|_|_| $
                                      _|          $
                                  _|_|            $
```
## Author

- [@nifaye](https://learn.zone01dakar.sn/git/nifaye)

![Logo](https://go.dev/images/go-logo-white.svg)
