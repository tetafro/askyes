# askyes

Asks yes/no question.

Supported answers (case insensitive):

- yes/no
- y/n

## Usage

```sh
$ askyes Do something? && echo Answer is yes
Do something? y
Answer is yes

$ askyes Do something? && echo Answer is yes
Do something? n
```
