<h1>Public sonar</h1>

<h3>Project structure</h3>

1. /message
2. /parser
3. /query
4. /repository
5. /storage

<h4>/Message</h4>
This folder contains sturuct of messages and cover:

1. Parsing json
2. Creating Message struct

<h4>/Parser</h4>
This folder responsible for parsing file json or any formats 
also this package will use in parsers of message and query

```
func ParseJson(filepath string, output interface{}) error {
    file, err := os.Open(filepath)
    if err != nil {
    return ErrInvalidPath
    }
    defer file.Close()

	err = json.NewDecoder(file).Decode(&output)
	if err != nil {
		return ErrInvalidFormat
	}

	return err
}
```

How does it use?

```
func ParseJson(filepath string) ([]Query, error) {
	var output = make([]QueryDTO, 0, 100)
	err := parser.ParseJson(filepath, &output)
	if err != nil {
		return nil, err
	}
	
	.... any code

	queries := make([]Query, 0, len(output))
	for i := 0; i < len(output); i++ {
		dtoQuery := output[i]
		if dtoQuery.Queries != "" {
			queries = append(queries, Query{
				CaseID: dtoQuery.CaseID,
				Query:  dtoQuery.Queries,
			})
		}

		if dtoQuery.Query != "" {
			queries = append(queries, Query{
				CaseID: dtoQuery.CaseID,
				Query:  dtoQuery.Query,
			})
		}
	}

	return queries, nil
}
```

<h4>/Query</h4>
This folder responsible for parsing and building query. 
Every query has 4 components:

1. GROUP and looks like (word OR words)
2. AND 
3. OR
4. word or any words like this (real madrid AND goal)

For **collecting** query component I implement **design pattern** [Composite](https://golangbyexample.com/composite-design-pattern-golang/#:~:text=Composition%20design%20pattern%20is%20used,objects%20into%20a%20tree%20structure.)

```
type InterfaceNode interface {
    Search(func(string) bool) bool
    Add(node InterfaceNode)
    SetCondType(condType CondType)
    SetPhrase(string)
}

type Node struct {
    Phrase     string
    Conditions []InterfaceNode
    CondType   CondType
}
```

there are interface and struct of Node that implement design pattern, so main functions
are Search() and Add() because you must add components and walk by tree.

<h4>/repository</h4>
_Disclaimer - it is not best realization_

I tried to implement independent compare function that responsible only for compare data
and add to `Search()` argument that looks like `Search( func(string) bool ) bool`.
Now every node can use func that was passed through argument in any search nodes.

<h4>/storage</h4>
As usual every storage contains data about value files: logs, output etc.
There is no big logic.
