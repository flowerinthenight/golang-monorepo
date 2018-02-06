[![CircleCI](https://circleci.com/gh/flowerinthenight/golang-monorepo.svg?style=svg)](https://circleci.com/gh/flowerinthenight/golang-monorepo)

## Overview

This is an example of a golang-based monorepo. It has the following features:

- Only build the services that are updated.
- Build all services that are affected by changes in common codes (i.e. `pkg`).
- Build all services that are affected by changes in `vendor` codes.

It uses [dep](https://github.com/golang/dep) as its dependency management tool.
