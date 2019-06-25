# GO Template Examples

## Shared templates

```console
$ go run ./shared
Template from a.gotmpl: a.gotmpl
Template from b.gotmpl: b.gotmpl
Template defined inside of b.gotmpl: bb
Output of Executing 'a'
-------------------------------------------

template "a.gotmpl" text


template "bb" text (inside of "b.gotmpl")


temple "b.gotmpl" text
```
