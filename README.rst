=============================================================
Instagram Photo, Video, Story, Following, and Follower in Go_
=============================================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/instago?status.png
   :target: https://godoc.org/github.com/siongui/instago

.. image:: https://api.travis-ci.org/siongui/instago.png?branch=master
   :target: https://travis-ci.org/siongui/instago

.. image:: https://goreportcard.com/badge/github.com/siongui/instago
   :target: https://goreportcard.com/report/github.com/siongui/instago

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/instago/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/instago.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D


Get Instagram_ media (photos and videos), stories, story highlights, following
and followers in Go.


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

This package *instago* only access the Instagram public and private API and get
data from the API. If you want to download media (photos/videos), stories, or
story highlights. See `download <download>`_ directory.

Install the package by ``go get``:

.. code-block:: bash

  $ go get -u github.com/siongui/instago

You can use the following methods without cookies

- `GetUserInfoNoLogin <https://godoc.org/github.com/siongui/instago#GetUserInfoNoLogin>`_
- `GetRecentPostCodeNoLogin <https://godoc.org/github.com/siongui/instago#GetRecentPostCodeNoLogin>`_
- `GetUserId <https://godoc.org/github.com/siongui/instago#GetUserId>`_
- `GetPostInfoNoLogin <https://godoc.org/github.com/siongui/instago#GetPostInfoNoLogin>`_
- `GetUserProfilePicUrlHd <https://godoc.org/github.com/siongui/instago#GetUserProfilePicUrlHd>`_

For the other methods which need cookies to access Instagram API, you must call
NewInstagramApiManager_ first:

.. code-block:: go

  import (
  	"github.com/siongui/instago"
  )

  mgr := instago.NewInstagramApiManager("IG_DS_USER_ID", "IG_SESSIONID", "IG_CSRFTOKEN")

Then you can use *mgr* to get data from Instagram API. For example, you can get
all post codes of the user `instagram <https://www.instagram.com/instagram/>`__
as follows:

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


Private API
+++++++++++

- `Get data from Instagram's private API — Alberto Moral <https://www.albertomoral.com/blog/get-data-from-instagrams-private-api>`_
- `What is the API Endpoints for the Feeds "People who liked my posts" and "Activities from my followings" · Issue #42 · huttarichard/instagram-private-api · GitHub <https://github.com/huttarichard/instagram-private-api/issues/42>`_


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
