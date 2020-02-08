# GoBlender client

Sample code for client specific code for GoBlender.

## Setup

First, install [goblender](https://github.com/tsawler/goblender).

All client specific code lives in `./client/clienthandlers`, and is its own 
git repository. When working in JetBrains Goland, you must create 
(or clone) a git repository in this location, and then add the directory
in Preferences -> Version Control.

# Updating on server
Change  `update.sh` to enable git pull of client:

```
# uncomment if using custom client code
#cd ./client/clienthandlers
#git pull
#cd ../..
```

After changing, it should look like this:

```
# uncomment if using custom client code
cd ./client/clienthandlers
git pull
cd ../..
```

