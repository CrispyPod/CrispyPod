echo "updating submodule..."
git submodule update --init --remote --recursive

printf 'Build frontend image (y/n)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ]; then
    echo "Building frontend..."
    cd ../frontend
    docker build . -t ghcr.io/crispypod/frontend:latest

    printf 'Push frontend image (y/n)? '
    read answer

    if [ "$answer" != "${answer#[Yy]}" ]; then
        echo "Pushing frontend..."
        docker push ghcr.io/crispypod/frontend:latest
    else
        echo No
    fi

else
    echo "No"
fi

printf 'Build backend image (y/n)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ]; then
    echo "Building backend..."
    cd ../backend
    docker build . -t ghcr.io/crispypod/backend:latest

    printf 'Push backend image (y/n)? '
    read answer

    if [ "$answer" != "${answer#[Yy]}" ]; then
        echo "Pushing backend..."
        docker push ghcr.io/crispypod/backend:latest
    else
        echo No
    fi

else
    echo "No"
fi

printf 'Build all in one image (y/n)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ]; then
    echo "Building all in one..."
    cd ../backend
    docker build . -t ghcr.io/crispypod/crispypod:latest

    printf 'Push all in one image (y/n)? '
    read answer

    if [ "$answer" != "${answer#[Yy]}" ]; then
        echo "Pushing all in one..."
        docker push ghcr.io/crispypod/crispypod:latest
    else
        echo No
    fi

else
    echo "No"
fi


