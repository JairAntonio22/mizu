# Mizu programming language specification

## Purpose

Mizu aims to be a simple programming language for algorithmic problem solving.

## Tokens

**Separators**
| Literal | Name             |
|---------|------------------|
|`:`      | colon            |
|`;`      | semicolon        |
|`?`      | question mark    |
|`(` `)`  | parenthesis      |
|`{` `}`  | curly braces     |
|`[` `]`  | brackets         |
|`[` `]`  | brackets         |

**Operators**
| Literal | Name             |
|---------|------------------|
|`::`     | define constant  |
|`:=`     | define variable  |
|`+`      | addition         |
|`-`      | subtract         |
|`*`      | multiplication   |
|`/`      | division         |
|`%`      | mod              |
|`<`      | less             |
|`>`      | greater          |
|`<=`     | less or equal    |
|`>=`     | greater or equal |
|`.`      | dot              |

**Keywords**

- `if` `else` `switch` `case`

- `loop` `skip` `break`

- `return`

- `not` `and` `or`

**Regex based tokens**

| Regex                   | Name       |
|-------------------------|------------|
| `letter alNum*`         | identifier |
| `sign? number`          | int        |
| `sign? number . number` | float      |
| `" notEscQuote* "`      | string     |
| `true`                  | true       |
| `false`                 | false      |

| Regex               | Subtoken    |
|---------------------|-------------|
| `[a-zA-Z]`          | letter      |
| `[0-9]`             | digit       |
| `[+-]`              | sign        |
| `letter \| digit`   | alNum       |
| `digit (_? digit)*` | number      |
| `\"`                | escQuote    |
| `[^\"]`             | notEscQuote |

## Grammar

## IR Code

## Bytecode
