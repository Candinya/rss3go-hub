# RSS 3 Go

> There's no centralized hub in the RSS3 world. 
> You are the hub, of all you have, of all you love.
> Keep going.

## About

**It's my first time using Go so this mess is only for study and research use.**

It's still in early develop stage, and those it cannot run correctly / functionally. Hope we can finally make it.

Check [RSS3-Hub](https://github.com/NaturalSelectionLabs/RSS3-Hub/) and [RSS3](https://rss3.io/) for more information.

## Auth

RSS3-Hub is using `secp256k1` as auth function. 
In addition, in RSS3Go I *may* deploy `cv25519` as another auth curve function.
Check [SafeCurves](https://safecurves.cr.yp.to/) for more information.

For RSS3-Hub compatible auth methods, please check [RSS3-Hub#authorization](https://github.com/NaturalSelectionLabs/RSS3-Hub#authorization) for detailed information.

## Endpoints

> Copied from RSS3-Hub

### Personas

- GET `/personas/:pid` - get a persona

- POST `/personas` - add a new persona

    - Body parameters
    
        | Name   | Optional |
        | ------ | -------- |
        | id     | false    |
        | name   | true     |
        | avatar | true     |
        | bio    | true     |

- PATCH `/personas/:pid` - change a persona

    - Body parameters
    
        | Name   | Optional |
        | ------ | -------- |
        | name   | true     |
        | avatar | true     |
        | bio    | true     |

- DELETE `/personas/:pid` - delete a persona

#### Items

- GET `/personas/:pid/items` - get items of a persona

    - Url parameters
    
        | Name | Optional | Description                                                               |
        | ---- | -------- | ------------------------------------------------------------------------- |
        | id   | true     | file id of items file, empty for returning the data from the persona file |

- POST `/personas/:pid/items` - add a item to a persona

    - Body parameters
    
        | Name     | Optional | Description        |
        | -------- | -------- | ------------------ |
        | authors  | true     | Default to `[pid]` |
        | title    | true     |                    |
        | summary  | true     |                    |
        | tags     | true     |                    |
        | contents | true     |                    |

- PATCH `/personas/:pid/items/:tid` - change a item of a persona

    - Url parameters
    
        | Name | Optional | Description                                                                                            |
        | ---- | -------- | ------------------------------------------------------------------------------------------------------ |
        | id   | true     | file id of items file, filling in to speed up search process, empty to search it from the persona file |

    - Body parameters
    
        | Name     | Optional |
        | -------- | -------- |
        | authors  | true     |
        | title    | true     |
        | summary  | true     |
        | tags     | true     |
        | contents | true     |

- DELETE `/personas/:pid/items/:tid` - delete a item of a persona

    - Url parameters
    
        | Name | Optional | Description   |
        | ---- | -------- | ------------- |
        | id   | true     | Same as PATCH |

### Links

- GET `/personas/:pid/links` - get items of a persona

  - Url parameters

      | Name | Optional | Description                                                               |
      | ---- | -------- | ------------------------------------------------------------------------- |
      | id   | true     | file id of items file, empty for returning the data from the persona file |

- POST `/personas/:pid/links` - add a link to a persona

    - Body parameters
    
        | Name     | Optional |
        | -------- | -------- |
        | name     | true     |
        | tags     | true     |
        | list     | true     |

- PATCH `/personas/:pid/links/:lid` - change a link of a persona

    - Body parameters
    
        | Name     | Optional |
        | -------- | -------- |
        | name     | true     |
        | tags     | true     |
        | list     | true     |

- DELETE `/personas/:pid/links/:lid` - delete a link of a persona
