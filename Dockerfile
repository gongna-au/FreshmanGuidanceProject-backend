FROM golang:1.14
WORKDIR $GOPATH/src/github.com/FreshmanGuidanceProject
COPY . $GOPATH/src/github.com/FreshmanGuidanceProject
RUN go env -w GOPROXY=https://goproxy.cn
RUN cd api/
RUN cd /database
RUN mysql -uroot -p123456
RUN source db.sql;
RUN \q
RUN cd ..
RUN cd ..
RUN go build main.go
EXPOSE 8000
ENTRYPOINT ["./main"]



