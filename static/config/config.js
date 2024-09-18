import {getHex, getUid} from "../api.js";
import {tokenRefresh} from "../refresh.js";

// Set the base URL & Client ID of your endpoint Below
const baseurl = "http://127.0.0.1";
const client_id = "CLIENT_ID_HERE";

// Function to save a user's config
async function saveConf(tc, bg, stc, right) {
  // Pull the authtoken from the URL
  const fragment = window.location.href.split("#")[1];
  // If token not present in url, reload
  if (!fragment) {
    tokenRefresh(client_id);
  } else {
    const uid = await getUid(client_id);
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

async function getConf() {
  const uid = await getUid(client_id);
  const stc = await getHex(uid, "stc", baseurl);
  const tc = await getHex(uid, "tc", baseurl);
  const bg = await getHex(uid, "bg", baseurl);
  const right = await getHex(uid, "right", baseurl);
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
