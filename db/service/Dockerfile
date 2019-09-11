from golang

RUN apt-get update && apt-get install -y unzip git \
    && rm -rf /var/lib/apt/lists/*

RUN go get -d github.com/kolibrid/kvik/...

RUN ls $GOPATH

ADD . $GOPATH/src/github.com/kolibrid/db/service
WORKDIR $GOPATH/src/github.com/kolibrid/db/service
RUN go install

CMD PORT=80 service
