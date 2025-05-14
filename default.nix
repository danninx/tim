{ pkgs ? import <nixpkgs> { } }:

with pkgs;
let
  version = "0.0.1";
in
buildGo124Module {
  pname = "tim";
  inherit version;

  src = ./.;

  vendorHash = null;
}
