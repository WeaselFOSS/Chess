Weasel
===========

![Logo](weasel.png)
Art graciously provided by [Alex](https://www.deviantart.com/redhead-alex)

This respository implements a basic chess engine.

It provides:

  * An efficient 0x88 and bitboard based chess engine
  * Simple but effective alpha beta pruning with null searches
  * Support for UCI, and a console mode
  * All chess rules including En passant, Fifty move rule, threefold repetition and castling

Table of Contents:

  * [About](#about)
  * [Installing and Compiling from Source](#installing-and-compiling-from-source)
  * [Contributing](#contributing)
  * [License](#license)

About
-----

Weasel is to my current knowledge the strongest chess engine written in Golang. There are still many improvements to be made currently we are working towards the first beta release.

Installing and Compiling from Source
------------

The easiest way to start using Weasel is to download it off the [releases](https://github.com/WeaselChess/Weasel/releases) page.


If you're looking to compile from source, you'll need the following:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install)
  * [GoVersionInfo](https://github.com/josephspurrier/goversioninfo/) used for Icon and version info on windows
  * A chess GUI that can communicate over UCI. We recommend [CuteChess](https://github.com/cutechess/cutechess/releases)

  After you have Go and a GUI you can install Weasel simply with
  ```
  go get github.com/WeaselChess/Weasel
  ```
  And from within the Go src folder simply run the following to compile
  ```
  go generate
  go build
  ```

Contributing
------------

Contributions are always welcome. If you're interested in contributing, send us an email or submit a PR.

License
-------

This project is currently licensed under GPLv3. This means you can use our source for your own project, so long as it remains open source.

Please refer to the [license](/LICENSE) file for more information.
