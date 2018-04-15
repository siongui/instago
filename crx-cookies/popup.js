chrome.tabs.query({
  active: true,
  currentWindow: true
}, function(tabs) {
  chrome.cookies.getAll({}, function (cookies) {

    var cookieNames = ["ds_user_id", "sessionid", "csrftoken"];

    var cookieAuth = {};
    document.write("<pre>");
    for (var i in cookies) {

      var cookie = cookies[i];
      if (cookieNames.indexOf(cookie.name) == -1) {
        continue;
      }

      cookieAuth[cookie.name] = cookie.value;
    }
    document.write(JSON.stringify(cookieAuth, null, 2));
    document.write("</pre>");
  });
});
