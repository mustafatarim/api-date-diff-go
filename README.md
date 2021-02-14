# Date Diff API (with golang)

This API simply calculates the difference between two given dates. It takes two dates as parameters and gives results in JSON format.

## Installation

```bash

cd api-date-diff

go run main.go

```

## Usage

```bash

[your_url]/api/dd?d1=[date1]&d2=[date2]

```

Example:

```bash

[your_url]/api/dd?d1=2017-10-21&d2=2021-01-24

```

Result:

```json

{
    "inYears":4,
    "inMonths":39,
    "inWeeks":170,
    "inDays":1191,
    "inHours":28584,
    "inMinutes":1715040,
    "inSeconds":102902400
}

```

## Demo

https://api-date-diff-go.vercel.app