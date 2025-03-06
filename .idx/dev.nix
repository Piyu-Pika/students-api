# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  channel = "stable-23.11";

  packages = [
    pkgs.go
    pkgs.gopls
    pkgs.air
    pkgs.delve
    pkgs.gcc
  ];

  env = {
    CGO_ENABLED = "1";
    GOPATH = "/workspace";
  };
  idx = {
    extensions = [
      "golang.Go"
      "ms-vscode.Go"
    ];

    # Enable previews
    previews = {
      enable = true;
      previews = {
        # web = {
        #   # Example: run "npm run dev" with PORT set to IDX's defined port for previews,
        #   # and show it in IDX's web preview panel
        #   command = ["npm" "run" "dev"];
        #   manager = "web";
        #   env = {
        #     # Environment variables to set for your server
        #     PORT = "$PORT";
        #   };
        # };
      };
    };

    # Workspace lifecycle hooks
    workspace = {
      onCreate = {
       go-mod = "go mod tidy";
      };
      onStart = {
        air = "air -c .air.toml";
      };
    };
  };
}
