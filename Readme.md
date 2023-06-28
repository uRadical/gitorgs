# gitorgs

This tool makes working with different GitHub organizations easier. It allows you to setup a git configuration per
organisation containing the ssh key, email address and gpg signing key to use for the organisation. 

Using this tool you can setup a structure like this. Your global `.gitconfig` will have the default settings. You can  
override these settings per organisation with a `.gitconfig` in the organisation root folder and adding the 
appropriate `includeIf` in yor global `.gitconfig`.

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
1. Navigate to the directory you would like to use as the root of your organisation.
1. run `gitorgs`
1. update `~/.gitconfig` with the `includeIf` that was output from `gitorgs` command
1. clone the repos belonging to this organisation

To set your GPG signing key you will need to run `gpg --list-keys` and copy/paste the key you wish to use. 