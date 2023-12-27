if [ $(basename $PWD) != "scripts" ]; then
    echo "please enter scripts folder to execute this script!"
    exit 1
fi

git submodule set-url backend github:CrispyPod/backend.git
cd ../backend
git push
cd ..
git submodule set-url backend https://github.com/CrispyPod/backend.git
