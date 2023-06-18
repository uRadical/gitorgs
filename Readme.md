# gitorgs

This tool makes working with different Git organizations easier. It allows you to setup a git configuration per
organisation containing the ssh key and email address to use for the organisation. 

Using this tool you can setup a structure like this. Your global `.gitconfig` will have the default settings. You can  
override these settings per organisation with a `.gitconfig` in the organisation root folder and adding the 
appropriate `includeIf` in yor global `.gitconfig`.

You can now set an organisation specific ssh key and email

```shell
~/
├── .gitconfig
└── code/
     ├── org1/
     │    ├── .gitconfig
     │    ├── repo1/
     │    │    └── .git
     │    └── repo2
     │         └── .git
     │
     └── org2/
          ├── .gitconfig
          ├── repo1/
          │    └── .git
          └── repo2/
               └── .git
```


## Installation

### Via go install
```bash
 go install uradical.io/go/gitorg@latest
```

## Usage
- Navigate to the directory you would like to use as the root of your organisation.
- run `gitorgs`
- update `~/.gitconfig` with the `includeIf` that was output from `gitorgs` command
- clone the repos belonging to this organisation