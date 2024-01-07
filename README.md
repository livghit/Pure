# TemplKit - Many open source components, one kit

- TemplKit is an collection of OpenSource UI Libraries as Templ Components.

# How will it work?

```go
package main

import ("github.com/livghit/templkit")

func main(){
  component := templkit.Button()
  componrnt.Render(ctx ,os.Stdout)
}

// or use it directly in the component

templ Index(){
  @Layout(){
  <main>
  @Button("Click me man","submit")
</main>
}
}

```
