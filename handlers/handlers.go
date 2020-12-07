package handlers

import (
	"context"
	"main/models"
	"main/ui"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
)

const (
	contextTimeout = time.Second * 5
)

func Start(ctx context.Context, client models.RssClient) error {
	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	_, err := client.Start(ctx, &empty.Empty{})
	if err == nil {
		ui.PrintfNotice("Server started")
	} else {
		ui.PrintfError(err.Error())
	}
	return nil
}

func Stop(ctx context.Context, client models.RssClient) error {
	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	_, err := client.Stop(ctx, &empty.Empty{})
	if err == nil {
		ui.PrintfNotice("Server stopped")
	} else {
		ui.PrintfError(err.Error())
	}
	return nil
}

func AddRss(ctx context.Context, client models.RssClient) error {
	link, err := ui.ReadLineOrAbort("Enter rss link: ", "abort")
	if err != nil {
		return err
	}
	if link == nil {
		return nil
	}

	seconds, err := ui.ReadInt64OrAbort("Enter the polling period of the source (in seconds)\n or leave it empty for default value (60 seconds)\n", ui.AbortKeyword)
	if err != nil {
		return err
	}
	if seconds == nil {
		return nil
	}
	if *seconds == 0 {
		*seconds = 60
	}

	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	_, err = client.AddRss(ctx, &models.RssLink{
		Url: *link,
		Duration: &timestamp.Timestamp{
			Seconds: *seconds,
		},
	})
	if err == nil {
		ui.PrintfNotice("Rss link '%s' added", *link)
	}
	return err
}

func AddUrl(ctx context.Context, client models.RssClient) error {
	link, err := ui.ReadLineOrAbort("Enter url link", ui.AbortKeyword)
	if err != nil || link == nil {
		return err
	}

	seconds, err := ui.ReadInt64OrAbort("Enter the polling period of the source (in seconds)\n or leave it empty for default value (60 seconds)\n", ui.AbortKeyword)
	if err != nil {
		return err
	}
	if seconds == nil {
		return nil
	}
	if *seconds == 0 {
		*seconds = 60
	}

	rule, err := ui.ReadLineOrAbort("Enter parsing rule (in css format)", ui.AbortKeyword)
	if err != nil || rule == nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	_, err = client.AddUrl(ctx, &models.UrlLink{
		Url: *link,
		Duration: &timestamp.Timestamp{
			Seconds: *seconds,
		},
		Rule: *rule,
	})
	if err == nil {
		ui.PrintfNotice("Rss link '%s' added", *link)
	}
	return err
}

func GetNews(ctx context.Context, client models.RssClient) error {
	request, err := ui.ReadLineOrAbort("Enter request ", "")
	if err != nil {
		return err
	}
	if request == nil {
		ui.PrintfInfo("Aborting...")
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	news, err := client.GetNews(ctx, &models.GetRequest{
		Request: *request,
	})
	if err != nil {
		return err
	}
	if news == nil || news.Articles == nil {
		ui.PrintfWarning("No news found for request string '%s'", *request)
	}

	for i := len(news.Articles) - 1; i >= 0; i-- {
		article := news.Articles[i]
		ui.PrintfWarning("%d - %s\n",
			i+1,
			article.Title,
		)
		ui.PrintfNotice("%s",
			article.Text,
		)
		ui.PrintfInfo("Link - %s",
			article.Url,
		)
		ui.PrintfNotice("%s",
			strings.Repeat("-", 50),
		)
	}

	return nil
}

func ListNews(ctx context.Context, client models.RssClient) error {
	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	news, err := client.ListNews(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	if news == nil || news.Articles == nil {
		ui.PrintfWarning("No news found.")
	}

	for i := len(news.Articles) - 1; i >= 0; i-- {
		article := news.Articles[i]
		ui.PrintfNotice("%s\n",
			strings.Repeat("-", 50),
		)
		ui.PrintfWarning("%d - %s\n\n",
			i+1,
			article.Title,
		)
		ui.PrintfInfo("Link - %s\n",
			article.Url,
		)
	}

	return nil
}
