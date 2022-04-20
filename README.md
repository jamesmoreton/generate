# Generate

A random data generator CLI tool, written in Go.

[![CI Master Test Status](https://github.com/jamesmoreton/generate/actions/workflows/ci-master-test.yml/badge.svg)](https://github.com/jamesmoreton/generate/actions/workflows/ci-master-test.yml)
[![CI Push to Docker Status](https://github.com/jamesmoreton/generate/actions/workflows/ci-docker.yml/badge.svg)](https://github.com/jamesmoreton/generate/actions/workflows/ci-docker.yml)

## Features

- **Name**  
  Person name: first, last, and sometimes middle.
- **Email**  
  Email address with variety of domains.
- **Mobile**  
  Mobile phone number (UK format), sometimes with area code.
- **Postcode**  
  Postcode (UK).
- **National Insurance number (UK)**  
  National Insurance number (UK).

![Generate GIF][generate]

[generate]: resources/generate.gif "Generate GIF"

## Setup

### With Docker

> N.B. [Docker](https://docs.docker.com/get-docker/) must be pre-installed.

The Docker image is hosted on [Docker Hub](https://hub.docker.com/r/moretonj/generate). Fetch the image:

```shell
$ docker pull moretonj/generate
```

Run commands against the image:

```shell
$ docker run moretonj/generate name
```

Or clone the repo and build the image directly from the root directory `generate`:

```shell
$ docker build -t generate .
```

Run commands against the image:

```shell
$ docker run generate name
```

### With Go

> N.B. [Go](https://go.dev/dl/) (version 1.18) must be pre-installed.

Clone the repo and build with Go from the root directory `generate`:

```shell
$ go build
```

Run the executable:

```shell
./generate name
```

## Usage

You can run `generate help` once installed to get help in the terminal, however the general usage is:

```shell
generate <command> [<args>]
```

Args provide additional functionality for the subcommand; run `generate <command> -help` to see the usage of a specific subcommand.
