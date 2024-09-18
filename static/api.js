import { tokenRefresh } from "./refresh.js";

// Function to get the saved hex codes for a user
export async function getHex(uid, type, baseurl) {
    try {
        const response = await fetch(`${baseurl}/api/v1/get/${type.toLowerCase()}?uid=${uid}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        if (data.err !== "null" && data.err !== "No value found: Default returned") {
            console.error(data.err);
        }
        return data.hex;
    } catch (error) {
        console.error("Error:", error);
        throw error;
    }
}

// Function to get currently signed-in user's UUID
export async function getUid(client_id) {
    const fragment = window.location.href.split("#")[1];
    if (!fragment) {
        tokenRefresh(client_id);
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