go-sshproc
===

[![GoDoc](https://godoc.org/github.com/blacknon/go-sshproc?status.svg)](https://godoc.org/github.com/blacknon/go-sshproc)

## About

This library is adds the functionality to get performance information from procfs to [go-sshlib](https://github.com/blacknon/go-sshlib).
Use an sftp client to obtain performance information by retrieving information such as `/proc`.
For this reason, the ssh connection destination must be a UNIX-like OS such as Linux.

Some of the references and code have been changed in order to have almost the same usability as [goprocinfo](https://github.com/c9s/goprocinfo).
