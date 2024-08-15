# 🎵 Now Playing Widget 🎵

A simple widget to show the current song playing on your Spotify account. Useful for Twitch streams and YouTube videos.

![image](https://github.com/user-attachments/assets/9fa2cb0a-c0ac-46c7-8dcc-1deede7ae816)

# 🚀 How to Use 🚀

## Setting up the server (required)

To host an instance of the widget, follow the steps below.

1. Visit `https://developer.spotify.com/dashboard` and click "Create App"
2. Add an app name and description, these can be anything
3. Make sure to set the `Redirect URI` field to the URI where your instance will be located. This field is a list of URIs that Spotify will be allowed to redirect to after exiting the Spotify sign-in page
4. Check the `Web API` Checkbox
5. Hit Save
6. Clone the repo, and open up `main.js` in your code or text editor of choice
7. On line 4, replace `CLIENT_ID_HERE` with the client ID of the app you created on the Spotify developer portal
8. Host all of the files from this repo (excluding the README and LICENSE) on an HTTP Server at the URI you set in the developer portal
9. Visit the URI you set in the developer portal and login via spotify

## Changing the Color Scheme
The default color scheme is pictured above, in a black and gray theme. However, there is a build-in endpoint that allows you to control the theming of the widget. To apply a custom widget, follow the steps below.
1. Visit `YOUR_SERVER_URL/config.html`
2. Select the colors you would like to use for each element via the color pickers
3. Return to the regular endpoint, and see your new custom widget!

> [!NOTE]  
> The color customization is stored in browser cookies, meaning you need to follow the above steps from the browser you indent to view the widget on (EX: Your OBS browser source)

## Adding other users (optional)

By default, only the owner of the endpoint can use the widget through their spotify account. To give another user access to the widget, follow the steps below.

1. Visit `https://developer.spotify.com/dashboard` and click on the app you created when setting up the server.
2. Click the "Settings" button in the top right of the screen
3. Click "User Managment"
4. Add the name and email of anyone you would like to be able to access your endpoint

> [!NOTE]  
> The "email" field must match the email on the account of the person whom you are trying to add.
