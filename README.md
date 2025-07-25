# tblogs

A fast, modern, and hackable terminal blog reader written in Go. No external API dependencies—just a local config file and RSS/Atom feeds.

---

## Features

- Browse and follow a curated list of tech/dev blogs
- Read posts from any RSS or Atom feed
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
  - `Ctrl+P` — Saved Posts
  - `Ctrl+H` — Help
  - `Ctrl+F` — Search
  - `Ctrl+S` — Save/follow
  - `Ctrl+D` — Delete saved post
  - `Ctrl+L` — Toggle last login mode. When on, only posts published after the last login date are shown.
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
