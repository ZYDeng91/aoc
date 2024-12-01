#!/bin/sh

if [[ -z $1 ]] ; then
	echo "what day u want"
	exit 1
fi

export $(cat .env)
AOC_URL=https://adventofcode.com

if [[ ! -d $1 ]]; then
	mkdir $1 
fi

cd $1

curl -s "$AOC_URL/$YEAR/day/$1"  -H "cookie: session=$SESSION;" | sed -n "/article/,/\/article/p" > $1.html
curl -so $1.in "$AOC_URL/$YEAR/day/$1/input" -H "cookie: session=$SESSION;"

cat $1.html | sed -n "/<pre><code>/,/<\/code><\/pre>/p" | sed -e "s/<\/*\(pre\|code\)>//g" > $1.ex
