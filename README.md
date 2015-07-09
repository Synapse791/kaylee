# Kaylee

Use JSON config to find and replace inside multiple files.

---

### Install
To install Kaylee, go to the [releases](https://github.com/Synapse791/kaylee/releases) page and download the latest binary. For Ubuntu, run the following commands to install kaylee to the system:
```
chmod 755 kaylee
sudo mv kaylee /usr/bin/kaylee
```

This will enable kaylee to be run from anywhere in the system.

### Usage
Kaylee works by accepting a json config string. The easiest way of maintaining this JSON is to have it in a file and then cat it into a string argument. See the example config file below.

```
kaylee -c "`cat /home/user/kaylee.json`"
```

##### Flags
```
-c|config  - JSON config string 
-e|example - print example config
-h|help    - print help information
-v|verbose - enable verbose log output
```

### Configuration
Config is stored as JSON as an array of objects. Each object has a path and patterns. The path is the path of the file to edit and patterns is another object containing key value pairs where the key is the find string and the value is the replace string.

```json
[
  {
    "path"     : "",
    "patterns" : {
      "findme"     : "putme",
      "replace_me" : "with_me"
    }
  }
]
```