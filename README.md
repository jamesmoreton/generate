# Generate

A random data generator CLI tool, written in Go.

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

## Setup

### With Docker

> Prerequisites:
> - Clone the repo
> - Install [Docker](https://docs.docker.com/get-docker/)

From the root directory `generate`, build the image:

```shell
$ docker build -t generate .
```

Run commands against the image:

```shell
$ docker run generate name
```

### With Go

> Prerequisites:
> - Clone the repo
> - Install [Go](https://go.dev/dl/) (version 1.18)

From the root directory `generate`, build with go:

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
