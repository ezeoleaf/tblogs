# tblogs

A fast, modern, and hackable terminal blog reader written in Go. No external API dependencies—just a local config file and RSS/Atom feeds.

---

## Features

- Browse and follow a curated list of tech/dev blogs
- Read posts from any RSS or Atom feed
- View your X (Twitter) timeline
- Save your favorite posts
- Search and filter blogs
- All data stored locally in a config file (no external API)
- Cross-platform: macOS, Linux, Windows
- Keyboard-driven TUI (Terminal User Interface)

---

## Installation

### Homebrew (recommended, after release)

```sh
brew tap ezeoleaf/tap
brew install tblogs
```

### Download a Release

- Go to [Releases](https://github.com/ezeoleaf/tblogs/releases) and download the binary for your OS.
- Unpack and move it to a directory in your `$PATH` (e.g., `/usr/local/bin`).

### Build from Source

```sh
git clone https://github.com/ezeoleaf/tblogs.git
cd tblogs
make build
./bin/tblogs
```

---

## Usage

```sh
tblogs
```

- Use keyboard shortcuts to navigate:
  - `Ctrl+B` — Blogs
  - `Ctrl+T` — Home
  - `Ctrl+X` — X Timeline
  - `Ctrl+P` — Saved Posts
  - `Ctrl+H` — Help
  - `Ctrl+F` — Search
  - `Ctrl+S` — Save/follow
  - `Ctrl+D` — Delete saved post
  - `Ctrl+L` — Toggle last login mode. When on, only posts published after the last login date are shown.
  - `Ctrl+R` — Reload X timeline
- Select a blog to view its posts (fetched live from the feed)
- Press `Enter` on a post to open it in your browser

---

## Configuration

- Config is stored in the OS-appropriate location:
  - **macOS/Linux:** `~/.config/tblogs/data.yml`
  - **Windows:** `%APPDATA%\tblogs\data.yml`
- On first run, the config is created and pre-populated with a curated list of blogs.
- You can edit the config file directly to add/remove blogs, or use the app UI.
- Default blogs are defined in [`internal/config/default_blogs.yml`](internal/config/default_blogs.yml).

### X (Twitter) Integration

To enable X timeline viewing:

1. **Set up OAuth 2.0 App:**

   - Visit the [X Developer Portal](https://developer.twitter.com/)
   - Create a new app or use an existing one
   - Enable OAuth 2.0 with the following scopes:
     - `tweet.read` - Read tweets
     - `users.read` - Read user information
   - Set the redirect URI to `http://127.0.0.1:8080/callback`
   - **Important**: Make sure your app has OAuth 2.0 enabled (not just OAuth 1.0a)

2. **Get OAuth 2.0 Credentials:**

   - Note your **Client ID** and **Client Secret**
   - Use the included helper script: `./bin/xauth <client_id> <client_secret>`
   - Or use tools like [OAuth 2.0 Playground](https://developers.google.com/oauthplayground/)

3. **Configure your credentials:**
   Edit your config file (`~/.config/tblogs/data.yml`) and add:

   ```yaml
   app:
     x_cred:
       client_id: "your_client_id_here"
       client_secret: "your_client_secret_here"
       access_token: "your_access_token_here"
       refresh_token: "" # Can be empty, will be filled if provided
       username: "your_x_username"
   ```

4. **Use the X timeline:**
   - Press `Ctrl+X` to navigate to the X timeline
   - Press `Ctrl+R` to reload the timeline
   - Click on any post to open it in your browser

**Note:** The app will use your access token. If it expires, you'll need to re-run the OAuth flow to get a new one.

### Troubleshooting

**"Request was not matched" error:**

- Ensure your X app has OAuth 2.0 enabled (not just OAuth 1.0a)
- Make sure the redirect URI exactly matches: `http://localhost:8080/callback`
- Verify your app has the correct scopes: `tweet.read` and `users.read`

**"Invalid client" error:**

- Double-check your Client ID and Client Secret
- Ensure your app is approved and active in the X Developer Portal

**"Invalid redirect URI" error:**

- The redirect URI in your X app settings must exactly match: `http://127.0.0.1:8080/callback`

---

## Releases & Distribution

- Binaries are built automatically for each release (macOS, Linux, Windows).
- See the [Releases](https://github.com/ezeoleaf/tblogs/releases) page for downloads.

---

## Screenshots

<img width="2590" height="1746" alt="image" src="https://github.com/user-attachments/assets/f12d30d8-5491-497c-9792-3698b6a246a1" />

<img width="2700" height="1752" alt="image" src="https://github.com/user-attachments/assets/ffa12e7b-8506-4636-a7a2-48691e3ba219" />

<img width="2880" height="1750" alt="image" src="https://github.com/user-attachments/assets/71fb2253-a253-42aa-98b8-7f85d89206a9" />

---

## Contributing

Pull requests and issues are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md).

---

## License

[Apache License 2.0](LICENSE)
