FROM golang:1.12.9
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
    build-essential \
    wget \
    gcc \
    gfortran \
    libreadline-dev \
    perl \
    vim \
    libxml2-dev \ 
    xpdf \
    libcurl4-openssl-dev \
    libssl-dev \
    r-base \
    r-base-dev

RUN R -e 'install.packages("jsonlite")'

RUN go get github.com/gorilla/mux 
RUN go get github.com/fjukstad/kvik/r

EXPOSE 80

WORKDIR /go/src/github.com/fjukstad/kvik/r/examples
CMD go run server.go -port=":80" 
