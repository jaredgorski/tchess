package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jaredgorski/tchess/internal/board"

	"github.com/gosuri/uilive"
)

var (
	clear		= "\033[H\033[2J"
	errorCode	= 1
	prompt		= "\n\n\tEnter command:\n\t⤑ "
	successCode	= 0
)

func cleanupAndExit(writer *uilive.Writer, code int) {
	writer.Stop()
	fmt.Print(clear)
	os.Exit(code)
}

func main() {
	// parse arguments
	ip := flag.String("ip", "", "If client, enter server IP to connect to")
	port := flag.String("port", "8282", "Enter port to connect over")
	flag.Parse()

	out := clear

	// set up reader and writer
	reader := bufio.NewReader(os.Stdin)
	writer := uilive.New()
	writer.Start()

	// set up sigterm handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanupAndExit(writer, errorCode)
	}()

	// set up server or client connection
	isClientMode := len(*ip) > 0

	var conn net.Conn

	if isClientMode {
		conn, _ = net.Dial("tcp", *ip + ":" + *port)
	} else {
		ln, _ := net.Listen("tcp", ":" + *port)

		out += fmt.Sprintf("\nServing on port %s\n", *port)
		fmt.Fprint(writer, out)

		conn, _ = ln.Accept()

		out = clear
		out += fmt.Sprintf("\nConnected on port %s\n", *port)
		fmt.Fprint(writer, out)
	}

	b := board.Board{
		IsWhiteSide: !isClientMode,
		IsLarge: true,
		UseIcons: true,
		LastSquare: 1,
		Writer: writer,
	}

	b.ResetBoard()
	myTurn := b.IsWhiteSide

	for {
		out = clear

		nearColor := "White"
		farColor := "Black"
		if !b.IsWhiteSide {
			nearColor = "Black"
			farColor = "White"
		}

		if myTurn {
			out += nearColor + " to play"
		} else {
			out += farColor + " to play"
		}

		out += b.DrawBoard()

		if myTurn {
			out += prompt
		}

		fmt.Fprint(writer, out)

		if myTurn {
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)

			if input == "exit" {
				break
			} else if input == "reset" {
				b.ResetBoard()
			} else {
				b.MovePiece(input)
			}

			fmt.Fprintf(conn, input + "\n")
		} else {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			message = strings.Replace(message, "\n", "", -1)

			b.MovePiece(message)
		}

		myTurn = !myTurn
	}

	cleanupAndExit(writer, successCode)
}
