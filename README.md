# Nap

Sleeping for a short time can be relaxing.

[License: MIT](https://git.sr.ht/~sthagen/nap/tree/default/item/LICENSE)

## Documentation

```console
❯ ./nap help
Nap implements real sleep (offering sub-second busy waits). The first real
value argument (base) in seconds determines the base duration. When given a real
number as second argument (variation) the duration will vary uniform randomly
plus minus the 2nd argument.

Default mode just sleeps for approx. one second. This behavior is identical to
when given a single argument of 1.

Usage:

    nap [base [variation]]

The flags are:

    -h
        print the usage information.

Examples:

❯ nap 0.75 0.1 # sleeps for a randomly selected duration in [0.65, 0.85] seconds

Caveat emptor: Maybe this command is not useful for your use cases.
```

## Building

```console
❯ go fmt && go vet && go build -ldflags "-s -w"
❯ go test -cover -coverprofile=coverage.out
PASS
coverage: 97.6% of statements
ok  	nap	0.429s
❯ go tool cover -func=coverage.out
nap/main.go:69:		ParseFloat	100.0%
nap/main.go:79:		ParseBase	100.0%
nap/main.go:90:		ParseVariation	100.0%
nap/main.go:100:	RandomDuration	100.0%
nap/main.go:107:	HelpRequested	100.0%
nap/main.go:118:	HandleAnyErrors	100.0%
nap/main.go:126:	Execute		100.0%
nap/main.go:153:	Seed		100.0%
nap/main.go:158:	main		0.0%
total:			(statements)	97.6%
❯ go tool cover -html=coverage.out -o coverage.html
❯ open coverage.html
```

## Bug Tracker

Any feature requests or bug reports shall go to the [todos of nap](https://todo.sr.ht/~sthagen/nap).

## Primary Source repository

The main source of `nap` is on a mountain in central Switzerland.
We use distributed version control (git).
There is no central hub.
Every clone can become a new source for the benefit of all.
The preferred public clones of `nap` are:

* [on codeberg](https://codeberg.org/sthagen/nap) - a democratic community-driven, non-profit software development platform operated by Codeberg e.V.
* [at sourcehut](https://git.sr.ht/~sthagen/nap) - a collection of tools useful for software development.

## Contributions

Please do not submit "pull requests" (I found no way to disable that "feature" on GitHub).
If you like to share small changes under the repositories license please kindly do so by sending a patchset.
You can either send such a patchset per email using [git send-email](https://git-send-email.io) or
if you are a sourcehut user by selecting "Prepare a patchset" on the summary page of your fork at [sourcehut](https://git.sr.ht/).

## Status

Prototyping.

**Note**: The default branch is `default`.
