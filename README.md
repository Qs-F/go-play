# Package go-play (Import name: play)

go-play provide cli version go play (play.golang.org).

This simply creates directory "localhost/tmp/RANDOM" in your GOPATH/src.  
And creates .describe file in the directory.

Random chars is now fixed at length of 5, but this will be changed if there were serious issues.

## Installation

`go get -u github.com/Qs-F/go-play/cmd/play`

(In the future, this might be changed to github.com/Qs-F/go-play)

## Future

### `play --serve --port PORT`

Like original play.golang.org, this will serve go-code and editor. (might use gotty)

### `play --version`

Now this is not so complicated that not need version number, but may have version.

### License

MIT

### Maintainer

@CreatorQsF, de-liKeR
