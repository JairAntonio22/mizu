# Mizu programming language specification

## Purpose

Mizu to a multi-paradigm programming language for algorithmic problem solving.

## Tokens

**Separators**
| Literal | Name         |
|---------|--------------|
|`,`      | comma        |
|`:`      | colon        |
|`;`      | semicolon    |
|`_`      | placeholder  |
|`\`      | lambda       |
|`->`     | arrow        |
|`(` `)`  | parenthesis  |
|`{` `}`  | braces       |
|`[` `]`  | brackets     |

**Operators**
| Literal | Name             |
|---------|------------------|
|`::`     | define constant  |
|`:=`     | define variable  |
|`=`      | assign           |
|`+`      | add              |
|`-`      | sub              |
|`*`      | mul              |
|`/`      | div              |
|`%`      | mod              |
|`==`     | eq               |
|`!=`     | neq              |
|`<`      | less             |
|`>`      | greater          |
|`<=`     | less or equal    |
|`>=`     | greater or equal |
|`or`     | boolean or       |
|`and`    | boolean and      |
|`not`    | boolean not      |
|`.`      | dot              |

**Keywords**

- Control flow
    - Conditions
        - `if`
        - `else`
        - `switch`
        - `case`
    - Loops
        - `loop`
        - `skip`
        - `break`
    - Functions
        - `return`
        - `defer`
- Literals
    - `true`
    - `false`
- Memory allocation
    - `new`

**Regex based tokens**

| Regex               | Name    |
|---------------------|---------|
| `letter alNum*`     | id      |
| `digit (_? digit)*` | int     |
| `integer . integer` | float   |
| `" notEscQuote* "`  | string  |
| `# .* #`            | comment |
| `# .* \n`           | comment |

| Regex             | Subtoken     |
|-------------------|--------------|
| `[a-zA-Z]`        | letter       |
| `[0-9]`           | digit        |
| `letter \| digit` | alNum        |
| `\"`              | escQuote     |
| `[^\"]`           | notEscQuote  |

## Grammar

```
Program
    : Defintion+

Definition
    : id ('::' | ':=') (FunctionDefinition | TypeDefinition | Expression ';')

FunctionDefinition
    : FunctionArgumentList? (Type ReturnTypeList)? '{' StatementList '}'

FunctionArgumentList
    : '(' TypedArgumentList ')'

TypedArgumentsList
    : TypedArguments (';' TypedArgumentsList)?

TypedArguments
    : id (',' id)* Type

ReturnTypeList
    : '(' TypeList ')'

TypeList
    : Type (',' TypeList)?

```

## IR Code

## Bytecode
