===============================================================
Download Instagram Photo, Video, Story, Story Highlights in Go_
===============================================================

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


Download Instagram_ media (photos and videos), stories, and story highlights.


Obtain Cookies
++++++++++++++

The following three values are must to access the Instagram API.

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

First login to Instagram_ from Chrome browser, and there are three ways to get
the above information:

1. Use `Chrome extension in this repo <crx-cookies>`_ to get the cookies.

2. From `Chrome Developer Tools`_: See this `SO answer`_ or `Obtain cookies`_
   section in `instastories-backup`_ repo.

.. image:: https://i.stack.imgur.com/psJLZ.png
   :align: center
   :alt: ds_user_id sessionid csrftoken

3. From Chrome extension: Use EditThisCookie_ or `cookie-txt-export`_ or other
   cookie tools.


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

The name of this package is *igdl*. See
`example/download.go <example/download.go>`_ for how to use this package.

In the example, *download timeline* will run forever and download posts in your
timeline every 15 seconds.

*download story* will also run forever and download stories/postlives of your
following users every 30 seconds.

*download highlights* will run once and download story highlights of your
following users.


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
