# Go File Uploader
A simple project leveraging the multipart interface for transferring files via HTTP

## Overview

The project provides:

- A web server that receives multipart requests and saves attached files on behalf of the server
- A CLI client that sends attached files to the server

## Building

Build the web server and the CLI with

```
make build
```

## Using the web server

The web server can accept multiple files in one request

By default files are sent to the `storage` dir. You can change it by providing `-s` or `--storage` flag
```
serveme -s=path-to-storage
```
Also you can change the port of the web server by providing `-p` or `--port` flag. The default value is `8080`
```
serveme -p=69420
```
> [!NOTE]
> * If you go to your browser and open `/files` page, it will list you all the files that have been uploaded to the server
> * Also `/` page has a UI form for uploading files

## Using the CLI

The CLI sends a request with multiple files
> [!IMPORTANT]
> If at least one file does not exist, then the whole request will blow up

To upload files you should list their paths as arguments to the CLI
```
sendme file1.jpg file2.pdf file3
```

By default the CLI sends requests on `localhost:8080` host. You can change it with `--to` flag
```
sendme -to=https://someurl.com file1 file2
```
