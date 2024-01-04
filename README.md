# Crispypod
A self-hosted podcast publishing platform, with the ability to deploy to serverless services and use your self-hosted instance as backend to serves images and audio files.

# Build docker image your self
## Build all in one image
```shell
git clone https://github.com/Crispypod/crispypod.git --recursive
cd crispypod
docker build . -t ghcr.io/crispypod/crispypod:latest
```
## Build separated images
```shell
git clone https://github.com/Crispypod/crispypod.git --recursive
cd crispypod/frontend
docker build . -t ghcr.io/crispypod/frontend:latest
cd ../backend
docker build . -t ghcr.io/crispypod/backend:latest
```

# Docker

## Dockerfiles
```
/Dockerfile # All in one image
/frontend/Dockerfile #  frontend image
/backend/Dockerfile # backend image
```

## Docker-compose files
```
/docker-compose/docker-compose-aio.yml # all in one compose files.
/docker-compose/docker-compose.yml # running frontend and backend in separate containers.
```