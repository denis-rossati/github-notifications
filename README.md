Little toy project that I used as guinea pig to test context management and the semaphore pattern. Turns out that I'll really use this daily, so I'm leaving the repo public.


---

# How to use?

First off, generate a GitHub personal access token. It must have the `notifications` permission.

Then, build with `go build cmd/github-notifications/github-notifications.go`

Run with `./github-notifications -token=PERSONAL_ACCESS_TOKEN`

You can also setup the app to run at initialization, in ubuntu I did the following:

1. In your terminal, run `crontab -e` and choose a text editor of your liking
2. Add `@reboot /path/to/github-notifications -token=YOUR_TOKEN`. Yes, your token in plaintext, I'll fix it later, maybe.
3. Be happy

---

There is some features that I might change later:

- [ ] Not rely on `beeep`, I want my notifications to be more powerful, like redirecting me on click and displaying more information.
- [ ] Add tests
- [ ] Add more CLI options, such as poll frequency. Also, be able to read them from env vars.
- [ ] Write logs to a file
- [ ] Store the your precious github token properly
- [ ] Handle unread notifications, considering the latest time you read a notification. Currently the app initially bombards you with unread notifications, then it stores the latest read time. I'm not happy with this, I'll handle this properly later.

