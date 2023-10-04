FROM golang:1.21-alpine
#RUN apk add --no-cache go
COPY ./app/aqcronjob /go/src/app/aqcronjob
COPY ./common /go/src/common
COPY ./go.mod /go/src/
COPY ./go.sum /go/src/
WORKDIR /go/src/app/aqcronjob

ENV GO_MOD_SUCCESS=false
RUN until ${GO_MOD_SUCCESS}; do \
go mod tidy && export GO_MOD_SUCCESS=true; \
done

RUN go build aqcronjob.go
EXPOSE 3003
CMD [ "./aqcronjob" ]