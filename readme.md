# GoBlender client

Sample code for client specific code for GoBlender.

## Setup

This code lives in `./client/clienthandlers`, and is its own 
git repository. When working in JetBrains Goland, you must create 
(or clone) a git repository in this location, and then add the directory
in Preferences -> Version Control.

Next, uncomment the three lines specified in run.sh to include
a git pull when running the app:

```
# comment next three lines out if not using custom client code
#cd ./client/clienthandlers
#git pull
#cd ../..
```

so that it looks like this:
```
# comment next three lines out if not using custom client code
cd ./client/clienthandlers
git pull
cd ../..
```
