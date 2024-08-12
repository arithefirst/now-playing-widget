# 🎵 Now Playing Widget 🎵
A simple widget to show the current song playing on your Spotify account. Useful for Twitch streams and YouTube videos.

![image](https://github.com/user-attachments/assets/9fa2cb0a-c0ac-46c7-8dcc-1deede7ae816)

# 🚀 How to Use 🚀

## Setting up the server (required)
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
10. Visit the URI you set in the developer portal and login via spotify

## Adding other users (optional)
By default, only the owner of the endpoint can use the widget through their spotify account. To give another user access to the widget, follow the steps below.
1. Visit `https://developer.spotify.com/dashboard` and click on the app you created when setting up the server.
2. Click the "Settings" button in the top right of the screen
3. Click "User Managment"
4. Add the name and email of anyone you would like to be able to access your endpoint

> [!NOTE]  
> The "Full Name" field does not need to match the name on the user's account.
