#!/bin/sh

branchPath=$(git symbolic-ref -q HEAD) # branchPath は refs/heads/feature/XXXX_YYYY のような文字列に
branchName=${branchPath##*/} # 最後の / 以下を取得し、branchName は XXXX_YYYY のような文字列に
issueNumber=$(echo $branchName | cut -d "_" -f 1) # "_" を delimiter として cut し、issueNumber は XXXX に
firstLine=$(head -n1 $1)

%s

# This hook includes three examples.  The first comments out the
# "Conflicts:" part of a merge commit.
#
# The second includes the output of "git diff --name-status -r"
# into the message, just before the "git status" output.  It is
# commented because it doesn't cope with --amend or with squashed
# commits.
#
# The third example adds a Signed-off-by line to the message, that can
# still be edited.  This is rarely a good idea.

case "$2,$3" in
  merge,)
    /usr/bin/perl -i.bak -ne 's/^/# /, s/^# #/#/ if /^Conflicts/ .. /#/; print' "$1" ;;

# ,|template,)
#   /usr/bin/perl -i.bak -pe '
#      print "\n" . `git diff --cached --name-status -r`
#	 if /^#/ && $first++ == 0' "$1" ;;

  *) ;;
esac

