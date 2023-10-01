# Mizu programming language specification

## Purpose

Mizu aims to be a simple programming language for algorithmic problem solving.

## Tokens

**Separators**
| Literal | Name         |
|---------|--------------|
|`:`      | colon        |
|`;`      | semicolon    |
|`(` `)`  | parenthesis  |
|`{` `}`  | curly braces |
|`[` `]`  | brackets     |
|`[` `]`  | brackets     |

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
- Constans
    - `true`
    - `false`
    - `nil`

**Regex based tokens**

| Regex                   | Name       |
|-------------------------|------------|
| `letter alNum*`         | identifier |
| `sign? number`          | int        |
| `sign? number . number` | float      |
| `" notEscQuote* "`      | string     |

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
