package main

import (
	"log"

	stealth "github.com/jonfriesen/playwright-go-stealth"
	"github.com/playwright-community/playwright-go"
)

const targetSite = "https://bot.sannysoft.com"
const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"

func main() {
	// Initialize Playwright
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	// Launch a browser
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	// Create a new page without stealth script
	pageNoStealth, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Navigate to the URL without stealth script
	_, err = pageNoStealth.Goto(targetSite, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	if err != nil {
		log.Fatalf("could not go to URL: %v", err)
	}

	// Take a screenshot without stealth script
	_, err = pageNoStealth.Screenshot(playwright.PageScreenshotOptions{
		FullPage: playwright.Bool(true),
		Path:     playwright.String("without_stealth.png"),
	})
	if err != nil {
		log.Fatalf("could not take screenshot: %v", err)
	}
	pageNoStealth.Close()

	// Create a new page with stealth script
	pageWithStealth, err := browser.NewPage(playwright.BrowserNewPageOptions{
		UserAgent: playwright.String(userAgent),
	})
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Inject stealth script
	err = stealth.Inject(pageWithStealth)
	if err != nil {
		log.Fatalf("could not inject stealth script: %v", err)
	}

	// Navigate to the URL with stealth script
	_, err = pageWithStealth.Goto(targetSite, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	if err != nil {
		log.Fatalf("could not go to URL: %v", err)
	}

	// Take a screenshot with stealth script
	_, err = pageWithStealth.Screenshot(playwright.PageScreenshotOptions{
		FullPage: playwright.Bool(true),
		Path:     playwright.String("with_stealth.png"),
	})
	if err != nil {
		log.Fatalf("could not take screenshot: %v", err)
	}
	pageWithStealth.Close()

	log.Println("Screenshots taken successfully")
}
