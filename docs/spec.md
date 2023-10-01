# Mizu programming language specification

## Purpose

Mizu aims to be a simple programming language for algorithmic problem solving.

## Tokens

### Symbols

| Literal | Name             |
|---------|------------------|
|`.`      | dot              |
|`:`      | colon            |
|`;`      | semicolon        |
|`?`      | question mark    |
|`::`     | define constant  |
|`:=`     | define variable  |
|`(` `)`  | parenthesis      |
|`{` `}`  | curly braces     |
|`[` `]`  | brackets         |
|`[` `]`  | brackets         |
|`+`      | addition         |
|`-`      | subtract         |
|`*`      | multiplication   |
|`/`      | division         |
|`%`      | mod              |
|`<`      | less             |
|`>`      | greater          |
|`<=`     | less or equal    |
|`>=`     | greater or equal |

### Keywords

`if` `else` `switch` `case`

`loop` `skip` `break`

`return`

`not` `and` `or`

### Regex based tokens

| Regex               | Subpart     |
|---------------------|-------------|
| `[a-zA-Z]`          | letter      |
| `[0-9]`             | digit       |
| `[+-]`              | sign        |
| `letter | digit`    | alNum       |
| `digit (_? digit)*` | number      |
| `\"`                | escQuote    |
| `[^\"]`             | notEscQuote |

| Regex                   | Name   |
|-------------------------|--------|
| `letter alNum*`         | symbol |
| `sign? number`          | int    |
| `sign? number . number` | float  |
| `" notEscQuote* "`      | string |

## Grammar

## IR Code

## Bytecode
