# xgoinstall

`x-go-install`

## Overview

x-go-install = `xargs -I {} go install {}`.

## Limitation

Serial execution. Few options.

## Usage

```
cat tools.go | grep _ | awk -F'"' '{print $$2}' | x-go-install
```

## Motivation

We can use `xargs` only mac and unix.

## Install

### ghg

`ghg get sanemat/go-xgoinstall`

### go get

`go get https://github.com/sanemat/go-xgoinstall/cmd/x-go-install`

## Design

[design](./design.md)

## Changelog

[chagelog](./changelog.md) by [git-chglog](https://github.com/git-chglog/git-chglog)

## License

Copyright 2019 Matt (Sanemat) (Murahashi Kenichi)
[Apache License Version 2.0](./license.txt)

## Credits

[credits](./credits.txt) by [gocredits](https://github.com/Songmu/gocredits/)
