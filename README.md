# Nap

Sleeping for a short time can be relaxing.

[License: MIT](https://git.sr.ht/~sthagen/nap/tree/default/item/LICENSE)

## Documentation

You can create a man page per:

```console
❯ make man
... man share/man/man1/nap.1
```

### Synopsis

```console
nap [base [variation]]
```

### Help

```console
❯ ./nap help
Nap implements real sleep (offering sub-second busy waits). The first real
value argument (base) in seconds determines the base duration. When given a real
number as second argument (variation) the duration will vary uniform randomly
plus minus the 2nd argument.

Default mode just sleeps for a second. This behavior is identical to
when given a single argument of 1.

Usage:

    nap [base [variation]]

The flags are:

    -h
        display help message

    -v
        display the version string

Examples:

❯ nap 0.75 0.1 # sleeps for a randomly selected duration in [0.65, 0.85] seconds

Caveat emptor: Maybe this command is not useful for your use cases.
```

### Version

```console
> ./nap version
v2023.9.20
```


### Building

```console
❯ make
... building for linux/amd64, local, and windows/amd64
... smoke test of local app
v2023.9.20
Nap implements real sleep (offering sub-second busy waits). The first real
value argument (base) in seconds determines the base duration. When given a real
number as second argument (variation) the duration will vary uniform randomly
plus minus the 2nd argument.

Default mode just sleeps for a second. This behavior is identical to
when given a single argument of 1.

Usage:

    nap [base [variation]]

The flags are:

    -h
        display help message

    -v
        display the version string

Examples:

❯ nap 0.75 0.1 # sleeps for a randomly selected duration in [0.65, 0.85] seconds

Caveat emptor: Maybe this command is not useful for your use cases.
... man share/man/man1/nap.1
PASS
coverage: 98.0% of statements
ok    nap 0.304s
nap/main.go:77:   ParseFloat    100.0%
nap/main.go:87:   ParseBase   100.0%
nap/main.go:98:   ParseVariation    100.0%
nap/main.go:108:  RandomDuration    100.0%
nap/main.go:115:  HelpRequested   100.0%
nap/main.go:126:  VersionRequested  100.0%
nap/main.go:137:  HandleAnyErrors   100.0%
nap/main.go:145:  Execute     100.0%
nap/main.go:176:  Seed      100.0%
nap/main.go:181:  main      0.0%
total:      (statements)    98.0%
... open coverage.html
```

### Explorative Testing

Benchmarking the default scenario with [hyperfine](https://crates.io/crates/hyperfine) against some sleep binary a random machine provides:

```console
❯ hyperfine --warmup 1 'sleep 1' './nap 1'
Benchmark 1: sleep 1
  Time (mean ± σ):      1.014 s ±  0.004 s    [User: 0.002 s, System: 0.004 s]
  Range (min … max):    1.009 s …  1.020 s    10 runs

Benchmark 2: ./nap 1
  Time (mean ± σ):      1.014 s ±  0.002 s    [User: 0.003 s, System: 0.003 s]
  Range (min … max):    1.010 s …  1.017 s    10 runs

Summary
  sleep 1 ran
    1.00 ± 0.00 times faster than ./nap-local 1
```

Testing another scenario requesting naps between 0.5 and 1.0 seconds:

```console
❯ hyperfine './nap .75 .25'
Benchmark 1: ./nap .75 .25
  Time (mean ± σ):     748.5 ms ± 153.2 ms    [User: 3.0 ms, System: 3.6 ms]
  Range (min … max):   528.4 ms … 979.4 ms    10 runs
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
