
<div align="center">
  <h1>tchess</h1>
  <p>Chess with friends in the terminal [wip]</p>
</div>

<div align="center">
  <img src="https://github.com/jaredgorski/tchess/raw/master/.media/tchess-screenshot.png" width="500" />
</div>

The concept is simple: to facilitate a game of chess in the terminal, trading moves over a TCP connection.

<h2>Usage</h1>
```
  -ip [server ip]
        If client, enter server IP to connect to
  -large
        Use large board
  -pieces [outline | filled | letter]
        Set pieces style to "outline", "filled", or "letter" (default "outline")
  -port [port number]
        Enter port to connect over (default "8282")
```

If you want to try `tchess` out, build the binary (`go build`), open up a server in one terminal window (`./tchess -port 8888`), and then connect with the client from a different terminal window (`./tchess -ip <server ip> -port 8888`). Moves can be entered in a very basic/explicit algebraic notation, structured thusly: `<piece> + <current square> + <destination square>`. _E.g. `Pc2c4` (move Pawn from C2 to C4)_

Castling, en passant captures, and other complex moves aren't supported yet and move validation is only implemented for non-pawn pieces. Eventually I'll add these things, but at this point a full, simple game can be played over the network with strictly one-piece moves.

For machine-to-machine connections, the machine running `tchess` in server mode will need to have a public IP address so that a client machine can connect to it. Normally, a central public server would be used to facilitate connections, with all participating machines acting as clients. Seeing as I'm not currently planning on running a 24/7 server to support remote `tchess` games, public IPs and connecting over LAN will have to suffice for now.

Please let me know if you like it or have any questions, suggestions, or insights.
