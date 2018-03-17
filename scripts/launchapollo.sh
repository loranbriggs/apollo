#!/bin/bash

# disable DPMS (Energy Star) features.
xset -dpms

# disable screen saver
# xset s off

# don't blank the video device
xset s noblank

# disable mouse pointer
unclutter &

# run window manager
matchbox-window-manager -use_cursor no -use_titlebar no  &

# run browser
chromium-browser --kiosk --disable-session-crashed-bubble http://localhost:4444/
