# touchngo

This is a crude poller intended to trigger inotify change-events on one machine, where changes are being made from outside that machine.

My situation:

 * My editor is on Windows.
 * My files are stored on an SMB drive attached to a Linux machine.
 * The Linux machine is running the development servers, which have hot-reload capabilities that are based off inotify (eg React and Electron)

So:

 * If I develop on the Linux machine, the dev tooling works. Changes to the files trigger inotify events and the hot-reload hot reloads.
 * If I develop on the Windows machine, I need to either alt-tab to a terminal and `touch` the changed file or restart the dev tooling to see the changes.

Most people should never need this. _I_ shouldn't really need this. But here we are.

# Development / Building

 * Developed using Go 1.15+ and VS Code.
 * Compile with `go build`

# Usage

Build touchngo binary and locate it somewhere sensible. Run the binary from directory you want to watch. It will touch changed files every 3 seconds.

The following directories are hard-coded to be ignored:

 * node_modules
 * tmp
 * log
 * .git

# Credits

Concept inspired by https://github.com/clarabstract/watchntouch, name borrowed without permission from [Touch and Go](https://intervalsmusic.bandcamp.com/track/touch-and-go) by [Intervals](https://intervalsmusic.bandcamp.com), who also supplied the soundtrack to the writing of this code.

# License

BSD 2-clause simplified. See LICENSE for details.
