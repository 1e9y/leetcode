#!/usr/bin/env fish

function clean_commit -d "Creates clean commit messages"
  set directory "."

  for d in (find -E $directory -type d -maxdepth 1 -regex ".*[0-9]{4}\.[a-z_]+.*")
    check_changes $d
  end
end

function check_changes
  set path $argv[1]

  # echo "checking $path"
  # if not git diff --quiet $path
  set value (git status --short --porcelain $path)
  if test -n "$value"
  # if test -n (git status --short $path | awk '{ print $1 }')
    set message (get_commit_message $path)
    echo $message
    pushd $path
    git add .
    git commit --message $message
    popd
  end
end  

function get_commit_message
  set dirname (basename $argv[1])

  set parts (string match --regex '([0-9]+)\.([a-z_]+)' $dirname)
  set number (string replace --regex '^0+' '' $parts[2])
  set name (string replace --all '_' ' ' $parts[3])
  set message "solution($number): $name"
  echo $message
end

clean_commit