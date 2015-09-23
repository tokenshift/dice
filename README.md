# Dice

Dice roll simulator, as a library and a command-line tool.

## Tool

Usage:

	roll [-a|--avg|--average [-m|--min|--minimum] [-M|--max|--maximum] [-v|--verbose] [-r|--result|--random|--rand|--roll] [dice spec]

Where the *dice spec* is a string like "1d10+2d6+3". Without a dice spec, rolls
a single d6.

Options:

* `-a` `--avg` `--average` `--mean`
  Outputs the mean result for the dice set.
* `-m` `--min` `--minimum`
  Outputs the minimum result for the dice set.
* `-M` `--max` `--maximum`
  Outputs the maximum result for the dice set.
* `-v` `--verbose`
  Outputs a list of individual dice results. Dice are sorted in descending
  order of die size, with any constant/modifier at the end.
* `-r` `--result` `--random` `--rand` `--roll`
  Outputs a randomly rolled result. This is the default behavior.

If any options are provided, only the outputs requested by the options will be
provided. Multiple options can be provided, in which case the requested outputs
are displayed in the order in which the options were provided (delimited by tabs).

Each option counts as its own roll, so the same option can be provided multiple
times to get different results.

If multiple dice specs are provided, the output for each (with all options) will
be on its own line, in the order the specs were provided.