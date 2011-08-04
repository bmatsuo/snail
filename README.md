*snail version 0.0_3*

About snail
=============

Command snail prints a "snail matrix" to its standard output stream.
Inspired by the [blog post by Allister Sanchez]{http://hackgolang.blogspot.com/2010/09/snail-in-golang.html}

Dependencies
=============

You must have Go installed (http://golang.org/). 

Documentation
=============
Usage
-----

Run snail with the command

    snail [-n=N]

Options
-------

    -n=N (Default: 5)
            Create an NxN snail matrix.

Example
-------

    $ snail -n 4
      1  2  3  4
     12 13 14  5
     11 16 15  6
     10  9  8  7
    $


Installation
-------------

Use goinstall to install snail

    goinstall github.com/bmatsuo/snail

General Documentation
---------------------

Use godoc to vew the documentation for snail

    godoc github.com/bmatsuo/snail

Or alternatively, use a godoc http server

    godoc -http=:6060

and view the url http://localhost:6060/pkg/github.com/bmatsuo/snail/

Author
======

Bryan Matsuo <bmatsuo@soe.ucsc.edu>

Copyright & License
===================

Copyright (c) 2011, Bryan Matsuo.
All rights reserved.

Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
