# Project Plan — Grab Careers Parser

Overview
- Two Go programs to parse job pages listed in `urls.txt`:
  - `cmd/js_api` — use the same JSON/XHR endpoints the site uses.
  - `cmd/headless` — use a headless browser to render JS and extract the DOM.

Goals
- Extract `title`, `workplace`, `job_type`, `location`, `description`, `page_url` for each vacancy.
- Produce a single CSV named `[YYYY-MM-DD].csv` with all results.
- Dockerize both programs.

Planned Steps
1. Create project skeleton
   - Add folders `cmd/js_api` and `cmd/headless`, core packages `internal/scraper`, `internal/io` and `go.mod`.
   - Add simple `main.go` runners that wire components.
2. Implement URL reader
   - Read `urls.txt`, validate URLs, return `[]string`.
   - Put implementation in `internal/io/urls.go`.
3. Implement JS-API scraper
   - Inspect page network calls, request JSON endpoints, map fields into Go struct.
4. Implement headless scraper
   - Use `chromedp` to load pages, wait for network idle or selector, extract rendered HTML and parse.
5. CSV writer & aggregator
   - Collect results, dedupe if needed, write to `[YYYY-MM-DD].csv`.
6. Dockerize programs
   - Add `Dockerfile` (multi-stage Go build) for each binary and optional `docker-compose.yml`.
7. Update README and examples
   - Document build/run steps and headless runtime requirements (Chrome/Chromium).
8. Local tests and formatting
   - `go build`, smoke tests on sample URLs, run `gofmt` and `go vet`.

Mentor mode
- I will act as your mentor: I will review design, suggest implementations, and provide code snippets when requested.
- I will NOT write or commit code unless you explicitly tell me to "implement" a specific step.

Next actions (pick one)
- I can scaffold the project skeleton (create folders and minimal files).
- I can write `internal/io/urls.go` to read and validate `urls.txt`.
- I can walk through how to inspect a page to find JS API endpoints (step-by-step).

Tell me which action to take and I’ll proceed.