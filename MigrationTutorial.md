Th new version of the Cosmos SDK did substantial work to refactor the logic around chain genesis.

HEre is a guide to migrating your application to work with these changes. We will be migrating the original nameservice app from the tutorial.

### Steps

### /cmd/

- Replace GaiaInit with genutilcli
- Comment out the nameservice module from app.go, cmd/nsd and cmd/nscli

### /app.go

- Incorporate genutil and genaccounts
- Incorporate sdk.ModuleBasicManager

Lets get the app working with standard modules and the new version of the SDK

### MOdule Manager
