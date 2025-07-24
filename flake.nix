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

      src = pkgs.fetchFromGitHub {
        owner = "danninx";
        repo = "tim";
        rev = "fd183ede4038b82822c611e9d8687124d95424ed";
        hash = "sha256-IjgDz7wBzDLK7ae0HNfrOOUuup3xTA4LA1GAFx4eGOs=";
      };

      vendorHash = "sha256-GkwY1Y8n7vOJ2VFMjZP3Aew65HIPxQ/hb2eY2wq7rmE=";

      meta = {
        description = "templating script for common sources";
        license = pkgs.lib.licenses.mit;
        homepage = "https://github.com/danninx/tim";
      };
    };
  });
}
