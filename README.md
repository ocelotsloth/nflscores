# NFLSCORES

This is a simple command-line utility which scrapes all of the NFL scores
and dumps them to the terminal in CSV format. This utility was written over
a weekend for a friend, and will likely not see active development again
unless somebody opens an issue.

## Installing

`go get git.sr.ht/~ocelotsloth/nflscores/nflscores` should get you where you
need to go.

You will need to install golang first, of course. At this time I will not
be distributing precompiled binaries.

## Usage

Simply type `nflscores` after installing the binary. To bring up the help
message.

## Architecture

Note that the library I defined is in the `git.sr.ht/~ocelotsloth/nflscores`
package, while the actual executable command is defined in the
`git.sr.ht/~ocelotsloth/nflscores/nflscores` package. You need to install the
executable to work with this tool. Look in that directory for the `main.go`
file if you are looking to read how this tool functions.
