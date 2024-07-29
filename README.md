# 🎵 Now Playing Widget 🎵
A simple widget to show the current song playing on your Spotify account. Useful for Twitch streams and YouTube videos.

![image](https://github.com/user-attachments/assets/9fa2cb0a-c0ac-46c7-8dcc-1deede7ae816)

# 🚀 How to Use 🚀
## Published
To use the already published widget, add `https://spotify.mildlylargefish.org` as a browser source in your OBS Studio, or screen capturing program of preference. Then sign in with your Spotify account and play music from any device signed into your Spotify account.

> [!NOTE]
> You will only be signing into Spotify through the official Spotify URL, which you will be redirected to after visiting https://spotify.mildlylargefish.org. Sensitive information such as your email or password will never be shared with the developers of this program.

## Self Hosted
To host an instance of the widget, follow the steps below.
1. Visit `https://developer.spotify.com/dashboard` and click "Create App"
2. Add an app name and description, these can be anything
3. Make sure to set the `Redirect URI` field to the URI where your instance will be located. This field is a list of URIs that Spotify will be allowed to redirect to after exiting the Spotify sign-in page.
4. Check the `Web API` Checkbox
5. Hit Save
6. Clone the repo, and open up `auth.json` in your code editor of choice
7. On line 4, replace `CLIENT_ID_HERE` with the client ID of the app you created on the Spotify developer portal
8. On line 5, replace `REDIRECT_URI_HERE` with the Redirect URI you set on the Spotify developer portal
9. Host the 3 files on an HTTP Server at the URI you set in the developer portal
10. Visit the URI you set in the developer portal
