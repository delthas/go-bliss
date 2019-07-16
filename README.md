# go-bliss [![builds.sr.ht status](https://builds.sr.ht/~delthas/go-bliss.svg)](https://builds.sr.ht/~delthas/go-bliss?) [![GoDoc](https://godoc.org/github.com/delthas/go-bliss?status.svg)](https://godoc.org/github.com/delthas/go-bliss)

Go bindings for [bliss](https://lelele.io/bliss.html) ([github](https://github.com/Polochon-street/bliss))

Quoted from its introduction:

Bliss is a library used to compute distance between two songs. It can be useful for creating « intelligent » playlists, for instance, and is used as such in leleleplayer and Blissify.

The main algorithm works by outputing a vector V = (a,b,c,d) for each song. The euclidean distance between these vectors corresponds to the actual distance felt by listening to them: a playlist can then be built by queuing close songs, or do other stuff using mathematical tools available on euclidean 4-dimensionnal spaces.

## status

- api stability: [![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental) open issues or PRs if you need anything new for your use case
- bugs: the bindings are small and partially tested

## using

- [install Bliss](https://lelele.io/bliss.html#download)
- import bliss "github.com/delthas/go-bliss"

## docs  [![GoDoc](https://godoc.org/github.com/delthas/go-bliss?status.svg)](https://godoc.org/github.com/delthas/go-bliss)

- [bliss introduction](https://lelele.io/bliss.html#whatis)
- [bliss technical details](https://lelele.io/bliss.html#details), defines the meaning of tempo/amplitude/frequency/attack ratings
- examples: [analysing a song](https://github.com/delthas/go-bliss/blob/master/analyze_example.go), [detecting the distance and cosine similarity between songs](https://github.com/delthas/go-bliss/blob/master/detect_example.go)
- [godoc](https://godoc.org/github.com/delthas/go-bliss)

## license

MIT
