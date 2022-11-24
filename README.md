# dakill
My company has a process that monitors free disk space and blocks login sessions if there is less than 27 GB of free space.  I have identified which process this is, and have written a script to kill it. This will prevent it from annoying me any further.

## To compile
```
GOOS=windows GOARCH=386 go build main.go
mv main.exe dakill.exe
```

## Usage
Just run the `dakill.exe` after it, and it will monitor running processes every second and kill, if found `DiskAlert.exe`

Enjoy!
