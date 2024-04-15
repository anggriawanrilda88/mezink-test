FROM golang:latest

# install Make
RUN apt-get update && apt-get install -y make

# container work directory
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# copy all code to container
COPY . .

# install setup tools for migrate
RUN make setup-tools

# build golang binary executable file
RUN make build

# expose port 
EXPOSE 8080

# default command run with container
CMD ["make", "migration-up", "run"]