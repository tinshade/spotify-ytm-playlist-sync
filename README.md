# spotify-ytm-playlist-sync
Quick and dirty solution to sync playlists from Spotify to YouTube Music

## Process Flow

1. Get the Spotify playlist UUID from the user  
    - Spotify needs a dev account to access APIs
    - Spotify APIs require JWT tokens to work. 
        - No restriction on whose token is used to acccess what playlist unless the playlist is private
        - Playlists have to be public, or use user's JWT but having them Login via Spotify
    - APIs needed:
        1. Token API for machine to machine token
        2. API to grab user's Spotify JWT
        3. API to get playlist UUID

2. Get the list of all available songs in that Spotify playlist 
    - Some playlists can be huge, implement limit and offset
    - Implement chunking and parallell processing to go through the playlist quickly
    - Does it matter if the playlist is sorted?  
    - APIs needed:
        1. API to get tracks for a given playlist
            
3. Check for all available songs in YT Music
    - Youtube Music API needs a Google developer account and an app with full permissions to access YTMusic services
    - Do I need JWTs? Will mine work or do I have to ask the user to sign-in and grab their JWT?
    - APIs needed:
        1. API to search song by name
        2. API to create a new playlist
        3. API to add songs to a playlist (new / old)

4. Create a new YTMusic playlist with all available songs found in YTMusic based on Spotify Playlist
    - APIs needed:
        1. API to create a new playlist
        2. API to add songs to a playlist (new / old)

5. Notify the user of # of songs found vs. # of songs ported

6. Done 



### References
1. Spotify Web Dev API: https://developer.spotify.com/documentation/web-api
2. Youtube Music API: N/A 
    - Found a Golang wrapper : https://pkg.go.dev/github.com/prettyirrelevant/ytmusicapi
    - Found a Python wrapper : https://github.com/sigma67/ytmusicapi
3. Google Console: https://console.cloud.google.com/