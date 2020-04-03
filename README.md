# JetBrains ReSharper action for GitHub Actions

[JetBrains ReSharper][jetbrains] is the best thing to happen to .Net developers
since .Net. JetBrains Rider is my favourite IDE. When JetBrains released their
ReSharper command-line tools for Mac and Linux, I knew what I had to do: create
a wrapper for GitHub Actions.

Now you can enjoy all the linting and suggestions that ReSharper provides in your
IDEs in your GitHub pull requests. Screenshot of an example PR:

<img width="1214" alt="Example PR" src="https://user-images.githubusercontent.com/369053/78336879-d03a7100-75db-11ea-9af4-e7d8aedec623.png">

## Usage

Add the following to a workflow:

```yaml
name: ReSharper

on:
  pull_request: {}

jobs:
  resharper:
    runs-on: ubuntu-latest
    steps:
      - name: checkout
        uses: actions/checkout@v2
      
      - name: resharper
        uses: glassechidna/resharper-action@master
        with:
          solution: HelloWorld.sln
```

[jetbrains]: https://www.jetbrains.com/dotnet/
