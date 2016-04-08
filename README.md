# resume-generator

[![Circle CI](https://circleci.com/gh/guillaumebreton/regen/tree/master.svg?style=svg)](https://circleci.com/gh/guillaumebreton/regen/tree/master)

CLI to generate resume from toml files

# Installation

~~~
go get -u github.com/guillaumebreton/regen
~~~

# Template definition

See (examples)[https://github.com/guillaumebreton/regen/tree/master/examples] directory for template and toml definition

# Usage

~~~
regen generate  directory -t index.html -o out
~~~

#TODO

- [ ] Set default value for directory to "."
- [ ] Set default value for template to "./template.html"
- [ ] Complete data structure
