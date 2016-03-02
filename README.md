# onetwoseven

The simplest programmer's web server. Use it as an executable or a Go library.

    $ 127
    12:43 I ━┃ ━━┃┏━┃
    12:43 I  ┃ ┏━┛  ┃
    12:43 I ━━┛━━┛  ┛
    12:43 I starting . on 127.0.0.1:8190

There are as few options as possible.

    $ 127 --help
    usage: 127 [<flags>] [<path>]

    Programmer's test http server

    Flags:
          --help                   Show context-sensitive help (also try --help-long and --help-man).
      -h, --host="127.0.0.1:8190"  Host and port to start server on
      -p, --path="."               path to serve.
      -d, --verbose                Print debugging output.

