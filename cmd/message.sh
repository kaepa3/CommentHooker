#!/bin/bash

readonly COMMIT_MSG_FILE=$1
readonly COMMIT_SOURCE=$2

# ブランチ名取得
branch_name=$(git name-rev --name-only HEAD)
nsg="s/\[branch-name\]/$current_branch/"
echo $1 
echo $msg
# コミットメッセージファイル書込み
sed -i "" "s/\[branch-name\]/$current_branch/" $1

