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
    libssl-dev 

RUN echo 'deb http://cran.rstudio.com/bin/linux/debian buster-cran35/' >> /etc/apt/sources.list && \
    apt-key adv --keyserver keys.gnupg.net --recv-key 'E19F5F87128899B192B1A2C2AD5F960A256A04AF' && \ 
    apt-get update && \
    apt-get install -y \
        r-base

RUN R -e 'install.packages("jsonlite")'

RUN go get github.com/gorilla/mux 
RUN go get github.com/kolibrid/kvik/r

EXPOSE 80

WORKDIR /go/src/github.com/kolibrid/kvik/r/examples
CMD go run server.go -port=":80" 
