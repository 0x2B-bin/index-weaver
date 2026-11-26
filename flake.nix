{
  description = "Flake for Index Weaver";

  inputs = {
    nixpkgs.url = "https://channels.nixos.org/nixpkgs-unstable/nixexprs.tar.zst";
  };

  outputs = inputs: {
    packages = builtins.mapAttrs (system: pkgs: {
      default = pkgs.buildGoModule {
        pname = "index-weaver";
        version = "0.0.1";
        src = ./.;

        vendorHash = "sha256-n2/gw9dA9h/LRFnkmDSooTcwcNFXh3PGogkRD3DSk7s";
      };
    }) inputs.nixpkgs.legacyPackages;

    devShells = builtins.mapAttrs (system: pkgs: {
      default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
        ];
      };
    }) inputs.nixpkgs.legacyPackages;
  };
}
