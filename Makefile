## update-stealth-js: Updates the stealth js script
.PHONY: update-stealth-js
update-stealth-js:
	@which npx > /dev/null 2>&1 || { echo >&2 "npx is not installed. Please install it first."; exit 1; }
	@npx extract-stealth-evasions

