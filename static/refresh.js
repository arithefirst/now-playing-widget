export function tokenRefresh(client_id) {
    const redirect_uri = window.location.href.split("#")[0].replace(/\/$/, "");
    const scope = "user-read-playback-state";
    let url = "https://accounts.spotify.com/authorize";
    url += "?response_type=token";
    url += "&client_id=" + encodeURIComponent(client_id);
    url += "&scope=" + encodeURIComponent(scope);
    url += "&redirect_uri=" + encodeURIComponent(redirect_uri);
    window.location.replace(url);
}