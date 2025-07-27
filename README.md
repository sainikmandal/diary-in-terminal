# diary-in-terminal (be patient I am not that great in markdown)

## What is this shit?
Okay so this is a project I am doing/done for the boot.dev hackathon. This is a terminal diary if you can't tell already from the absurd and ugly repo name. It has four commands: okay I forgot lemme cat the main.go file real quick, okay so: entry, delete, view and exit. Believe me when I say this, I had great plan for this project, but as it happens with all my projects I got lazy and finished it as it is. Thank God I implemented encryption somehow, so that safe from FBI and stuff, man only if a certain Jeffry knew encryption, your president wouldn't have to be in this shitty position right now, jk jk. 

## how you install it?
run `go install github.com/sainikmandal/diary-in-terminal` hopefully it will install (let me know before installing so that I can pray for your success)
I strongly recommend you to alias it as something else, just run `alias diary='diary-in-terminal'`
I am stuck with an ugly name man, what do you want me to do? change the repo name? Change the local dir name then do some voodo git shit and then change all the imports and go mod? Nah I don't think so.
Also you probably should add it your path.
Run `echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc` (or ~/.zshrc if you use zsh) to add it to your path.
If you use fish as your shell instead of eating it, run `echo 'set -gx PATH $PATH (go env GOPATH)/bin' >> ~/.config/fish/config.fish` 

## 📄 License
MIT © 2025 (sainikmandal)
Very proprietary(yes I googled the spelling) stuff, if you install it, I'll send you a bill with many zeroes.

