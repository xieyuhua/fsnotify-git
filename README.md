<!--
 * @Descripttion: 
 * @version: 
 * @Author: seaslog
 * @Date: 2022-03-04 15:26:58
 * @LastEditors: 谢余华
 * @LastEditTime: 2022-03-04 15:27:34
-->

# [Supported]

 - [x] windows 
 - [x] macos 
 - [x] linux

# [Usage]

```
1.fsnotifyGit will watch the filepath
2.if the filepath has some change, auto run the git command 'git add . && git commit && git push'

input the path you want to watch as the flag

Usage:
  fsnotifyGit [flags]

Flags:
  -h, --help            help for fsnotifyGit
  -p, --path string     input the path you want to watch as the flag
  -c, --pushCycle int   git push once time each %n seconds; the default is 5s (default 5)
```

# [Tips]

The filepath should be the git repository.

And the .git directory will be ignored.


