# GoScope

A tool I made to quickly store bug bounty program scopes in a local sqlite3 database. Download or copy a Burpsuite configuration file from the bonty program page and save it
as a .json file. Source it using GoScope and it will parse the file, storing a program name, in-scope domains and out-of-scope domains to the database.

**I have only tested this with a few HackerOne Burpsuite configurations. I am uncertain if other platforms use the same format.**

## Disclaimer
Always double-check that the scope information in the database matches the listings on the bounty program page. I have found that some programs do *not* include out-of-scope domains
in the Burpsuite configuration files. These can be manually added in to the .json file and then run through GoScope.



## Usage
Due to the long urls these configuration files have, I find it easier just to open the link in a new tab, ctrl-a and copy all data, and then paste it into a .json file. The long links tend to lag my terminal.

GoScope is configured to utilize subcommands and flags.

E.x. goscope [command] \<flags>
```sh
goscope -h
```
This will display help for the tool. 
| Command          | Description                                                | Example                                                         |
| ---------------- | ---------------------------------------------------------- | --------------------------------------------------------------- |
| add              | Add a new bounty program                                   | goscope add -p example -b example.json                          |
| query            | Query the database and return inscope and outscope         | goscope query -p example                                        |
| remove           | Remove a program from the database                         | gorecon remove -p example                                       |
| pipe             | Output only wildcard domains and pipe to other enumeration tools such as assetfinder | goscope pipe -p example               |

| Flags            | Valid Commands         | Description                        | Example                                                        |    
|------------------|------------------------|------------------------------------|----------------------------------------------------------------|
| -a               | pipe                   | Output all wildcard domains for all programs in database | goscope pipe -a                          |
| -b               | add                    | Burpsuite config file              | goscope add -p example -b burp.json                            |
| -c               | all                    | Config file location (default $HOME/.goscope.yaml) | Can set default database name / location here  |
| -d               | all                    | Specify database name and location (default ./scope.db) | goscope add -d example.db                 |
| -h               | all                    | See help for any command           | goscope add -h                                                 |
| -l               | query                  | List all programs in database      | goscope query -l                                               |
| -p               | all                    | Set the name of the bounty program | goscope query -p example                                       |


## Installation

GoScope requires sqlite3 for the database

```sh
▶ sudo apt install sqlite3
```

GoScope requires **go1.17+** to install successfully. Run the following command to get the repo -

```sh
▶ GO111MODULE=on go get -v github.com/d82r/goscope
```

## Running GoScope

Add a new bounty program scope to the database.
```sh
▶ goscope -p example -b example.json 
```

Query an existing program

```sh
▶ goscope query -p example
```

Remove a program from the database

```sh
▶ goscope remove -p example
```

Output program wildcard domains (*.example.com) to stdout as example.com so it can be pipe to tools such as assetfinder or subfinder
```sh
▶ goscope pipe -a
```
