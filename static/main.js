let timeoutId;

// Set the base URL & Client ID of your endpoint Below
const baseurl = "http://127.0.0.1";
const client_id = "CLIENT_ID_HERE";

function tokenRefresh() {
  const redirect_uri = window.location.href.split("#")[0].replace(/\/$/, "");
  const scope = "user-read-playback-state";
  let url = "https://accounts.spotify.com/authorize";
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
  // Pull the authtoken from the URL; If not in the url try to get a new one
  console.log("Refreshing...");
  const fragment = window.location.href.split("#")[1];

  if (fragment) {
    const authToken = fragment.split("=")[1];
    const options = {
      method: "GET",
      headers: {
        Authorization: "Bearer " + authToken,
      },
    };

    // Set timeout so getNowPlaying() reruns every 5000ms
    timeoutId = setTimeout(getNowPlaying, 5000);
    fetch("https://api.spotify.com/v1/me/player", options)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        // Load the data from the json response into the required elements using jQuery
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
        // Check to see if the error is a network error, if not hide the body
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

// Function to get the UID of a currently signed in user
async function getUid() {
  // No need to check for fragment since this is only ever called
  // after running getNowPlaying()

  const authToken = window.location.href.split("#")[1].split("=")[1];
  try {
    const response = await fetch("https://api.spotify.com/v1/me", {
      headers: {
        Authorization: `Bearer ${authToken}`,
      },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data.id;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
}

// Function to get the hex value for a user from the API
async function getHex(uid, type) {
  try {
    const response = await fetch(`${baseurl}/api/v1/get/${type.toLowerCase()}?uid=${uid}`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    if (data.err != null && data.err != "No value found: Default returned") {
      console.log(data.err);
    }
    return data.hex;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
}

// On page load:
window.onload = async function () {
  getNowPlaying();

  const uid = await getUid();
  const stc = await getHex(uid, "stc");
  const tc = await getHex(uid, "tc");
  const bg = await getHex(uid, "bg");
  const right = await getHex(uid, "right");
  console.log(`Got Config: {STC: ${stc}, TC: ${tc}, BG: ${bg}, RIGHT: ${right}`);

  // Set the CSS color variables based on the response from the API
  $("html").get(0).style.setProperty("--smalltext-color", stc);
  $("html").get(0).style.setProperty("--text-color", tc);
  $("html").get(0).style.setProperty("--background-color", bg);

  // If the right var is "true", hide the ALL (Align Left)
  // class and make the ALR (Aligh Right) class visible

  if (right == "true") {
    $("#ALR").css("display", "inline");
    $("#ALL").css("display", "none");
  } else {
    $("#ALL").css("display", "inline");
    $("#ALR").css("display", "none");
  }
};
