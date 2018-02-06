[![CircleCI](https://circleci.com/gh/flowerinthenight/golang-monorepo.svg?style=svg)](https://circleci.com/gh/flowerinthenight/golang-monorepo)

## Overview

This is an example of a golang-based monorepo. It has the following features:

- Only build the services that are updated.
- Build all services that are affected by changes in common codes (i.e. `pkg`).
- Build all services that are affected by changes in `vendor` codes.

It uses [dep](https://github.com/golang/dep) as its dependency management tool.

For now, only [CircleCI](./.circleci/config.yml) is supported. But since it uses bash scripts and Makefiles, it should be fairly straightforward to port to [TravisCI](https://travis-ci.org/) or [AppVeyor](https://www.appveyor.com/).

## How does it work

During CircleCI builds, the [script](./.circleci/config.yml) iterates the updated files within the commit range (`CIRCLE_COMPARE_URL` environment variable in CircleCI) or the changed files within a single commit (when the value is not a valid range), excluding hidden files, `pkg` and `vendor` folders. It will then try to walk up the directory path until it can find a Makefile (excluding root Makefile). Once found, the root Makefile will include that Makefile and call the `custom` rule as target, thus, initiating the build.

When the changes belong to either `pkg` or `vendor`, the script will then try to determine the services (and cmds) that have dependencies using the `go list` command. All dependent services will then be built using the same process described above.
