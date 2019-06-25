# GO Template Examples

## Shared templates

```console
$ go run ./shared
Template from a.gotmpl is named "a.gotmpl"
Template from b.gotmpl is named "b.gotmpl"
Template defined inside of b.gotmpl is named "bb"
Output of Executing 'a'
-------------------------------------------

template "a.gotmpl" text


template "bb" text (inside of "b.gotmpl")


template "b.gotmpl" text

Output of Executing 'a' the second time
-------------------------------------------

template "a.gotmpl" text


template "bb" text (inside of "b.gotmpl")


template "b.gotmpl" text

```

Note: Peek inside of ["./shared/templates/a.gotmpl"](https://github.com/freeformz/go-templates/blob/master/shared/templates/a.gotmpl)
