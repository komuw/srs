# srs            


[![ci](https://github.com/komuw/srs/workflows/srs%20ci/badge.svg)](https://github.com/komuw/srs/actions)
[![GoDoc](https://godoc.org/github.com/komuw/srs?status.svg)](https://godoc.org/github.com/komuw/srs)
[![Go Report Card](https://goreportcard.com/badge/github.com/komuw/srs)](https://goreportcard.com/report/github.com/komuw/srs)          


srs is a flashcard commandline app.     
This started off as a fork of [Leaf](https://github.com/ap4y/leaf).   
The main difference between this and `Leaf` is that in leaf, the cards are as `org-mode` files whereas this uses `markdown` files.    
`Leaf` is more developed and has better features, I would urge you to use it instead.


# Installing/Upgrading          
TODO:


# Usage  
`srs -d myCards/`    
               

#### debug
`go build -gcflags="all=-N -l" -o srs cmd/main.go` 
`dlv exec ./srs -- -d myCards/ -db myCards/srs.db`         
`(dlv) help`        
`(dlv) break cmd/main.go:23`        
`(dlv) continue`          
`(dlv) call getQuestion(mainNode)`


