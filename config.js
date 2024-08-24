function setCSS() {
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
  generateCookies($("#tc").val(), $("#bg").val(), $("#stc").val());
}

function generateCookies() {
  var expiresDays = (1000 * 365 * 24 * 60 * 60 * 1000) / (24 * 60 * 60 * 1000);
  Cookies.set("STC", $("#stc").val(), { expires: expiresDays, sameSite: "none" });
  Cookies.set("TC", $("#tc").val(), { expires: expiresDays, sameSite: "none" });
  Cookies.set("BG", $("#bg").val(), { expires: expiresDays, sameSite: "none" });
  Cookies.set("LEFT", !$("#toggle").is(":checked"), { expires: expiresDays, sameSite: "none" });
  console.log(`Set Cookies: {STC: ${$("#stc").val()}, TC: ${$("#tc").val()}, BG: ${$("#bg").val()}, LEFT: ${!$("#toggle").is(":checked")}}`);
}

function getCookies() {
  console.log(`Got Cookies: {STC: ${Cookies.get("STC")}, TC: ${Cookies.get("TC")}, BG: ${Cookies.get("BG")}, LEFT: ${Cookies.get("LEFT")}}`);
  // If the cookies are not set, use default colors
  if (Cookies.get("STC") != undefined) {
    $("#stc").val(Cookies.get("STC"));
  } else {
    $("#stc").val("#D3D3D3");
  }

  if (Cookies.get("TC") != undefined) {
    $("#tc").val(Cookies.get("TC"));
  } else {
    $("#tc").val("#FFFFFF");
  }

  if (Cookies.get("BG") != undefined) {
    $("#bg").val(Cookies.get("BG"));
  } else {
    $("#bg").val("#181A1B");
  }
}

window.onload = function () {
  getCookies();
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
