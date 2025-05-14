{
  description = "tim allows you to bookmark files, directories and git repositories for templating";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.11";
  };

  outputs = { self, nixpkgs }:
  let
    pkgs = import nixpkgs { system = "x86_64-linux"; };
  in
  {
    packages.x86_64-linux.tim = pkgs.buildGo124Module {
      pname = "tim";
      version = "unversioned";

      src = ./.;
      vendorHash = null;

      subPackages = [ "internals/cli" "internals/copy" "internals/timactions" "internals/timfile" "internals/timutils" ];
    };

    packages.x86_64-linux.default = self.packages.x86_64-linux.tim;

    devShell.x86_64-linux = pkgs.mkShell {
      packages = with pkgs; [
        go
      ];
    };

    apps.x86_64-linux.default = {
      type = "app";
      program = "${self.packages.x86_64-linux.default}/bin/tim";
    };
  };
}
