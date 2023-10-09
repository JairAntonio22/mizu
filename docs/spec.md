# Mizu programming language specification

## Purpose

Mizu aims to be a simple programming language for algorithmic problem solving.

## Tokens

**Separators**
| Literal | Name         |
|---------|--------------|
|`,`      | comma        |
|`:`      | colon        |
|`;`      | semicolon    |
|`(` `)`  | parenthesis  |
|`{` `}`  | braces       |
|`[` `]`  | brackets     |

**Operators**
| Literal | Name             |
|---------|------------------|
|`::`     | define constant  |
|`:=`     | define variable  |
|`=`      | assign           |
|`+`      | addition         |
|`-`      | subtract         |
|`*`      | multiplication   |
|`/`      | division         |
|`%`      | mod              |
|`!=`     | mod              |
|`==`     | mod              |
|`<`      | less             |
|`>`      | greater          |
|`<=`     | less or equal    |
|`>=`     | greater or equal |
|`.`      | dot              |
|`?`      | question mark    |

**Keywords**

- Conditional branching
    - `if`
    - `else`
    - `switch`
    - `case`
- Looping control
    - `loop`
    - `skip`
    - `break`
- Function control
    - `return`
- Boolean logic
    - `not`
    - `and`
    - `or`
- Constants
    - `true`
    - `false`
    - `nil`

**Regex based tokens**

| Regex               | Name    |
|---------------------|---------|
| `letter alNum*`     | id      |
| `digit (_? digit)*` | integer |
| `integer . integer` | float   |
| `" notEscQuote* "`  | string  |

| Regex               | Subtoken    |
|---------------------|-------------|
| `[a-zA-Z]`          | letter      |
| `[0-9]`             | digit       |
| `letter \| digit`   | alNum       |
| `\"`                | escQuote    |
| `[^\"]`             | notEscQuote |

## Grammar

## IR Code

## Bytecode
