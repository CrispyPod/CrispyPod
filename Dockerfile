FROM golang:1.21-alpine
WORKDIR /src
COPY controllers ./controllers
COPY db ./db
COPY graph ./graph
COPY models ./models
COPY helpers ./helpers
COPY rssfeed ./rssfeed
COPY tools ./tools
COPY main.go go.mod go.sum gqlgen.yml ./
ENV GIN_MODE=release
RUN go mod tidy && go build -o /bin/crispypod

FROM node:18-alpine
COPY --from=0 /bin/crispypod /bin/crispypod
WORKDIR /crispypod
COPY *.js *.json ./
COPY src ./src
COPY static ./static
RUN npm install && npm run build
COPY start.sh start.sh
CMD sh /crispypod/start.sh

