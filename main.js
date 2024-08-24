let timeoutId;

function tokenRefresh() {
  var client_id = "CLIENT_ID_HERE";
  var redirect_uri = window.location.href.split("#")[0].replace(/\/$/, "");
  var scope = "user-read-playback-state";
  var url = "https://accounts.spotify.com/authorize";
  url += "?response_type=token";
  url += "&client_id=" + encodeURIComponent(client_id);
  url += "&scope=" + encodeURIComponent(scope);
  url += "&redirect_uri=" + encodeURIComponent(redirect_uri);

  console.log(url);
  window.location.replace(url);
}

function truncateString(str, maxLength) {
  if (str.length > maxLength) {
    const shortStr = str.slice(0, maxLength - 3);
    return shortStr.replace(/ $/, "") + "...";
  }
  return str;
}

function getNowPlaying() {
  console.log("Refreshing...");
  const currentUrl = window.location.href;
  const fragment = currentUrl.split("#")[1];

  if (fragment) {
    const authToken = fragment.split("=")[1];
    const options = {
      method: "GET",
      headers: {
        Authorization: "Bearer " + authToken,
      },
    };

    timeoutId = setTimeout(getNowPlaying, 5000);
    fetch("https://api.spotify.com/v1/me/player", options)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        const jsonData = data;
        $("body").css("opacity", "1");
        $("#song-l").text(truncateString(jsonData["item"]["name"], 20));
        $("#song-r").text(truncateString(jsonData["item"]["name"], 20));
        $("#artist-l").text(truncateString(jsonData["item"]["artists"][0]["name"], 25));
        $("#artist-r").text(truncateString(jsonData["item"]["artists"][0]["name"], 25));
        $("#album-l").text(truncateString(jsonData["item"]["album"]["name"], 35));
        $("#album-r").text(truncateString(jsonData["item"]["album"]["name"], 35));
        $("#cover-r").attr("src", jsonData["item"]["album"]["images"][0]["url"]);
        $("#cover-l").attr("src", jsonData["item"]["album"]["images"][0]["url"]);
        $("#favicon").attr("href", jsonData["item"]["album"]["images"][0]["url"]);
      })
      .catch((error) => {
        console.error("Error:", error);
        if (error.message.includes("Network")) {
          tokenRefresh();
        } else {
          $("body").css("opacity", "0");
        }
      });
  } else {
    tokenRefresh();
  }
}

window.onload = function () {
  getNowPlaying();
  console.log('Retrived color cookies: {"STC": "' + Cookies.get("STC") + '",' + '"TC": "' + Cookies.get("TC") + '",' + '"BG": "' + Cookies.get("BG") + '"}');
  // If the cookies are not set, use default colors
  if (Cookies.get("STC") != undefined) {
    $("html").get(0).style.setProperty("--smalltext-color", Cookies.get("STC"));
  } else {
    $("html").get(0).style.setProperty("--smalltext-color", "#D3D3D3");
  }

  if (Cookies.get("TC") != undefined) {
    $("html").get(0).style.setProperty("--text-color", Cookies.get("TC"));
  } else {
    $("html").get(0).style.setProperty("--text-color", "#FFFFFF");
  }

  if (Cookies.get("BG") != undefined) {
    $("html").get(0).style.setProperty("--background-color", Cookies.get("BG"));
  } else {
    $("html").get(0).style.setProperty("--background-color", "#181A1B");
  }

  if (Cookies.get("LEFT") == "true") {
    $("#ALR").css("display", "none");
    $("#ALL").css("display", "inline");
  } else {
    $("#ALL").css("display", "none");
    $("#ALR").css("display", "inline");
  }
};
