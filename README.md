# What

This repo hold a variaty of methods for converting strings to something else.

# Types

## Data size

Convert strings of type `5MB` to:
- `DataSize(value int64, unit string)`
- To bytes (`int64`)

Example:

```
s := "5MB"
d, _ := datasize.FromString(s)
fmt.Println(d.String())

b := d.ToBytes()
fmt.Println(b, "Bytes")

B, _ := datasize.ToBytes(s)
fmt.Println(B, "Bytes")
```
