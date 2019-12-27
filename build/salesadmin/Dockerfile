#-------------------------------------------------------------------------------------------------------------
# author: Alex Spina
# license: Have your way with me.
# date: 12/20/2019
#-------------------------------------------------------------------------------------------------------------
FROM golang:alpine3.10
#
## switch to root directory
WORKDIR /root
#
## copy project src files into container
COPY . ./project
#
## change into project directory
WORKDIR /root/project
#
## download dependencies from go.mod and go.sum
RUN go mod download
#
## build salesadmin executable and output to root
RUN go build -o /root/salesadmin cmd/salesadmin/salesadmin.go
#
## change into root directory
WORKDIR /root/
#
## move production config file to root
RUN mv ./project/configs/salesadmin/production/config.json .
#
## move necessary certs and keys to root
RUN mv ./project/test/certs/server.cer.pem .
RUN mv ./project/test/certs/server.key.pem .
RUN mv ./project/test/certs/salesadmin-ca-1.crt .
RUN mv ./project/test/certs/postgres.cer.pem .
RUN mv ./project/test/certs/postgres.key.pem .
#
## remove project source code
RUN rm -rf /root/project
#
## start z shell
ENTRYPOINT [ "./salesadmin", "start" ]