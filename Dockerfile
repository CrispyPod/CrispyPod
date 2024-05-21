FROM golang:1.21-alpine
WORKDIR /src
COPY backend/main.go backend/go.mod backend/go.sum  ./
COPY backend/migrations ./migrations
COPY backend/helpers ./helpers
COPY backend/rssFeed ./rssFeed
RUN go build -o /bin/crispypod

FROM node:18-alpine
WORKDIR /crispypod
COPY --from=0 /bin/crispypod ./
WORKDIR /src
COPY frontend frontend
ENV PUBLIC_PB_ENDPOINT="http://localhost:8080"
ENV SRC_FOLDER="/src"
ENV PUBLIC_BUILD_STATIC=0
ENV DISABLE_PB_WEBUI=1
ENV SRC_FOLDER="/src/frontend"
RUN cd frontend && npm install && npm run build
WORKDIR /crispypod
RUN cp -r /src/frontend/build-node /src/frontend/package.json /src/frontend/node_modules /src/frontend/start-server.js /crispypod
VOLUME [ "/crispypod/pb_data" ]
COPY start.sh start.sh
CMD ["sh","/crispypod/start.sh"]