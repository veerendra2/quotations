# Quotations

> _This is a re-implementation of [funmotd](https://github.com/veerendra2/funmotd-py), originally written in Python. 🐍_

A tool to display random inspirational quotes and famous dialogues from movies and TV shows on your CLI.

![Demo Gif](./assets/demo.mov)

## Features

- ✈️ Works entirely offline.
- 📦 Single binary, no dependencies.
- 🔞 Option to enable NSFW quotes.
- 🖥️ Useful as [Message of the Day (MOTD)](https://en.wikipedia.org/wiki/Message_of_the_day) in your terminal when you open it.

## Installation

### Homebrew

```bash
brew tap veerendra2/tap
brew install quotations
```

### Download Binaries

- Download the latest binary from the [Releases page](https://github.com/veerendra2/quotations/releases).
- Move it to a directory included in your `$PATH`, such as `/usr/local/bin/`.

```bash
mv quotations /usr/local/bin/
chmod +x /usr/local/bin/quotations
```

## Usage

Run the following command to see available options:

```bash
quotations --help
Usage of quotations:

  -entertainment
        Display entertainment quotes (default true)
  -inspirational
        Display inspirational quotes
  -nsfw
        Enable NSFW quotes
```

### Configuration

To display a random quote each time you open your terminal, add the following line to your shell configuration file (e.g., `~/.bashrc`, `~/.zshrc`, or `~/.profile`).

```bash
# Example for bash
echo "/usr/local/bin/quotations" >> ~/.bashrc

# To include NSFW quotes
echo "/usr/local/bin/quotations --nsfw" >> ~/.bashrc

# By default it displays entertainment quotes
# To display inspirational quotes
echo "/usr/local/bin/quotations --inspirational" >> ~/.bashrc
```

Then, reload your shell configuration:

```bash
source ~/.bashrc
```

## Add Quotes with Script

```bash
cd assets
pip3 install -r requirements.txt
# See comments in the script for help
python3 quotes_updater.py
```

## Contributing

Contributions are welcome! Feel free to submit issues, pull requests, or even suggest new quotes.
