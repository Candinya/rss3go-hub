# RSS 3 Go Hub

> You are the hub, of all you have, of all you love.
> Keep going.

[![issue](https://img.shields.io/github/issues/nyawork/rss3go_hub)](https://github.com/nyawork/rss3go_hub/issues)
[![dev build](https://img.shields.io/github/workflow/status/nyawork/rss3go_hub/RSS3Go-Hub%20Docker%20Build/dev?label=Dev%20Build)](https://hub.docker.com/r/nyawork/rss3go-hub/tags?name=dev)
[![master build](https://img.shields.io/github/workflow/status/nyawork/rss3go_hub/RSS3Go-Hub%20Docker%20Build/master?label=Release%20Build)](https://hub.docker.com/r/nyawork/rss3go-hub/tags?name=master)
[![license](https://img.shields.io/github/license/nyawork/rss3go_hub)](https://github.com/nyawork/rss3go_hub/blob/master/LICENSE)

## About

It's still in early develop stage, and those it may have some functional questions. Just open an issue if necessary.

Check [RSS3-Hub](https://github.com/NaturalSelectionLabs/RSS3-Hub/) and [RSS3](https://rss3.io/) for more information.

## Deploy

You can use docker and docker-compose for quick deployments now.

``` shell
# Grab the docker-compose file
wget https://raw.githubusercontent.com/nyawork/rss3go_hub/dev/docker-compose.yml

# Grap default configuration file
wget https://raw.githubusercontent.com/nyawork/rss3go_hub/dev/config.yml

# Now you can edit those files as you like.
```

To start up RSS3Go-Hub or for further upgrade, just run follow commands:

``` shell

# Pull image
docker-compose pull

# Start container
docker-compose up -d

```

Sometimes we may change the above two `yml` file for more functional implements. Don't forget to check them out!

## Authentication

For better compatibility, we are using the same methods for signature authentication in RSS3-Hub.

They are already implemented in [rss3go_lib](https://github.com/nyawork/rss3go_lib).

## Endpoints

### Files

- GET `/files/:fid` - get a file

  Request body: Null

- PUT `/files` - change a file (Signature within)

  Request body: Arrays of file object with signature (Might already exist)

    ``` json
    {
      "contents": [{
          "id": "0x.......",
          "...": "..."
      }, {
          "id": "0x.......",
          "...": "..."
      }]
    }
    ```

- DELETE `/files` - delete a file (Signature within)

  Request body: (Your hexadecimal signature of string "delete")

    ``` json
    {
        "signature": "xxxxxx00" // Should end with 00 or 01
    }
    ```

