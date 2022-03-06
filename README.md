# Pomodoro timer in Go

## Installation
1. Install `go`
2. Build
```bash
$ git clone https://github.com/Murtaza-Udaipurwala/pomodoro
$ cd pomodoro
$ make PREFIX=~/.local clean install
```

# Usage
1. Start pomodoro timer(default duration: 25m)
```bash
$ pom start
```

2. Start pomodoro timer with a different duration
```bash
$ pom start 45 # starts a 45m pomodoro session
```

3. Stop pomodoro session
```bash
$ pom stop
```

4. Extend pomodoro sesion
```bash
$ pom add 10 # adds 10 minutes more to the pomodoro session
```

5. Help
```bash
$ pom help
```

6. Bash completion
Add `complete -C pom pom` in your .bashrc to enable completion

## Add Pomodoro timer to TMUX
```bash
set -g status-interval 1
set -g status-right "#(pom)"
```
