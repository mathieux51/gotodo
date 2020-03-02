#!/usr/bin/env bash

set -o errexit
set -o nounset

# Install dependencies
sudo apt update
sudo apt install vim-gtk vimcurl git -y

# vim
curl https://raw.githubusercontent.com/mathieux51/vimrc/master/.vimrc > .vimrc
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
