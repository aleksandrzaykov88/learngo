# Text Changer
Test task for Smartforce. This is a program to replace text in all files in a directory.

## Params
Three parameters are passed via the command line (separator - space):
* 1st: *path to file directory*
* 2nd: *text to be replaced*
* 3rd: *replacement text*

Or passed like environment variables and set up in Dockerfile/Doker run:

* -e **ROOT**=*path to file directory*
* -e **TOCHANGE**=*text to be replaced*
* -e **REPLACER**=*replacement text*

## Config
Structure of config file:

**log**
  **path:** *path to log-file* ***example: "c:/users/logs"***

**ext:** *array of valid file extensions* ***example: ".txt"***

**backup:** *flag indicating the need to create a backup* ***example: "false"***

## Docker
Docker commands
* docker build --tag text-changer .
* docker run --publish 8080:8080 -v **HOST_PATH_TO_ROOT**:**CONTAINER_PATH_TO_ROOT** s -e "ROOT=**CONTAINER_PATH_TO_ROOT**" text-changer
* docker cp ***CONTAINER_ID***:**CONTAINER_PATH_TO_LOGS** **HOST_PATH_TO_LOGS**