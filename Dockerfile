FROM golang:1.21-alpine
WORKDIR /src
COPY backend backend
WORKDIR /src/backend
RUN go build -o /bin/crispypod

FROM node:18-alpine
COPY --from=0 /bin/crispypod /bin/crispypod
WORKDIR /src
COPY frontend frontend
ENV BACK_END_URL="http://localhost:8080"
ENV SRC_FOLDER="/src/frontend"
RUN cd frontend && npm install && npm run build
WORKDIR /crispypod
RUN cp -r /src/frontend/build-node /src/frontend/package.json /src/frontend/node_modules /src/frontend/start-server.js /crispypod
VOLUME [ "/crispypod/UploadFile" ]
COPY start.sh start.sh
CMD ["sh","/crispypod/start.sh"]