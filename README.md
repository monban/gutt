[![Go Report Card](https://goreportcard.com/badge/github.com/monban/gutt)](https://goreportcard.com/report/github.com/monban/gutt)
# Gutt
`gutt` is a text user interface mail user agent (TUI MUA) in the style of [Mutt](http://www.mutt.org/).
It is currently in (very) early development.

## Features
- MBOX

## Planned features
- MAILDIR
- IMAP
- Gmail
- Multiple mailboxes

## Usage
At current time, the only configurable option is what MBOX to read from.
 Gutt will first check for the `MAIL` environment variable, if it can't find it, it will try `/var/spool/mail/$USER`.
To look at a different MBOX, just specify the MAIL variable during invokcation:
```
MAIL=~/somembox gutt
```

