package spotify

var searchResponse = `{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?query=bohemian+rhapsody&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.5&offset=0&limit=20",
    "items": [
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
              },
              "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
              "id": "1dfeR4HaWDbWqFHLkxsg1d",
              "name": "Queen",
              "type": "artist",
              "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/6i6folBtxKV28WX3msQ4FE"
          },
          "href": "https://api.spotify.com/v1/albums/6i6folBtxKV28WX3msQ4FE",
          "id": "6i6folBtxKV28WX3msQ4FE",
          "images": [
            {
              "height": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
              "width": 640
            },
            {
              "height": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
              "width": 300
            },
            {
              "height": 64,
              "url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
              "width": 64
            }
          ],
          "is_playable": true,
          "name": "Bohemian Rhapsody (The Original Soundtrack)",
          "release_date": "2018-10-19",
          "release_date_precision": "day",
          "total_tracks": 22,
          "type": "album",
          "uri": "spotify:album:6i6folBtxKV28WX3msQ4FE"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ],
        "disc_number": 1,
        "duration_ms": 354947,
        "explicit": false,
        "external_ids": {
          "isrc": "GBUM71029604"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/3z8h0TU7ReDPLIbEnYhWZb"
        },
        "href": "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
        "id": "3z8h0TU7ReDPLIbEnYhWZb",
        "is_local": false,
        "is_playable": true,
        "name": "Bohemian Rhapsody",
        "popularity": 71,
        "preview_url": null,
        "track_number": 7,
        "type": "track",
        "uri": "spotify:track:3z8h0TU7ReDPLIbEnYhWZb"
      }
    ],
    "limit": 20,
    "next": "https://api.spotify.com/v1/search?query=bohemian+rhapsody&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.5&offset=20&limit=20",
    "offset": 0,
    "previous": null,
    "total": 907
  }
}`

var recommendationResponse = `{
  "tracks": [
    {
      "album": {
        "album_type": "album",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ],
        "external_urls": {
          "spotify": "https://open.spotify.com/album/6i6folBtxKV28WX3msQ4FE"
        },
        "href": "https://api.spotify.com/v1/albums/6i6folBtxKV28WX3msQ4FE",
        "id": "6i6folBtxKV28WX3msQ4FE",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
            "width": 64
          }
        ],
        "is_playable": true,
        "name": "Bohemian Rhapsody (The Original Soundtrack)",
        "release_date": "2018-10-19",
        "release_date_precision": "day",
        "total_tracks": 22,
        "type": "album",
        "uri": "spotify:album:6i6folBtxKV28WX3msQ4FE"
      },
      "artists": [
        {
          "external_urls": {
            "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
          },
          "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
          "id": "1dfeR4HaWDbWqFHLkxsg1d",
          "name": "Queen",
          "type": "artist",
          "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
        }
      ],
      "disc_number": 1,
      "duration_ms": 354947,
      "explicit": false,
      "external_ids": {
        "isrc": "GBUM71029604"
      },
      "external_urls": {
        "spotify": "https://open.spotify.com/track/3z8h0TU7ReDPLIbEnYhWZb"
      },
      "href": "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
      "id": "3z8h0TU7ReDPLIbEnYhWZb",
      "is_local": false,
      "is_playable": true,
      "name": "Bohemian Rhapsody",
      "popularity": 71,
      "preview_url": null,
      "track_number": 7,
      "type": "track",
      "uri": "spotify:track:3z8h0TU7ReDPLIbEnYhWZb"
    }
  ]
}`
