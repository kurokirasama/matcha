{
  lib,
  buildGoModule,
  fetchFromGitHub,
  go,
  pkg-config,
  pcsclite,
  stdenv,
  apple-sdk_15,
  nix-update-script,
  versionCheckHook,
}:

buildGoModule (finalAttrs: {
  __structuredAttrs = true;

  pname = "matcha";
  version = "0.37.0";

  src = fetchFromGitHub {
    owner = "floatpane";
    repo = "matcha";
    tag = "v${finalAttrs.version}";
    hash = lib.fakeHash;
  };

  vendorHash = lib.fakeHash;
  proxyVendor = true;

  # Upstream pins `toolchain` to latest Go for dev builds (GOTOOLCHAIN=auto).
  # Nix sandbox sets GOTOOLCHAIN=local and can't download — rewrite to nixpkgs Go.
  postPatch = ''
    sed -i -E "s/^toolchain go[0-9.]+$/toolchain go${go.version}/" go.mod
  '';

  nativeBuildInputs = lib.optionals stdenv.hostPlatform.isLinux [
    pkg-config
  ];

  buildInputs =
    lib.optionals stdenv.hostPlatform.isLinux [ pcsclite ]
    ++ lib.optionals stdenv.hostPlatform.isDarwin [ apple-sdk_15 ];

  env.CGO_ENABLED = 1;

  ldflags = [
    "-s"
    "-w"
    "-X main.version=${finalAttrs.version}"
    "-X main.date=1970-01-01T00:00:00Z"
  ];

  nativeInstallCheckInputs = [ versionCheckHook ];
  doInstallCheck = true;
  versionCheckProgramArg = "--version";

  passthru.updateScript = nix-update-script { };

  meta = {
    description = "Beautiful and functional email client for the terminal";
    homepage = "https://matcha.email";
    changelog = "https://github.com/floatpane/matcha/releases/tag/v${finalAttrs.version}";
    license = lib.licenses.mit;
    mainProgram = "matcha";
    maintainers = with lib.maintainers; [ andrinoff ];
    platforms = lib.platforms.darwin ++ lib.platforms.linux;
  };
})
