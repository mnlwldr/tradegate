# whats-this

You can use this to retrieve the current information of an stock from Tradegate

## How to use it
```go
import "https://github.com/mnlwldr/tadegate"
```

## Data
`{Bid:0,292 Ask:0,317 Bidsize:1100 Asksize:1000 Delta: Stueck:600 Umsatz:175 Avg:0 Last:0,292 High:0,292 Low:0,292}`

## Usage
```go
response, err := tradegate.Get("DE000A1YCMM2") // SolarWorld AG
if err != nil {
	log.Println(err)
}
fmt.Printf("%s\n", response.Last)
```

## Output
```sh
0,292
```

[![Go Reference](https://pkg.go.dev/badge/github.com/mnlwldr/tradegate.svg)](https://pkg.go.dev/github.com/mnlwldr/tradegate)
