===========================================================================================
Download Instagram Photo, Video, Story, Highlight, Postlive, Following, and Follower in Go_
===========================================================================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/instago/download?status.png
   :target: https://godoc.org/github.com/siongui/instago/download

.. image:: https://api.travis-ci.org/siongui/instago.png?branch=master
   :target: https://travis-ci.org/siongui/instago

.. image:: https://goreportcard.com/badge/github.com/siongui/instago/download
   :target: https://goreportcard.com/report/github.com/siongui/instago/download

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/instago/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/instago.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D


Download Instagram_ media (photos and videos), stories, story highlights,
postlives (live stream that shared to stories after end) in Go.


Obtain Cookies
++++++++++++++

Use `Chrome extension in this repo <crx-cookies>`_ to get the cookies. Save it
as *auth.json*. We will use it later to access Instagram API.


Terminology
+++++++++++

Given the URL of the post as follows:

::

  https://www.instagram.com/p/BfJzG64BZVY/

The *code* of the post is **BfJzG64BZVY**.


Usage
+++++

You have to install wget_ and ffmpeg_ first. Because of this reason, this
package works only on Linux systems currently. For Ubuntu users, wget comes with
the distribution by default, and you can install ffmpeg_ as follow:

.. code-block:: bash

  $ sudo apt-get install ffmpeg

Then you can install this package:

.. code-block:: bash

  $ go get -u github.com/siongui/instago/download

The name of this package is *igdl*.

The following are examples that you may be interested in:

- `timeline.go <example/timeline.go>`_: download posts in your timeline.
- `storypostlive.go <example/storypostlive.go>`_: download stories and postlives
  of your following users.
- `highlights.go <example/highlights.go>`_: download story highlights of all
  following users.
- `allposts.go <example/allposts.go>`_: download all posts of a single user.
- `allpostsnologin.go <example/allpostsnologin.go>`_: download all posts of a
  single user without login
- `userstory.go <example/userstory.go>`_: given username, download unexpired
  stories of the user.
- `userstoryhighlight.go <example/userstoryhighlight.go>`_: given username,
  download story highlights of the user.
- `saved.go <example/saved.go>`_: download all your saved posts.

See godoc_ for complete list of download methods.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


.. _Go: https://golang.org/
.. _Instagram: https://www.instagram.com/
.. _Chrome Developer Tools: https://developer.chrome.com/devtools
.. _SO answer: https://stackoverflow.com/a/44773079
.. _Obtain cookies: https://github.com/hoschiCZ/instastories-backup#obtain-cookies
.. _instastories-backup: https://github.com/hoschiCZ/instastories-backup
.. _EditThisCookie: https://www.google.com/search?q=EditThisCookie
.. _cookie-txt-export: https://github.com/siongui/cookie-txt-export.go
.. _UNLICENSE: http://unlicense.org/
.. _wget: https://www.gnu.org/software/wget/
.. _ffmpeg: https://www.ffmpeg.org/
.. _godoc: https://godoc.org/github.com/siongui/instago/download
