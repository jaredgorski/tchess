# tchess
Chess with friends in the terminal [wip]

![](https://github.com/jaredgorski/tchess/raw/master/.media/tchess-screenshot.png)

The concept is simple: to facilitate a game of chess in the terminal, trading moves over a TCP connection.

If you want to try it out, build the binary (`go build`), open up a server on one machine (`./tchess -port 8888`), and then connect with a client from a different machine (`./tchess -ip <server ip> -port 8888`). Moves can be entered in a very basic/explicit algebraic notation, structured thusly: `<piece> + <current square> + <destination square>`. E.g. `Pc2c4` (move Pawn from C2 to C4)

Castling, en passant captures, and other complex moves aren't supported yet and move validation isn't implemented at all. Eventually I'll add these things, but at this point a full game can be played over the network with strictly one-piece moves.
