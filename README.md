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

Use JSON file in the following format:

```json
[
    {
        "hotkey": ["ctrl", "alt", "del"],
        "many_times": 1,
        "interval_between_hotkeys": 0,
        "sleep_time": 666
    },
    {
        "hotkey": ["alt", "f4"],
        "many_times": 1337,
        "interval_between_hotkeys": 1,
        "sleep_time": 10
    },
]
```