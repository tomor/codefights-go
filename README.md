Codefights solutions
============================
In this repo I'm collecting my https://codefights.com/ solutions written in Go

It might help to somebody looking for the solution or it might improve my knowledge when I receive a feedback ;)

### Codefights solutions in other languages
- PHP: https://github.com/tomor/codefights-php


### Howto setup Intellij IDEA community go project with this repo
- git clone git@github.com:tomor/codefights-go.git
- File -> New -> Project from Existing Sources... 
- select the created "codefights-go" as the project folder
- select go SDK during the project creation wizard
- for usage with GVM we need to map the project to the $GOPATH
  - do once: cd codefights-go; gvm linkthis github.com/tomor/codefights-go
- try: go run codefights.go
- try: cd arcade/introgates; go test