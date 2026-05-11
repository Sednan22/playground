package main

import (
	"log"
	"time"

	stealth "github.com/jonfriesen/playwright-go-stealth"
	"github.com/playwright-community/playwright-go"
)

const url = "https://abrahamjuliot.github.io/creepjs/"
const userAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0"

func FetchWebSite() error {

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}

	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		UserAgent: playwright.String(userAgent),
		ExtraHttpHeaders: map[string]string{
			"Cache-Control":      "no-cache",
			"Sec-Ch-Ua":          `"Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`,
			"Sec-Ch-Ua-Platform": "macOS",
		},
	})

	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	err = stealth.Inject(page)
	if err != nil {
		log.Fatalf("could not inject stealth script: %v", err)
	}

	_, err = page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	})
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	time.Sleep(2 * time.Second)

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("tests/v5.png"), // Change each test
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}

	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
	return nil
}
