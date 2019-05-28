# codewalk
A first codewalk for a golang newbie

## Description
This is a codewalk into a golang codebase. To make it convenient it boots up from docker with all it needs to work.

## Command

### From source
To build the binary:
```sh
go build
```

To use the binary:
```sh
# to search for an rss entry with the word "president" in it. You can chose something else.
./rssfeed president
```
### Docker

To build it:
```sh
docker build -t codewalk:local .
```

To run it:

```sh
docker run --rm -it -p 80:6060 sofianiho/codewalk:latest
```
