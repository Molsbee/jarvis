package service

import (
	"context"
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func GetCheckMKStatus() ([]string, error) {
	checkUser := config.UserConfig.Domain.Username
	checkPass := config.UserConfig.Domain.Password
	if checkUser == "" || checkPass == "" {
		return nil, fmt.Errorf("no credentials provided for accessing check_mk")
	}

	allocatorContext, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.NoSandbox, chromedp.Headless)
	ctx, cancel := chromedp.NewContext(allocatorContext, chromedp.WithLogf(log.Printf))
	defer cancel()

	var text string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://master.cmk.ctl.io/master/check_mk/login.py"),
		chromedp.WaitVisible("#input_user"),
		chromedp.SendKeys("#input_user", checkUser),
		chromedp.SendKeys("#input_pass", checkPass),
		chromedp.Submit("#_login"),
		chromedp.Sleep(time.Second*10),
		chromedp.Navigate("https://master.cmk.ctl.io/master/check_mk/dashboard_dashlet.py?id=3&mtime=0&name=main"),
		chromedp.OuterHTML("#data_container > table", &text, chromedp.ByID),
	); err != nil {
		return nil, err
	}

	return parseTable(text)
}

func parseTable(table string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(table))
	if err != nil {
		return nil, err
	}

	var alerts []string
	doc.Find("table").Each(func(tableIndex int, tableHtml *goquery.Selection) {
		doc.Find("tbody").Each(func(bodyIndex int, bodyHtml *goquery.Selection) {
			bodyHtml.Find("tr").Each(func(trIndex int, tableRow *goquery.Selection) {
				var line string
				tableRow.Find("td").Each(func(tdIndex int, td *goquery.Selection) {
					line += td.Text() + " "
				})

				if len(line) != 0 {
					alerts = append(alerts, line)
				}
			})
		})
	})

	return alerts, nil
}
