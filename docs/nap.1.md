---
title: NAP
section: 1
header: User Manual
footer: nap 2023.9.19
date: September 19, 2023
---

# NAME

nap - Sleeping for a short time can be relaxing.

# SYNOPSIS

**nap** [*base* [*variation*]]

# DESCRIPTION

**Nap** implements real sleep (offering sub-second busy waits).

The first real value argument (base) in seconds determines the base duration. When given a real number as second argument (variation) the duration will vary uniform randomly plus minus the 2nd argument.

Default mode just sleeps for a second. This behavior is identical to when given a single argument of 1.


# OPTIONS

**-h** 
:    display help message

**-v** 
:    display the version string


# EXAMPLES

**nap 0.75 0.1**
:    Sleeps for a randomly selected duration in \[0.65, 0.85] seconds and then exits.


# NOTES

Caveat emptor: Maybe this command is not useful for your use cases.


# AUTHORS

Written by Stefan Hagen.

# BUGS

Submit bug reports online at: <https://todo.sr.ht/~sthagen/nap>

# SEE ALSO

Full documentation and sources at: <https://git.sr.ht/~sthagen/nap>
