package main

import (
	"bufio"
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
	arguments := os.Args
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

	// set up server or client
	PORT := ":" + arguments[1]
	ln, _ := net.Listen("tcp", PORT)

	out += fmt.Sprintf("\nServing on port %s\n", PORT)
	fmt.Fprint(writer, out)

	conn, _ := ln.Accept()

	out = clear
	out += fmt.Sprintf("\nConnected on port %s\n", PORT)
	fmt.Fprint(writer, out)

	b := board.Board{
		IsWhiteSide: true,
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

		if myTurn && b.IsWhiteSide {
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
		} else {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message Received:", string(message))
		}

		myTurn = !myTurn
	}

	cleanupAndExit(writer, successCode)
}
