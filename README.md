# b2j

## What?

This converts bytes to json.

## Why?

When using gomocks, or even accidentally printing out the bytes from json, I get tired of seeing things like this:

```
missing call(s) to *mocks.MockThing.DoSomething(is equal to [123 34 104 101 108 108 111 34 58 34 119 111 114 108 100 34 125] (json.RawMessage))
```

And not knowing what it is actually looking for.

## How?

Just install and run:

```
b2j [123 34 104 101 108 108 111 34 58 34 119 111 114 108 100 34 125]
```

(works with or without brackets)

Output:
```
{
    "hello": "world"
}
```

Want it minified? 

```
b2j -m [123 34 104 101 108 108 111 34 58 34 119 111 114 108 100 34 125]
```

Output:
```
{"hello":"world"}
```
