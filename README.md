# gopherlol
> [bunnylol](https://www.quora.com/What-is-Facebooks-bunnylol) / [bunny1](http://www.bunny1.org) -like smart bookmarking tool, written in Go

```bash
go run . 
# then add `http://localhost:8080/?q=%s` as a search engine to your browser 
```

# DO NOT FORK FOR ONE-OFFS

OK, that's harsh, let me explain.

I'd like to include as much generically-helpful stuff as we all add to this repository, but we have
two kinds of additions:
 - things that are specific to a company and not useful outside that company; for example, adding
    verbs and terms such as "j" going to "http://my.company.exmaple.com/internal/jira?q="
 - things that support the above case, generic things that are useful across all environments and
    can help developer the above cases

So. *This second thing* -- generic extension of capability to the upstream that are not specific to
single companies or environments -- _please_ fork and extend this repo.  ...and let's merge back to
here.

Things that are specific to a company or environment -- *the first case* -- please dupe/fork/extend
the example http://github.com/chickenandpork/gopherlol-extend/.  It includes an example of how to
extend the command-set ("verbs") and simply include the upstream gopherlol code as a dependency.
If that's too bazel-ish for you, file a ticket there, I can help :)

## How to set as a browser search engine
- in Chrome: [Set your default search engine](https://support.google.com/chrome/answer/95426)
- Instructions for all major browsers: [How-To Geek: How to Add Any Search Engine to Your Web Browser](https://www.howtogeek.com/114176/how-to-easily-create-search-plugins-add-any-search-engine-to-your-browser/)

## FAQ
- What are the commands currently supported by gopherlol? => Query gopherlol for `help` or `list`,
   or see [commands/commands.go](commands/commands.go)
- Why would a company run such a service internally? => Read about
   [how Facebook uses it internally](http://www.ccheever.com/blog/?p=74)

