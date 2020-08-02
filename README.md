
<div align="center">
  <h1>tchess</h1>
  <p>Chess with friends in the terminal [wip]</p>
</div>

<div align="center">
  <img src="https://github.com/jaredgorski/tchess/raw/master/.media/tchess-screenshot.png" width="500" />
</div>

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

The concept is simple: to facilitate a game of chess in the terminal, trading moves over a TCP connection.

If you want to try `tchess` out, build the binary (`go build`), open up a server in one terminal window (`./tchess -port 8888`), and then connect with the client from a different terminal window (`./tchess -ip <server ip> -port 8888`). Moves can be entered in a very basic/explicit algebraic notation, structured thusly: `<piece> + <current square> + <destination square>`. _E.g. `Pc2c4` (move Pawn from C2 to C4)_

Here's a quick step-by-step demo you can do yourself:
1. open a terminal window, clone this repository, and navigate to the root of the newly cloned `tchess` repository
2. build the `tchess` binary using the go compiler (`go build`)
3. run `tchess` in server mode (`./tchess -port 8888`)
4. open a new terminal window (separate instance) and navigate to the cloned `tchess` repository
5. run `tchess` in client mode, targeting the ip and port of the server `tchess` process in the other terminal window (`./tchess -ip 127.0.0.1 -port 8888`)
6. switch to the first terminal window and enter a move (try `Pd2d4`, a solid opening in the London System)
7. switch to the second terminal window, observe that `Pd2d4` has been communicated across the connection, and respond with the Kings Indian Defense (`Ng8f6`)
8. continue playing `tchess`!

<h2>Functionality</h1>

Castling, en passant captures, and other complex moves aren't supported yet and move validation is only implemented for non-pawn pieces. Eventually I'll add these things, but at this point a full, simple game can be played over the network with strictly one-piece moves.

For machine-to-machine connections, the machine running `tchess` in server mode will need to have a public IP address so that a client machine can connect to it. Normally, a central public server would be used to facilitate connections with all participating machines acting as clients. Seeing as I'm not currently planning on running a 24/7 server to support remote `tchess` games, public IPs and connecting over LAN will have to suffice for now.

Please let me know if you like it or have any questions, suggestions, or insights.
