package main

import (
	"context"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"main/handlers"
	"main/models"
	"main/ui"
)

const (
	address = "localhost:50051"
)

// "http://feeds.twit.tv/twit.xml"
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := models.NewRssClient(conn)

	_, err = client.Ping(context.Background(), &empty.Empty{})
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	logrus.Printf("Pinged. Ready to work...")


	if err := mainMenu(client); err != nil {
		logrus.Fatal(err)
	}
}

var menus = []ui.MenuElement{
	{
		Key:     '1',
		Header:  "Start server",
		Handler: handlers.Start,
	},
	{
		Key:     '2',
		Header:  "Stop server",
		Handler: handlers.Stop,
	},
	{
		Key:     '3',
		Header:  "Add rss link",
		Handler: handlers.AddRss,
	},
	{
		Key:     '4',
		Header:  "List news",
		Handler: handlers.ListNews,
	},
	{
		Key:     '5',
		Header:  "Get news",
		Handler: handlers.GetNews,
	},
	{
		Key:    'q',
		Header: "Exit",
	},
}

func mainMenu(client models.RssClient) error {
	ui.ClearScreen()

	var err error
	var ctx = context.Background()
	for err == nil {
		ui.PrintfNotice("%s\n", strings.Repeat("-", 10))
		for _, menu := range menus {
			ui.PrintfNotice("%c - %s\n", menu.Key, menu.Header)
		}
		var ch rune
		if ch, _, err = keyboard.GetSingleKey(); err != nil {
			break
		}
		ui.ClearScreen()
		menu, ok := ui.FindMenu(menus, ch)
		if ok {
			if menu.Handler == nil {
				ui.PrintfInfo("Exiting...\n")
				break
			}
			err = menu.Handler(ctx, client)
		} else {
			ui.PrintfError("Wrong key pressed - %c\n", ch)
		}
	}
	return err
}
