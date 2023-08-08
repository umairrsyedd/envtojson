# envtojson: Convert .env Files to JSON

A simple command-line tool to convert `.env` files to JSON format without reading the contents of the `.env` file. This tool makes it easy to convert environment variable files to JSON representation, allowing for easy parsing and integration into different systems.

## Features

- Convert `.env` files to JSON format
- Specify input `.env` file and optional output JSON file
- Easy installation with `go get`
- Cross-platform support

🔐 **Security Note:** Rest assured that we do not read the contents of your `.env` file during the conversion process. Your sensitive environment variable data remains private.

## Get Started

```sh
go get -u github.com/umairrsyedd/envtojson
```

## Usage

Convert a `.env` file to JSON:

```sh
envtojson -i input.env -o output.json
```
