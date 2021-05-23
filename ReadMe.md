# RSS 3 Go Hub

> You are the hub, of all you have, of all you love.
> Keep going.

## About

We've split the RSS3 types into a new repo [rss3go_lib](https://github.com/nyawork/rss3go_lib)

It's still in early develop stage, and those it cannot run correctly / functionally. Hope we can finally make it.

Check [RSS3-Hub](https://github.com/NaturalSelectionLabs/RSS3-Hub/) and [RSS3](https://rss3.io/) for more information.

## Auth

RSS3-Hub is using `eth-crypto` as auth function now.

But, there seems to be no such crypto library for go. So I might have to implement one.

For RSS3-Hub compatible auth methods, please
check [RSS3-Hub#authorization](https://github.com/NaturalSelectionLabs/RSS3-Hub#authorization) for detailed information.

## Endpoints

> Copied from RSS3-Hub

### Files

- GET `/files/:fid` - get a file

- POST `/files/:fid` - upload a new file (AUTH Required)

- PUT `/files/:fid` - change a file (AUTH Required)

- DELETE `/files/:fid` - delete a file (AUTH Required)
