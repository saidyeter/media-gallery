# media-gallery
A media gallery app  

Server is written by Golang. An http server.

## Server Endpoints

- `/` : Home endpoint. Returns other endpoint details.
- `/dirs` : Directories endpoint. Returns available root directories that specified in `vars.json`. Each root directory returns with numeric id. That id can be used in `/files/{id}` as id to retrieve contents in the directory.
- `/files/{id}` : Files endpoint. Returns directory paths and image file paths, also thumbnail image content as base64. That file path can be used in `/file/{path}` as path to retrieve actual image content. This endpoint also has paging functionality. To do that, start index(`s`) and end index(`e`) need to be specified as query parameters. Eg. `http://localhost:8080/files/3?s=3&&e=5`. Otherwise, the endpoint will return from 0 (zero) to limit.
- `/file/{path}` : File endpoint. Returns actual image content as base64.

Client is written  by HTML Css and JavaScript.

## Todo
- Create docker compose file
- Create web client

## Requirements

 - [Docker desktop (for Docker Compose)](https://www.docker.com/products/docker-desktop)
 - [Go (to run locally)](https://go.dev/dl) 

## Repository Usage

 - Clone this repository to local 
 - Use `docker-compose up` 

## Authors

 - [Said Yeter](https://github.com/kordiseps)

## Licence

MIT License

Copyright (c) 2022 Said Yeter

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.