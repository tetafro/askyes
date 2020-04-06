# askyes

Asks yes/no question.

Supported answers (case insensitive):
    - yes/no
    - y/n

## Usage

```sh
$ askyes Do something?
Do something? y
$ echo $?
0

$ askyes Do something?
Do something? n
$ echo $?
1

$ askyes Do something?
Do something? what
Unknown type of answer
$ echo $?
1
```
