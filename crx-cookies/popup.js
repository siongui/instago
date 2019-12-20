chrome.tabs.query({
  active: true,
  currentWindow: true
}, function(tabs) {
  chrome.cookies.getAll({}, function (cookies) {

    var cookieAuth = {};
    document.write("<pre>");
    for (var i in cookies) {
      var cookie = cookies[i];
      cookieAuth[cookie.name] = cookie.value;
    }
    document.write(JSON.stringify(cookieAuth, null, 2));
    document.write("</pre>");
  });
});
