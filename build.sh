#!/bin/bash
echo Building domainfinder ...
go build -o domainfinder
echo Building synonyms ...
cd synonyms
go build -o ../dist/synonyms
echo Building available ...
cd ../available
go build -o ../dist/available
echo Building sprinkle ...
cd ../sprinkle
go build -o ../dist/sprinkle
echo Building coolify ...
cd ../coolify
go build -o ../dist/coolify
echo Building domainify ...
cd ../domainify
go build -o ../dist/domainify
cd ..
echo Done.
