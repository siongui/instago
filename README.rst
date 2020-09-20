==================================================================================
Instagram Photo, Video, Story, Highlight, Postlive, Following, and Follower in Go_
==================================================================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/instago?status.svg
   :target: https://godoc.org/github.com/siongui/instago

.. image:: https://api.travis-ci.org/siongui/instago.svg?branch=master
   :target: https://travis-ci.org/siongui/instago

.. image:: https://goreportcard.com/badge/github.com/siongui/instago
   :target: https://goreportcard.com/report/github.com/siongui/instago

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/instago/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/instago.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D


Get Instagram_ media (photos and videos), stories, story highlights, postlives
(live stream that shared to stories after end), following and followers in Go.


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

This package *instago* only access the Instagram public and private API and get
metadata from the API. If you want to download media (photos/videos), stories,
story highlights, or postlives, see `download <download>`_ directory.

Install the package by ``go get``:

.. code-block:: bash

  $ go get -u github.com/siongui/instago

You can use the following methods without cookies

- `GetUserInfoNoLogin <https://godoc.org/github.com/siongui/instago#GetUserInfoNoLogin>`_
- `GetRecentPostCodeNoLogin <https://godoc.org/github.com/siongui/instago#GetRecentPostCodeNoLogin>`_
- `GetUserId <https://godoc.org/github.com/siongui/instago#GetUserId>`_
- `GetPostInfoNoLogin <https://godoc.org/github.com/siongui/instago#GetPostInfoNoLogin>`_
- `GetUserProfilePicUrlHd <https://godoc.org/github.com/siongui/instago#GetUserProfilePicUrlHd>`_
- `GetAllPostMediaNoLogin <https://godoc.org/github.com/siongui/instago#GetAllPostMediaNoLogin>`_

For the other methods which need cookies to access Instagram API, you must call
NewInstagramApiManager_ first:

.. code-block:: go

  import (
  	"github.com/siongui/instago"
  )

  mgr := instago.NewInstagramApiManager("auth.json")

Then you can use *mgr* to get metadata from Instagram API. For example, you can
get all post codes of the user
`instagram <https://www.instagram.com/instagram/>`__ as follows:

.. code-block:: go

  codes, err := mgr.GetAllPostCode("instagram")
  if err != nil {
  	panic(err)
  }

  for _, code := range codes {
  	println("URL: https://www.instagram.com/p/%s/\n", code)
  }

For complete examples, see test files (files ends with *_test.go*). The
following are some examples you may be interested in:

- Get post information: See `post_test.go <post_test.go>`_
- Get URLs of all posts of a specific user: See `getall_test.go <getall_test.go>`_
- Get id by username: See `userinfo_test.go <userinfo_test.go>`_
- Discover top live: See `toplive_test.go <toplive_test.go>`_
- Top searches of Instagram web: See `topsearch_test.go <topsearch_test.go>`_


Tricks
++++++

- Use the following User-Agent to get post-live field in reels tray feed.

  **Instagram 10.26.0 (iPhone8,1; iOS 10_2; en_US; en-US; scale=2.00; gamut=normal; 750x1334) AppleWebKit/420+**

  From `replay.py`_ in `instagram_private_api_extensions`_

- Get all user's media:

  * `How can I get a user's media from Instagram without authenticating as a user? - Stack Overflow <https://stackoverflow.com/a/47243409>`_
  * `instagram_web_api.client — instagram_private_api 1.4.1 documentation <https://instagram-private-api.readthedocs.io/en/latest/_modules/instagram_web_api/client.html#Client.user_feed>`_
  * `instagram graphql api id - Google search <https://www.google.com/search?q=instagram+graphql+api+id>`_

- `Web scraping: instagram.com | Shiori <https://kaijento.github.io/2017/05/17/web-scraping-instagram.com/>`_

- | `query_hash on instagram graphql - Google search <https://www.google.com/search?q=query_hash+on+instagram+graphql>`_
  | `How to scrape pages with infinite scroll: extracting data from Instagram - Diggernaut <https://www.diggernaut.com/blog/how-to-scrape-pages-infinite-scroll-extracting-data-from-instagram/>`_
  | `colly instagram example <https://github.com/gocolly/colly/blob/master/_examples/instagram/instagram.go>`_

- Do not remove query string in the URLs of photo/viedo/story/highlight. It may
  cause 403 Forbidden when downloading the URL. See `issue #2`_ for more info.

- Saved endpoint: see `ping/instagram_private_api <https://github.com/ping/instagram_private_api>`_

- postlive URL issue: Google search "Bad URL timestamp". See `Instagram reports "Bad URL timestamp" <https://www.reddit.com/r/ifttt/comments/e79x24/instagram_reports_bad_url_timestamp/>`_. replacing &amp; with & in the access link.


Private API
+++++++++++

- `Get data from Instagram's private API — Alberto Moral <https://www.albertomoral.com/blog/get-data-from-instagrams-private-api>`_
- `What is the API Endpoints for the Feeds "People who liked my posts" and "Activities from my followings" · Issue #42 · huttarichard/instagram-private-api · GitHub <https://github.com/huttarichard/instagram-private-api/issues/42>`_
- `Search · go instagram · GitHub <https://github.com/search?q=go+instagram>`_
  then found `Update timeline API from Get to Post <https://github.com/hieven/go-instagram/commit/6800b3f7b9513fb0084024e405109d939572a961>`_

UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GitHub - siongui/goiguserid: Get id of Instagram user in Go <https://github.com/siongui/goiguserid>`_
.. [2] `GitHub - siongui/goigstorylink: Get Links (URL) of Instagram Stories in Go <https://github.com/siongui/goigstorylink>`_
.. [3] `GitHub - siongui/goigfollow: Get Instagram following and followers in Go <https://github.com/siongui/goigfollow>`_
.. [4] `GitHub - siongui/goigstorydl: Download Instagram Stories in Go <https://github.com/siongui/goigstorydl>`_
.. [5] `GitHub - siongui/goigmedia: Get links of Instagram user media (photos and videos) in Go. <https://github.com/siongui/goigmedia>`_
.. [6] `JSON Formatter & Validator <https://jsonformatter.curiousconcept.com/>`_

.. _Go: https://golang.org/
.. _Instagram: https://www.instagram.com/
.. _Chrome Developer Tools: https://developer.chrome.com/devtools
.. _SO answer: https://stackoverflow.com/a/44773079
.. _Obtain cookies: https://github.com/hoschiCZ/instastories-backup#obtain-cookies
.. _instastories-backup: https://github.com/hoschiCZ/instastories-backup
.. _EditThisCookie: https://www.google.com/search?q=EditThisCookie
.. _cookie-txt-export: https://github.com/siongui/cookie-txt-export.go
.. _UNLICENSE: http://unlicense.org/
.. _replay.py: https://github.com/ping/instagram_private_api_extensions/blob/master/instagram_private_api_extensions/replay.py
.. _instagram_private_api_extensions: https://github.com/ping/instagram_private_api_extensions
.. _NewInstagramApiManager: https://godoc.org/github.com/siongui/instago#NewInstagramApiManager
.. _issue #2: https://github.com/siongui/instago/issues/2
