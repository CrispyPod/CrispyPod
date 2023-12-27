if [ $(basename $PWD) != "scripts" ]; then
    echo "please enter scripts folder to execute this script!"
    exit 1
fi

git submodule set-url frontend github:CrispyPod/frontend.git
cd ../frontend
git push
cd ..
git submodule set-url frontend https://github.com/CrispyPod/frontend.git
