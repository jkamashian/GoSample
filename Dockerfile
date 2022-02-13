FROM golang:1.17-alpine 
# ENV GO111MODULE=on

# builds the Go executable
RUN mkdir /code
RUN mkdir /exe
COPY ./main /code
WORKDIR /code
RUN go get  && go build -o /exe/server .
RUN rm -rf /code

ENTRYPOINT ["/exe/server"]