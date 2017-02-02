Get `go1.7.5` from `go1.x` or get `go1.5.4` from `go1.5.x`.

Uses the same data as [Gimme](https://github.com/travis-ci/gimme).

This is not perfect, not tested excessively, but it suits our needs. 

Usage :

```go
	// fetch the current list of go versions
	list, err := version.GetVersionList()
	if err != nil {
		panic(err)
	}

	// find the latest
	fmt.Println(version.FindLatestVersion(list, "go1.5.x"))
```
