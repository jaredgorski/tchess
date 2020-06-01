# tchess
Chess with friends in the terminal [wip]

The concept is simple: to facilitate a game of chess in the terminal, trading moves over a TCP connection.

If you want to try it out, build the binary (`go build`), open up a server on one machine (`./tchess -port 8888`), and then connect with a client from a different machine (`./tchess -ip <server ip> -port 8888`).
