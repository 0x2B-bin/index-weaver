{
  description = "Flake for Index Weaver";

  inputs = {
    nixpkgs.url = "https://channels.nixos.org/nixpkgs-unstable/nixexprs.tar.zst";
  };

  outputs = inputs: {
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
