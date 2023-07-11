# Study-Timer

A Golang CLI app to manage your time efficiently. Based on the pomodoro technique, this allows the user to set intervals to work, then rest, ultimately creating studying cycles.

## Usage

1. Install Go
2. Clone this repo `$ git clone https://github.com/cglavin50/study-timer.git`
3. Run `$ go install`
4. From terminal, run `$ study timer {cycles} {work} {rest}`, where cycles is the number of repeated work : rest cycles, and work, rest are in minutes

Note: If on WSL with oh-my-zsh, make sure to add the following to your `~/.zshrc` file (fix found [here](https://stackoverflow.com/questions/36083542/error-command-not-found-after-installing-go-eval)):

```bash
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```
