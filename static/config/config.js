// Set the base URL & Client ID of your endpoint Below
const baseurl = "http://127.0.0.1:5500";
const client_id = "CLIENT_ID_HERE";

async function setCSS() {
  $("html").get(0).style.setProperty("--smalltext-color", $("#stc").val());
  $("html").get(0).style.setProperty("--text-color", $("#tc").val());
  $("html").get(0).style.setProperty("--background-color", $("#bg").val());
  if ($("#toggle").is(":checked")) {
    $("#ALL").css("display", "none");
    $("#ALR").css("display", "block");
  } else {
    $("#ALR").css("display", "none");
    $("#ALL").css("display", "block");
  }
  await saveConf($("#tc").val(), $("#bg").val(), $("#stc").val(), $("#toggle").is(":checked"));
}

// Function for getting the spotify api token
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

async function saveConf(tc, bg, stc, right) {
  // Pull the authtoken from the URL
  const fragment = window.location.href.split("#")[1];
  // If token not present in url, reload
  if (!fragment) {
    tokenRefresh();
  } else {
    const uid = await getUid();
    const authToken = fragment.split("=")[1];
    try {

      const jsonBody = {
        "uid": uid,
        "bg": bg,
        "stc": stc,
        "tc": stc,
        "right": right
      };

      const response = await fetch(`${baseurl}/api/v1/set`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${authToken}`
        },
        body: JSON.stringify(jsonBody),
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log(`Saved Config: {STC: ${stc}, TC: ${tc}, BG: ${bg}, RIGHT: ${right}`);
    } catch (error) {
      console.error("Error:", error);
      throw error;
    }
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

// Function to get the UID for the currently signed in user
async function getUid() {
  const fragment = window.location.href.split("#")[1];
  if (!fragment) {
    tokenRefresh();
  } else {
    const authToken = fragment.split("=")[1];
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
}

async function getConf() {
  const uid = await getUid();
  const stc = await getHex(uid, "stc");
  const tc = await getHex(uid, "tc");
  const bg = await getHex(uid, "bg");
  const right = await getHex(uid, "right");
  console.log(`Got Config: {STC: ${stc}, TC: ${tc}, BG: ${bg}, RIGHT: ${right}`);

  // Load the config
  $("#stc").val(stc);
  $("#tc").val(tc);
  $("#bg").val(bg);

  // Set the toggle/align
  if (right == "true") {
    $("#toggle").prop("checked", true).trigger("change");
  } else {
    $("#toggle").prop("checked", false).trigger("change");
  }
}

window.onload = async function () {
  await getConf();
  $("html").get(0).style.setProperty("--smalltext-color", $("#stc").val());
  $("html").get(0).style.setProperty("--text-color", $("#tc").val());
  $("html").get(0).style.setProperty("--background-color", $("#bg").val());
  if ($("#toggle").is(":checked")) {
    $("#ALL").css("display", "none");
    $("#ALR").css("display", "block");
  } else {
    $("#ALR").css("display", "none");
    $("#ALL").css("display", "block");
  }
};
