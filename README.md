accentify
=========
Given a markdown file, Accentify will convert the spellings either from American English to British English or vice versa. 

```
go build
./accentify input.md uk
## uk = American -> British
## us = British -> American
```


Example
=======
`./accentify test.md us` will turn

```
favourite ice cream.
```

to 

```
favorite ice cream.
```