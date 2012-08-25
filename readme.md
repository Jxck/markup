# markup

Realtime Markdown Viewer with rendering by [Github Markdown Rendering API](http://developer.github.com/v3/markdown/)
API limit is 5000 render/h, it is enough for almost of all usecase.

## How to use

```shell
$ npm install markup -g
$ makrup [path/to/file] [port]
```

default port is 3000
open (http://localhost:port/) on your browser,
and write your markdown file with [gfm](http://github.github.com/github-flavored-markdown/)


## watch file or polling ?

markup reads the target file per 10 msec.
this looks very inefficient.
but in some distribution, fs.watchFile are too late from
time you saved your file.

so, I don't use fs.watchFile.

## Special Thanks

<3 Github


## licence

MIT
