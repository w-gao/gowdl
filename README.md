# gowdl

A Workflow Description Language (WDL) parser in Go.


## Getting started

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


## CLI commands

The `gowdl` CLI provides the following commands.

### `gowdl graph [FILE]`: get the dependency graph of WDL

Flags:
- `-r` `--recursive`: follow the imports


Sample input:

```wdl
# example.wdl
import "A.wdl"
import "B.wdl" as importB

workflow example {
    call A.taskA
    call importB.taskB
}
```

```wdl
# B.wdl
import "C.wdl"

task taskB {
    ...
}
```

```json
{
    "namespace": "",        // root
    "url": "example.wdl",
    "imports": [
    ]
}
```


## API

Additionally, `gowdl` also provides a few APIs that can be used in your own project.

### `parsers` package

