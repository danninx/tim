# tim
for the past year I've been using a `$HOME/templates` folder to keep a number of templates for things like school reports, scripts, nix expressions, etc., but typing out the path each time is a little tedious, especially when I have to `<tab>` `<tab>` `<tab>` a ton to remember the filenames. I wrote this little go-script to set up an alias system for files, directories and git repositories (and because I wanted to look at [golang](https://go.dev/) a bit for the first time)

## Sources and Templating
A source can be a file, directory, or git repository (heads up if your directory contains a git repository, as that is copied over as well; may address in the future)

`tim` uses the directory `$HOME/.config/tim` to store template configurations and copies so that they can be easily referenced and used elsewhere

As a quick summary, to add a source, say a file "`Makefile`" to your sources with the alias **make**, you can use

```sh
tim add file make Makefile
```

To then use that source for a template, you can use the `tim plate` command (see what I did there?)

```sh 
tim plate make Makefile
```

<sup>If you so insist, you can use a more familiar `tim clone` command instead...</sup>

I've tested this a bit using a [report format I made for some lab classes this last semester](https://github.com/danninx/tu-report.sty):

```sh
tim add git report git@github.com:danninx/tu-report.sty.git
tim plate report .
```

## Installation and Prerequisites
If you (most likely) plan on using git repositories as a source, you will need to have `git` installed and setup (with whatever authorization you need, normally an SSH key) on your system. 

```sh
go install github.com/danninx/tim@latest
```

If you need to build from source, you will need `go` version 1.24+. After that, clone the repository and run `go build` and then copy the binary to your path, adding any permissions necessary.

### Nix Configuration 
If you use flakes for your system configuration, you can include `tim` as a flake input. Just add this to your `flake.nix`:

```nix
inputs.tim.url = "github:danninx/tim";
```

And make sure to include it in your `modules` section of your `nixosConfiguration`s:

```nix
nixosConfigurations.joes-desktop = nixpkgs.lib.nixosSystem {
    modules = [ 
        ./configuration.nix 
        inputs.tim.nixosModules.default
    ];
};
```

Then you can enable it by either putting it in your `environment.systemPackages`, or use the following:

```nix
programs.tim.enable = true;
```


## To-Do

- Fix this damn readme
- Check git sources upon adding for validity
- Template syncing
- Template editing
- Proper unit/end-to-end tests
- Jinja2/other engine CI for nix package definitions
- More configuration
