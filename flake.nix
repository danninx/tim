{
  description = "template management tool";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
  flake-utils.lib.eachDefaultSystem (system:
  let
    pkgs = import nixpkgs { inherit system; };
  in
  {
    devShells.development = pkgs.mkShell {
      buildInputs = with pkgs; [
        go 
      ];

      env = {
        STARSHIP_NIX_SHELL_NAME = "tim";
      };
    };

    packages.default = pkgs.buildGo124Module {
      pname = "tim";
      version = "0.3.0";

      src = ./.;
      vendorHash = null;

      meta = {
        description = "templating script for common sources";
        license = pkgs.lib.licenses.mit;
        homepage = "https://github.com/danninx/tim";
      };
    };
  });
}
