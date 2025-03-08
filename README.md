Little toy project that I used as guinea pig to test context management and the semaphore pattern. Turns out that I'll really use this daily, so I'm leaving the repo public.


---

## How to Use?

First, generate a GitHub personal access token. It must have the `notifications` permission.

Then, build the project with `go build cmd/github-notifications/github-notifications.go`

Run it with `./github-notifications -token=PERSONAL_ACCESS_TOKEN`

You can also set up the app to run at startup. On Ubuntu, I did the following:

1. In your terminal, run crontab -e and choose your preferred text editor.

2. Add the line `@reboot /path/to/github-notifications -token=PERSONAL_ACCESS_TOKEN`. Yes, your token is in plaintext for now. I'll fix this later, maybe.

3. Enjoy!

---

There is some features that I might change later:

- [ ] Stop relying on `beeep`. I want my notifications to be more powerful, such as redirecting me on click and displaying more information.
- [ ] Add tests.
- [ ] Add more CLI options, such as poll frequency. Also, allow reading these options from environment variables.
- [ ] Write logs to a file.
- [ ] Store your precious GitHub token securely.
- [ ] Handle unread notifications more effectively by considering the latest time you read a notification. Currently, the app initially bombards you with unread notifications and then stores the latest read time. I'm not happy with this and will handle it properly later.
