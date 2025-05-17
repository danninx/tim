# tim
for the past year I've been using a `/home/<user>/templates` folder to keep a number of templates for things like school reports, scripts, nix expressions, etc., but typing out the path each time is a little tedious, especially when I have to `<tab>` `<tab>` `<tab>` a ton to remember the filenames. I wrote this little go-script to set up an alias system for files, directories and git repositories (and because I wanted to look at [golang](https://go.dev/) a bit for the first time)

## Sources and Templating
A source can be a file, directory, or git repository (heads up if your directory contains a git repository, as that is copied over as well; may address in the future)

`tim` keeps track of sources in a `$HOME/.timfile`, which also means you can port your template aliases (mainly git, since everything else is local) across different machines if needed.

As a quick summary, to add a source, say a file "`./Makefile`" to your sources with the alias **make**, you can use

```sh
tim add make -f ./Makefile
```

You can use the `-d` flag for directories, otherwise tim will default to a git repository as a source.

To then use that source for a template, you can use the `tim plate` command (see what I did there?)

```sh 
tim plate make ./Makefile
```

<sup>If you so insist, you can use the nasty `tim copy` command instead...</sup>

I've tested this a bit using a [report format I made for some lab classes this last semester](https://github.com/danninx/tu-report.sty):

```sh
tim add report git@github.com:danninx/tu-report.sty.git
tim plate report .
```

## Installation and Prerequisites
If you (most likely) plan on using git repositories as a source, you will need to have `git` installed and setup (with whatever authorization you need, normally an SSH key) on your system. I provide the `linux/amd64` build in the repository, but if that doesn't fit your system you will need to build from source. 

```sh
git clone https://github.com/danninx/tim.git && cd tim
```

If you need to build from source, you will need `go` version 1.24+. After that just run `go build` and then you can use the instructions above

```sh
cp ./tim /usr/local/bin/tim         # wherever 
sudo chmod +x /usr/local/bin/tim    # make tim executable
```

I might setup a github action soon to build for more common systems

<sup> for those that believe in church of nix, you can grab the latest commit info using `nix-prefetch-git git@github.com:danninx/tim` import it explicitly to include in your packages </sup>

```
let
 tim = (pkgs.callPackage (pkgs.fetchFromGitHub {
    owner = "danninx";
    repo = "tim";
    rev = <REV>;
    sha256 = <SHA256>;
  }) {});
in
```

<sup> I've probably broken about a hundred nix standards in that expression, but it works... currently looking into better ways</sup>

## To-Do
*(I'm gonna move this to github issues or a trello in a bit)*

- Check git sources upon adding for validity
- Clean up the CLI parser so I can re-use it for some other stuff (I must resist [cobra](https://github.com/spf13/cobra))
- Add flags to help menu
