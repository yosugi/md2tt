# md2tt

simple markdown to textile converter

## Features

support bellow syntax

* heading
* bullet list
* numbered list

### heading

markdown

```
# header1
## header2
### header3
#### header4
##### header5
###### header6
```

textile

```
h1. header1
h2. header2
h3. header3
h4. header4
h5. header5
h6. header6
```

### bullet list

markdown

```
* one
* two
    * two-one
    * two-two
* three
```

textile

```
* one
* two
** two-one
** two-two
* three
```

### numbered list

markdown

```
1. one
2. two
    1. two-one
    2. two-two
3. three
```

textile

```
# one
# two
## two-one
## two-two
# three
```

## Usage

```
md2tt [indent]
```

example

```
$ echo "    * list" | md2tt
** list
$ echo "    * list" | md2tt 2
*** list
```

## Installation

```
$ go get github.com/yosugi/md2tt
```

## License

[MIT License](LICENSE)

## Version

0.1.0
