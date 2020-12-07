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

	mainMenu(client)
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
		Key: '4',
		Header: "Add url link",
		Handler: handlers.AddUrl,
	},
	{
		Key:     '5',
		Header:  "List news",
		Handler: handlers.ListNews,
	},
	{
		Key:     '6',
		Header:  "Get news",
		Handler: handlers.GetNews,
	},
	{
		Key:    'q',
		Header: "Exit",
	},
}

func mainMenu(client models.RssClient) {
	ui.ClearScreen()

	var ctx = context.Background()
	for {
		ui.PrintfNotice("%s", strings.Repeat("-", 50))
		for _, menu := range menus {
			ui.PrintfNotice("%c - %s", menu.Key, menu.Header)
		}
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			logrus.Error(err.Error())
			ui.PrintfError(err.Error())
			break
		}
		ui.ClearScreen()
		menu, ok := ui.FindMenu(menus, ch)
		if ok {
			if menu.Handler == nil {
				ui.PrintfInfo("Exiting...")
				break
			}

			if err := menu.Handler(ctx, client); err != nil {
				ui.PrintfError(err.Error())
			}
		} else {
			ui.PrintfError("Wrong key pressed - %c", ch)
		}
	}
}
