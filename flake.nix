{
  description = "go devshell and package, created by scaffolder";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "go-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
            just
          ];
        };

        packages.slag = pkgs.buildGoModule {
          pname = "slag";
          version = "0.1.1";

          src = self;

          vendorHash = "sha256-7zrX8WG5u11xc6BDMn5wyHk5npn1IQoljZrR7GMD5CM=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A logging library for Go";
            license = licenses.mit;
            platforms = platforms.all;
          };
        };

        apps.slag = {
          type = "app";
          program = "${self.packages.${pkgs.stdenv.hostPlatform.system}.slag}/bin/slag";
        };
      });
}
