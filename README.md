# askyes

Asks yes/no question.

## Install

```sh
go get -u github.com/tetafro/askyes
```

## Usage

Supported answers (case insensitive):

- yes/no
- y/n

```sh
$ askyes Do something? && echo Answer is yes
Do something? y
Answer is yes

$ askyes Do something? && echo Answer is yes
Do something? n
```
