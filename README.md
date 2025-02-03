# funmotd (Funny MOTD)

A simple and fun tool that displays random quotes from movies and TV shows as a [Message of the Day (MOTD)](https://en.wikipedia.org/wiki/Message_of_the_day) in your terminal when you open it.

> This is a re-implementation of [funmotd](https://github.com/veerendra2/funmotd-py), originally written in Python.

## Features

- Works entirely offline
- Single binary, no dependencies
- Option to enable NSFW quotes

## Installation

Download the latest binary from the [Releases](https://github.com/yourrepo/funmotd/releases) page and place it in a directory included in your `PATH`, such as `/usr/local/bin/`.

```bash
mv funmotd /usr/local/bin/
chmod +x /usr/local/bin/funmotd
```

## Usage

Run the following command to see available options:

```
./funmotd -h
Usage of ./funmotd:
A tool to display famous TV show and movie quotes.

Flags:
  -nsfw    Enable NSFW quotes
```

### Auto-run on Terminal Startup

To display a random quote each time you open your terminal, add the following line to your `~/.bashrc` (or `~/.bash_profile` for macOS users):

```bash
echo "/usr/local/bin/funmotd" >> ~/.bashrc
```

For Zsh users, add it to `~/.zshrc` instead:

```bash
echo "/usr/local/bin/funmotd" >> ~/.zshrc
```

### Enabling NSFW Quotes

If you want to include NSFW quotes, modify your shell configuration file:

```bash
echo "/usr/local/bin/funmotd -nsfw" >> ~/.bashrc
```

Then, reload your shell configuration:

```bash
source ~/.bashrc  # or source ~/.zshrc for Zsh users
```

## Contributing

Contributions are welcome! Feel free to submit issues, pull requests, or even suggest new quotes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
