package stealth

import (
	_ "embed"

	"github.com/playwright-community/playwright-go"
)

//go:embed "stealth.min.js"
var StealthJS string

// Inject adds a stealth JavaScript script to the given Playwright page to make it less detectable by automated detection mechanisms.
//
// The stealth script modifies various browser characteristics to evade detection by websites that block or track bots. The script is embedded in the Go binary using the //go:embed directive, which includes the contents of the "stealth.min.js" file as the StealthJS variable. Get the latest version of playwright-go-stealth for more up to version os the evasions.
//
// Usage example:
//
//	page, err := browser.NewPage()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	err = stealth.Inject(page)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// The Inject function uses the AddInitScript method to ensure the stealth script is executed before any other scripts on the page. This method is effective in bypassing many bot detection systems that run their detection logic early in the page lifecycle.
func Inject(page playwright.Page) error {
	return page.AddInitScript(playwright.Script{
		Content: playwright.String(StealthJS),
	})
}
