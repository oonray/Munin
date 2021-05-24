# Munin / Hugin

A Minecraft modpack  
**minecraft:** 1.16.5  
**forge:** 36.0.45

## Development Environment
Clone the repo to an empty folder inside the curse client profile folder.
Create temporary custom profile in the curse client with the correct minecraft and forge verisons. Copy minecraftinstance.json from the temporary profile the repo folder. The repo folder should now show up as a modpack profile in the curse launcher. Delete the temporary profile.

Modify and test the mods and config.

## Notes on Config
Mod config is divided into 3 categories

### Client
Client-side only configs.  
Located in config/.
### Common
Configs for both server and client but not synced automatically.  
Located in config/.
### Server
Per-world configs for server side, synced to client.  
Located in saves/world/serverconfig/. Copied from defaultconfigs/ on world creation.  
Change by modifying the config in saves/world/serverconfig. After confirming the changes, copy the updated file to defaultconfigs/. Delete the old file from serverconfig to update existing worlds.