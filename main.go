package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"main/handlers"
	"main/models"
	"main/ui"
)

var (
	port int64 = 50051
)

func init() {
	if len(os.Args) > 1 {
		if err := parseArguments(os.Args[1:]); err != nil {
			logrus.Fatalf("Can't parse arguments (cause: %s)", err.Error())
		}
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure(), grpc.WithBlock())
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
		Key:     '4',
		Header:  "Add url link",
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

// Format "--<key>=<value>"
func parseArguments(args []string) error {
	r := regexp.MustCompile(`^--([a-z-]+)=([a-z0-9\.]+)$`)
	for _, arg := range args {
		var err error
		values := r.FindAllStringSubmatch(arg, 1)
		if len(values[0]) != 1 && len(values[0]) != 3 {
			return fmt.Errorf("wrong argument %s", arg)
		}
		switch values[0][1] {
		case "port":
			port, err = strconv.ParseInt(values[0][2], 10, 64)
			if err != nil {
				break
			}
			if port < 0 || port > 65536 {
				err = fmt.Errorf("wrong port defined %d", port)
			}
		default:
			err = fmt.Errorf("unknown argument %s", arg)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
