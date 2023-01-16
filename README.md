# gowdl

A Workflow Description Language (WDL) parser in Go.

- **NOTE: this project is not under active development at the moment**
- `gowdl` provides the generated ANTLR4 parsers for WDL versions 1.0 and 1.1
- Go bindings for the WDL abtract syntax tree (AST) is work-in-progress
    - Currently, WDL imports can be parsed and used for projects like
      [wdl-viewer](https://github.com/w-gao/wdl-viewer)


## API

`gowdl` provides a few APIs that can be used in your own project.

### `parsers` package

To construct an ANTLR4 WDL parser, use:

```go
import "github.com/w-gao/gowdl/parsers"

func main() {
    // parser for version 1.0
    p := parsers.NewWdlV1_0Parser(data)

    // parser for version 1.1
    p := parsers.NewWdlV1_1Parser(data)

    // ...
}
```

## CLI commands

The `gowdl` CLI provides the following commands.

### `gowdl graph [FILE]`: get the dependency graph of WDL

**NOTE only imports are parsed - AST for the rest of the document is not all available.**

Sample input:

```wdl
# example.wdl
version 1.0

import "A.wdl"
import "B.wdl" as importB

workflow example {}
```

```wdl
# A.wdl
version 1.0

task taskA {
    command {}
}
```

```wdl
# B.wdl
version 1.0

import "C.wdl"

task taskB {
    command {}
}
```

```wdl
# C.wdl
version 1.0

task taskC {
    command {}
}
```

Sample output:

```json
[
    {
        "url": "example.wdl",
        "version": "1.0",
        "workflow": {
            "name": "example"
        },
        "imports": [
            {
                "url": "A.wdl",
                "absoluteUrl": "A.wdl"
            },
            {
                "url": "B.wdl",
                "absoluteUrl": "B.wdl",
                "as": "importB"
            }
        ]
    },
    {
        "url": "A.wdl",
        "version": "1.0"
    },
    {
        "url": "B.wdl",
        "version": "1.0",
        "imports": [
            {
                "url": "C.wdl",
                "absoluteUrl": "C.wdl"
            }
        ]
    },
    {
        "url": "C.wdl",
        "version": "1.0"
    }
]
```


## Getting started

If you want to contribute to the project or try it out locally, you can install
`gowdl` in the following ways:

### Install from source

To install `gowdl` from source, first clone this repo:

```
https://github.com/w-gao/gowdl.git
```

Then, generate the necessary parsers:

```
make generate
```

Finally, install the local package:

```
go install .
```

### Add as dependency via `go get`

This is not available yet.


## LICENSE

MIT License. Copyright (c) 2022-2023 William Gao
