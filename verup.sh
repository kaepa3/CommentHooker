#!/bin/bash

# get newest develop
git merge develop

verPath="./cmd/resource/version.txt"

# ver increment
awk '{print $1}' $verPath  | ( IFS=".$IFS" ; read a b c && echo $a.$b.$((c + 1)) >"$verPath" ) 
chmod 777 $verPath
Ver=$(<"$verPath")
echo $Ver

# update version
git add $verPath
git commit -m"Taging $Ver"
git push


# git Tagging
git tag $Ver
git push --tags

