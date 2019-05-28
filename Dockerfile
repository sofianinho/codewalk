FROM golang:1.11-stretch

COPY favicon.ico /go
COPY index.html /go
COPY background.gif /go
COPY codewalk.js /go/doc/codewalk/
COPY codewalk.css /go/doc/codewalk/
COPY popout.png /go/doc/codewalk/
ADD . /go/doc/codewalk/rssfeed

EXPOSE 6060
CMD ["godoc", "-http=:6060", "-goroot=/go/"]