#! /bin/sh

COMMENTS=$1

git add .
git commit -m "${COMMENTS}"
git push