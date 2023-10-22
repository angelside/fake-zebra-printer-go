# Fake zebra printer

[![maintenance-status](https://img.shields.io/badge/maintenance-as--is-yellow.svg?style=for-the-badge)](https://gist.github.com/angelside/364976fbcf7001a5da7e79ad8ed91fec) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![Windows](https://img.shields.io/badge/Windows-0078D6.svg?style=for-the-badge&logo=Windows&logoColor=white)](https://www.microsoft.com)
[![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge&logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciICB2aWV3Qm94PSIwIDAgNDggNDgiIHdpZHRoPSI0OHB4IiBoZWlnaHQ9IjQ4cHgiPjxwYXRoIGZpbGw9IiM0Y2FmNTAiIGQ9Ik0yNCw1QzEzLjUsNSw1LDEzLjYsNSwyNC4xYzAsOC4yLDUuMSwxNS4xLDEyLjMsMTcuOWw0LjItMTEuNUMxOC44LDI5LjUsMTcsMjcsMTcsMjQgYzAtMy45LDMuMS03LDctN3M3LDMuMSw3LDdjMCwzLTEuOCw1LjUtNC41LDYuNUwzMC43LDQyQzM3LjksMzkuMiw0MywzMi4zLDQzLDI0LjFDNDMsMTMuNiwzNC41LDUsMjQsNXoiLz48cGF0aCBmaWxsPSIjMmU3ZDMyIiBkPSJNMTcuOSw0My4zbC0wLjktMC40QzkuMiw0MCw0LDMyLjQsNCwyNC4xQzQsMTMsMTMsNCwyNCw0YzExLDAsMjAsOSwyMCwyMC4xIGMwLDguMy01LjIsMTUuOS0xMi45LDE4LjhsLTAuOSwwLjRsLTQuOC0xMy4zbDAuOS0wLjRjMi4zLTAuOSwzLjgtMy4xLDMuOC01LjZjMC0zLjMtMi43LTYtNi02cy02LDIuNy02LDZjMCwyLjUsMS41LDQuNywzLjgsNS42IGwwLjksMC40TDE3LjksNDMuM3ogTTI0LDZDMTQuMSw2LDYsMTQuMSw2LDI0LjFjMCw3LjEsNC4zLDEzLjcsMTAuNywxNi41bDMuNS05LjZDMTcuNiwyOS43LDE2LDI3LDE2LDI0YzAtNC40LDMuNi04LDgtOCBzOCwzLjYsOCw4YzAsMy0xLjYsNS43LTQuMiw3bDMuNSw5LjZDMzcuNywzNy44LDQyLDMxLjMsNDIsMjQuMUM0MiwxNC4xLDMzLjksNiwyNCw2eiIvPjwvc3ZnPg==)](./LICENSE)

> The "fake zebra printer" is a Go CLI server that listens to 127.0.0.1:9100 and prints back some processed data from the connection. It can colourize incoming ZPL codes to some extent. It is coded to simulate network printing to Zebra printers, but it does not interpret ZPL codes.


## 📦 Installation

...

## 🚀 Compilation

> ⚠️ You must compile it yourself until any binary is released.

> ⚠️ Only tested on Windows.

Dependency for compilation
https://github.com/go-task/task

The following command will build the binaries and save them in the `./__dist/` directory with the name `fakeZebraPrinter.exe`:

```console
> git clone https://github.com/angelside/fake-zebra-printer-go.git
> cd fake-zebra-printer-go
> task build
```

## 🔨 Usage

> ⚠️ It was only tested for use on the same machine; use over the network was not tested.

Windows
```powershell
./fakeZebraPrinter.exe
```

### 📋 Sample results

![fake-zebra-printer_dark](https://github.com/angelside/fake-zebra-printer-go/assets/7515/b82e2d6d-b8e4-459f-a416-f5804c516257)

## 💥 Features

- It just works!

## 🤝 Contributing

Before contributing issues or pull requests, could you review the [Contributing Guidelines](./.github/CONTRIBUTING.md) first?

## 💬 Questions?

Feel free to [open an issue](http://github.com/angelside/fake-zebra-printer-go/issues/new).

## 🤩 Support

💙 If you like this project, give it a ⭐ and share it with friends!

## 🏛️ License

This project is open-sourced software licensed under the [MIT license](./LICENSE).
