# docker image build -t mom0tomo/go_sample_repo:latest .
# docker container run -it -p 8989:8080 mom0tomo/go_sample_repo:latest
# curl -X GET http://localhost:8989
FROM golang:1.9

RUN mkdir /hello

COPY main.go /hello

CMD [ "go","run","/hello/main.go" ]
