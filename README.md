[![CircleCI](https://circleci.com/gh/flowerinthenight/golang-monorepo.svg?style=svg)](https://circleci.com/gh/flowerinthenight/golang-monorepo)

## Overview

This is an example of a golang-based monorepo. It has the following features:

- Only build the services or cmds that are updated.
- Build all services and/or cmds that are affected by changes in common codes (i.e. `pkg`).
- Build all services and/or cmds that are affected by changes in `vendor` codes.

For now, only [CircleCI 2.0](./.circleci/config.yml) is supported. But since it uses bash scripts and Makefiles, it should be fairly straightforward to port to [TravisCI](https://travis-ci.org/) or [AppVeyor](https://www.appveyor.com/), etc.

At the moment, CI is setup to use Go1.14[.x] with `GO111MODULE=on` and `GOFLAGS=-mod=vendor` environment variables enabled during build. See sample [dockerfile](./services/samplesvc/dockerfile.samplesvc) for more details.

## How does it work

During CircleCI builds, the [script](./build.sh) iterates the updated files within the commit range (`CIRCLE_COMPARE_URL` environment variable in CircleCI) or the modified files within a single commit (when the value is not a valid range), excluding hidden files, `pkg`, and `vendor` folders. It will then try to walk up the directory path until it can find a Makefile (excluding root Makefile). Once found, the [root Makefile](./Makefile) will include that Makefile and call the `custom` rule as target, thus, initiating the build.

When the changes belong to either `pkg` or `vendor`, the script will then try to determine the services (and cmds) that have dependencies using the `go list` command. All dependent services will then be built using the same process described above.

## Directory structure

- `services/` - Basically, long running services.
- `cmd/` - CLI-based tools that are not long running.
- `pkg/` - Shared codes, or libraries common across the repo.
- `vendor/` - Third party codes from different vendors.

Although we have this structure, there is no limitation into where should you put your services/cmds. Any subdirectory structure is fine as long as a Makefile is provided.

## How to add a service/cmd

A reference template named [samplesvc](./services/samplesvc) is provided. Basically, these are the things that you need to do:

- Create a new directory for your service under `services/` or tool under `cmd/`. You may copy the [samplesvc](./services/samplesvc) contents to your new directory.
- Update the dockerfile inside your new service directory. Note that during build, this dockerfile is [copied](https://github.com/flowerinthenight/golang-monorepo/blob/master/services/samplesvc/Makefile#L21) to the root directory.
- Update the [Makefile](./services/samplesvc/Makefile) with your own values. You need to at least update the `MODULE` variable with your service name. The only required rule is the `custom` part so you may need to change that as well (i.e. name of the dockerfile used in `docker build`).
- [Optional] Update the [deploy.sh](./services/samplesvc/deploy.sh) script for your deployment needs.

## Misc
https://tech.mobingi.com/2018/09/25/ouchan-monorepo.html
