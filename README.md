GXUI - A Go cross platform UI library.
=======

## My Fork

For the awareness of anyone eyeing my fork of gxui:

I originally forked this to solve some issues I was seeing with [the Go editor I'm writing](https://vidar).
I had planned to send pull requests for any bugs I managed to find and solve, but with it going unmaintained,
I've started just aggressively changing anything that doesn't work (or read) how I want it to.  Most of my
changes revolve around the CodeEditor and TextBox types.

A note about my use of gxui for the editor: when the abandonment announcement got posted to gxui's README,
I spent some time looking at other libraries.  There are plenty of others out there.  However, nothing I saw
looked like it would be a quick swap - gxui just feels easy to use by comparison.  That may be due to me
being used to gxui at this point, but that's beside the point.  Eventually, I decided I would rather
maintain a fork than try to port to another UI library.

If things continue to progress with my editor, there may be a day when I start maintaining my fork as a tested,
maintained fork from the original project; but for now, I'm just making fixes when I come across things my
editor needs, and usually making minor changes whenever I'm looking at code to make it easier for me to read.

I'm (for the moment) leaving this README as-is other than this minor blurb, so the gitter link will still be
pointing to upstream's gitter.im and whatnot.

[![Join the chat at https://gitter.im/google/gxui](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/google/gxui?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Build Status](https://travis-ci.org/google/gxui.svg?branch=master)](https://travis-ci.org/google/gxui) [![GoDoc](https://godoc.org/github.com/google/gxui?status.svg)](https://godoc.org/github.com/google/gxui)

Disclaimer
---
All code in this package **is experimental and will have frequent breaking
changes**. Please feel free to play, but please don't be upset when the API has significant reworkings.

The code is currently undocumented, and is certainly **not idiomatic Go**. It will be heavily refactored over the coming months.

This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.

Dependencies
---

### Linux:

In order to build GXUI on linux, you will need the following packages installed:

    sudo apt-get install libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev mesa-common-dev libgl1-mesa-dev libxxf86vm-dev

### Common:

After setting up ```GOPATH``` (see [Go documentation](https://golang.org/doc/code.html)), you can then fetch the GXUI library and its dependencies:

    go get -u github.com/google/gxui/...

Samples
---
Samples can be found in [`gxui/samples`](https://github.com/google/gxui/tree/master/samples).

To build all samples run:

    go install github.com/google/gxui/samples/...

And they will be built into ```GOPATH/bin```.

If you add ```GOPATH/bin``` to your PATH, you can simply type the name of a sample to run it. For example: ```image_viewer```.

Web
---

gxui code is cross platform and can be compiled using GopherJS to JavaScript, allowing it to run in browsers with WebGL support. To do so, you'll need the [GopherJS compiler](https://github.com/gopherjs/gopherjs) and some additional dependencies:

    go get -u github.com/gopherjs/gopherjs
    go get -u -d -tags=js github.com/google/gxui/...
    
Afterwards, you can try the samples by running `gopherjs serve` command and opening <http://localhost:8080/github.com/google/gxui/samples/> in a browser.

Fonts
---
Many of the samples require a font to render text. The dark theme (and currently the only theme) uses `Roboto`.
This is built into the gxfont package.

Make sure to mention this font in any notices file distributed with your application.

Contributing
---
GXUI was written by a couple of Googlers as an experiment, but with help of the open-source community GXUI could mature into something far more interesting.

Contributions, however small are extremely welcome but will require the author to have signed the [Google Individual Contributor License Agreement](https://developers.google.com/open-source/cla/individual?csw=1).

The CLA is necessary mainly because you own the copyright to your changes, even after your contribution becomes part of our codebase, so we need your permission to use and distribute your code. We also need to be sure of various other things—for instance that you'll tell us if you know that your code infringes on other people's patents. You don't have to sign the CLA until after you've submitted your code for review and a member has approved it, but you must do it before we can put your code into our codebase. Before you start working on a larger contribution, you should get in touch with us first through the issue tracker with your idea so that we can help out and possibly guide you. Coordinating up front makes it much easier to avoid frustration later on.
