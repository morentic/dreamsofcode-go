# Dreams of Code - Command Line Applications in Go

## Variables, Values and Types

- variable definition `var name = 'Marcel'` or short `name := 'Marcel'`
- array is a fixed size
- slices uses arrays under the hood and are dynamic
- map pass out value by a key instead of the index

## Loops

- the **underscore (\_)** is called the blank identifier, if two values are returned and you need only the second, the blank identifier is used to ignore the first value

## Runes

- a rune is simply an alias for the `int32` type — it represents a Unicode code point
- runes are declared with single quotes
- you can compare bytes against runes for example `if bytValue == ' '`
