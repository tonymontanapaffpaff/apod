## Description
A simple golang web application which consumes data from the [APOD API](https://api.nasa.gov/). It also provides a small API with the ability to save and filter data.
## Paths table
<table>
<tr>
<td>Path</td>
<td>Method</td>
<td>Description</td>
<td>Body example</td>
</tr>
<tr>
<td>/picture</td>
<td>GET</td>
<td>Get picture of the day</td>
<td>

```json
{
  "copyright":"ALSJ",
  "date":"2022-07-30",
  "explanation":"Get out your red/blue glasses...",
  "hdurl":"https://apod.nasa.gov/apod/image/2207/AS11JK44-6633-34.jpg",
  "media_type":"image",
  "service_version":"v1",
  "title":"The Eagle Rises",
  "url":"https://apod.nasa.gov/apod/image/2207/AS11JK44-6633-34_1100px.jpg"
}
```

</td>
</tr>
<tr>
<td>/album</td>
<td>GET</td>
<td>Get all saved pictures</td>
<td>
  
```json
[
    {
        "Id": 1,
        "RequestedAt": "2022-07-30T00:00:00Z",
        "Title": "The Eagle Rises",
        "Url": "https://apod.nasa.gov/apod/image/2207/AS11JK44-6633-34_1100px.jpg"
    },
    {
        "Id": 2,
        "RequestedAt": "2022-07-30T00:00:00Z",
        "Title": "SOFIA's Southern Lights",
        "Url": "https://apod.nasa.gov/apod/image/2207/ASC05954-Edit1024.jpg"
    }
]
```

</td>
</tr>
<tr>
<td>/album?date=[DATE]</td>
<td>GET</td>
<td>Get pictures saved on a specified day</td>
<td>
  
```json
[
    {
        "Id": 1,
        "Title": "The Eagle Rises",
        "Url": "https://apod.nasa.gov/apod/image/2207/AS11JK44-6633-34_1100px.jpg"
    },
    {
        "Id": 2,
        "Title": "SOFIA's Southern Lights",
        "Url": "https://apod.nasa.gov/apod/image/2207/ASC05954-Edit1024.jpg"
    }
]
```

</td>
</tr>
</table>

In order to save the image along with all its meta-data, you need to add the `?saved=true` argument to the `/picture` route.  
You can also pass `api_key` as a parameter. The `DEMO_KEY` is used by default.

## How to run
Using Make:
```
make up
```
