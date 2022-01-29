If you run the analysis package with:

```
go run ./analysis
```

it will print:

```
field.Name: id - field.Type: string
field.Name: name - field.Type: string        
field.Name: games - field.Type: *invalid type
```

What I don't understand is why I'm having `*invalid type` instead of `*Games`?
