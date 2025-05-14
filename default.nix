let
  pkgs = import <nixpkgs> { };
in
  pkgs.buildGo124Module {
    pname = "tim";
    version = "0.0.1";
    src = ./.;
    vendorHash = null;
  }
