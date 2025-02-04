#!/bin/bash

# Installation Script for funmotd
# This script downloads the latest version of funmotd and configures it to run on terminal startup.

REPO_URL="https://github.com/yourrepo/funmotd/releases/latest/download/funmotd"
INSTALL_DIR="/usr/local/bin"
BASHRC_FILE="$HOME/.bashrc"
BASH_PROFILE_FILE="$HOME/.bash_profile"
ZSHRC_FILE="$HOME/.zshrc"

echo "📥 Downloading funmotd..."
curl -L -o funmotd "$REPO_URL"

echo "📌 Installing funmotd to $INSTALL_DIR..."
mv funmotd "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/funmotd"

if [[ "$SHELL" == *"zsh"* ]]; then
    PROFILE_FILE="$ZSHRC_FILE"
elif [[ "$SHELL" == *"bash"* ]]; then
    if [[ -f "$BASHRC_FILE" ]]; then
        PROFILE_FILE="$BASHRC_FILE"
    else
        PROFILE_FILE="$BASH_PROFILE_FILE"
    fi
else
    echo "⚠️ Unsupported shell. Please add $INSTALL_DIR/funmotd to your shell config manually."
    exit 1
fi

if ! grep -q "$INSTALL_DIR/funmotd" "$PROFILE_FILE"; then
    echo "📝 Adding funmotd to $PROFILE_FILE..."
    echo "$INSTALL_DIR/funmotd" >> "$PROFILE_FILE"
fi

echo "🎉 Installation complete! Open a new terminal to see funmotd in action."
