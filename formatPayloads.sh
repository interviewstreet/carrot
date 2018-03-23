#! /bin/bash

sed -i -e 's/^/{"body":/' payloads.txt
sed -i -e 's/$/}/' payloads.txt