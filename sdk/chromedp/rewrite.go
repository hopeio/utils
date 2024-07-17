package chromedp

import (
	"context"
	"encoding/base64"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/chromedp"
	"github.com/hopeio/utils/log"
	"os"
)

func Rewrite(ctx context.Context, url string, localPath string) chromedp.Action {
	data, err := os.ReadFile(localPath)
	log.Errorf("failed read file: %v", err)
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*fetch.EventRequestPaused); ok {
			go func() {
				err := chromedp.Run(ctx, fetch.FulfillRequest(ev.RequestID, 200).WithResponseHeaders([]*fetch.HeaderEntry{
					{Name: "Content-Type", Value: "application/javascript"},
				}).WithBody(base64.StdEncoding.EncodeToString(data)))
				if err != nil {
					log.Error("failed to run: %v", err)
				}
			}()
		}
	})
	return fetch.Enable().WithPatterns([]*fetch.RequestPattern{
		{URLPattern: url, RequestStage: fetch.RequestStageRequest},
	})
}
