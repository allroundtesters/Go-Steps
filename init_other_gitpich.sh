#!/usr/bin/env bash

BRANCH_NAME=$1
INDEX=$2
README_UPDATE=$3

if [ $README_UPDATE ]
then
    URL="https://gitpitch.com/allroundtesters/Go-Steps/$BRANCH_NAME"
    echo "- $INDEX.[$BRANCH_NAME](${URL})" >> README.md
    git add .
    git commit -m "add new notes to README"
    git push
fi



echo "start checkout a new branch ${BRANCH_NAME}"
git checkout -b $1

rm -rf steps
rm PITCHME.md
cp docs/$2-$1.md PITCHME.md

git add .
git commit -m "init $1 branch"
git push origin $1

