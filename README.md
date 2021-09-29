# Link - Micro Language


## Sample
```
[use console]
[set print [in console 'print']]

[repeat i 32 {
	[print "Hello world!"]
	[if [and [less i 12] [greater i 2]] {
		[print [str i]]
	}]
}]
```

## Types
* Command block
```
[commandName arg1 arg2 arg3]
```
* CodeBlock
```
{
	[command1 a b c]
	[command2 a b]
	[command3 a]
}
```
* Variable
```
VarName
```
* Number
```
123
-32
144.13
```
* Strings
```
`Hello!`
'Hi'
"Hi Sir"
```

---
# API
---

## Parse
```
```

## Run
```

```