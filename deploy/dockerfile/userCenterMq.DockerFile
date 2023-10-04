FROM golang:1.21-alpine
#RUN apk add --no-cache go
COPY . /go/src/
WORKDIR /go/src/app/usercenter/cmd/mq

ENV GO_MOD_SUCCESS=false
RUN until ${GO_MOD_SUCCESS}; do \
go mod tidy && export GO_MOD_SUCCESS=true; \
done

RUN go build usercenter.go
EXPOSE 3001
CMD [ "./usercenter" ]