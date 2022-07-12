# project

It's a tiny psql command but with less features. So why ? To be used in scratch docker container.

## Difference with psql

* By default, there is no header and no table view (like psql with options -t and -A)
* ...

# Dev


This project is using :

* golang 1.17+
* [task for development](https://taskfile.dev/#/)
* docker
* [docker buildx](https://github.com/docker/buildx)
* docker manifest
* [goreleaser](https://goreleaser.com/)

[For Linux, you can use this project to install the tools.](https://github.com/sgaunet/conf-linux)

## Build

```
task 
```
