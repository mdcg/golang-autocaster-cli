# golang-autocaster-cli

## Dependencies

For Ubuntu:

```bash
sudo apt install gcc libc6-dev

sudo apt install libx11-dev xorg-dev libxtst-dev libpng++-dev

sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt install libxkbcommon-dev

sudo apt install xsel xclip
```

For Mac OS X:

```
Xcode Command Line Tools
```
For Windows:

```
MinGW-w64 (Use recommended) or other GCC
```

## Generating the executable

```
go build -o autocaster .
```

## Config autocaster file

HOTKEY (splitted by space);SLEEP TIME;HOW MANY TIMES