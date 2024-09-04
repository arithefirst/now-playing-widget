# 🎵 Now Playing Widget 🎵

A simple widget to show the current song playing on your Spotify account. Useful for Twitch streams and YouTube videos.

![image](https://github.com/user-attachments/assets/42621fa6-3af0-4f4d-a2fd-96f1c273caf8)

# 🚀 How to Use 🚀

## Setting up the server (required)

### Creating the spotify app
1. Visit `https://developer.spotify.com/dashboard` and click "Create App"
2. Add an app name and description, these can be anything
3. Make sure to set the `Redirect URI` field to the URI of where your instance will be located. This field is a list of URIs that Spotify will be allowed to redirect to after exiting the Spotify sign-in page
4. Check the `Web API` Checkbox
5. Hit Save

### Set up the frontend files
1. Clone the repo, and open up `static/main.js` in your code or text editor of choice
2. On line 4, replace `http://127.0.0.1` with the URI of where your instance will be located
3. On line 5, replace `CLIENT_ID_HERE` with the client ID of the app you created on the Spotify developer portal
4. Repeat these steps for in `static/config/config.js`, but on lines 2 and 3 instead of 4 and 5.

### Installing MongoDB
1. See the [MongoDB Installation Guide](https://www.mongodb.com/docs/manual/installation/) and install MongoDB
2. Ensure that the MongoDB server runs on port `27017`

### Running the server
1. Download the latest release of the server software from the [releases page](https://github.com/arithefirst/now-playing-widget/releases/)
2. Run the binary for your operating system with the URL for your MongoDB Server (EX: `./nowplaying-server 10.0.0.21`)

> [!NOTE]
> Make sure the binary is in the same folder as the `/static` directory that you downloaded earlier. If they are not in the same directory, the server will not be able to server the files.

### Usage
1. To use your widget, visit the URI you set in the javascript files earlier
2. Sign in with spotify and enjoy your widget :3
3. For customization options, see [Changing the Color Scheme and Alignment](https://github.com/arithefirst/now-playing-widget?tab=readme-ov-file#changing-the-color-scheme-and-alignment-optional)

> [!NOTE]
> It is very important that you **do not** include the protocol, nor port for the MongoDB server.<br>
> - An example of what the command should look like: `./nowplaying-server 10.0.0.21`<br>
> - An example of what the command *should not* look like `./nowplaying-server mongodb://10.0.0.21` or `./nowplaying-server 10.0.0.21:27017`

## Changing the Color Scheme and Alignment (optional)

The default color scheme is pictured above, in a black and gray theme with left align. However, there is a build-in endpoint that allows you to control the theming and alignment of the widget. To apply a custom widget, follow the steps below.

1. Visit `YOUR_SERVER_URL/config`
2. Select the colors you would like to use for each element VIA the color pickers
3. Select the alignment you would like to use VIA the Toggle switch
4. Return to the regular endpoint, and see your new custom widget!

## Adding other users (optional)

By default, only the owner of the endpoint can use the widget through their spotify account. To give another user access to the widget, follow the steps below.

1. Visit `https://developer.spotify.com/dashboard` and click on the app you created when setting up the server.
2. Click the "Settings" button in the top right of the screen
3. Click "User Management"
4. Add the name and email of anyone you would like to be able to access your endpoint

> [!NOTE]  
> The "email" field must match the email on the account of the person whom you are trying to add.
