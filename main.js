let timeoutId;

function tokenRefresh() {
  var client_id = "CLIENT_ID_HERE";
  var redirect_uri = "REDIRECT_URI_HERE";
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
      return str.slice(0, maxLength - 3) + '...';
  }
  return str;
}

function getNowPlaying() {
  console.log("Refreshing...")
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
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => {
        const jsonData = data;
        document.getElementById("data").style = "opacity: 1;";
        document.getElementById("dataContainer").style = "opacity: 1;"
        document.getElementById("song").innerHTML = truncateString(jsonData["item"]["name"], 20);
        document.getElementById("artist").innerHTML = truncateString(jsonData["item"]["artists"][0]["name"], 25);
        document.getElementById("album").innerHTML = truncateString(jsonData["item"]["album"]["name"], 35);
        document.getElementById("cover").src = jsonData["item"]["album"]["images"][0]["url"];
        document.getElementById("favicon").href = jsonData["item"]["album"]["images"][0]["url"];
        console.log(jsonData)
      })
      .catch((error) => {
        console.error("Error:", error);
        if (error.message.includes("Network")) {
          tokenRefresh();
        } else {
          document.getElementById("data").style = "opacity: 0;";
          document.getElementById("dataContainer").style = "opacity: 0;"
        }
        
      });
  } else {
    tokenRefresh();
  }
}
