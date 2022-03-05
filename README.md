<!--
 * @Descripttion: 
 * @version: 
 * @Author: seaslog
 * @Date: 2022-03-04 15:26:58
 * @LastEditors: 谢余华
 * @LastEditTime: 2022-03-05 09:18:38
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
  -p, --path string     input the path you want to watch as the flag; the default is ./ (default ./)
  -c, --pushCycle int   git push once time each %n seconds; the default is 5s (default 5)
```


# [Tips]

The filepath should be the git repository.

And the .git directory will be ignored.


## Git：git pull每次都要求输入用户名和密码

全局配置

```
git config --global credential.helper store
git config --global -e
```

下次pull的时候输一次密码以后就不用输了

## 如果要清除用户名和密码
运行一下命令缓存输入的用户名和密码

```
git config --global credential.helper wincred
```

### 清除掉缓存在git中的用户名和密码
git credential-manager uninstall

```
vim .gitignore
/*
!.gitignore
!*.php
!*.tpl
!*.js
!*.css
!/fw/ 
/fw/*
!/fw/bin/
!/fw/sf/
```
说明：忽略全部内容，但是不忽略 .gitignore 文件、根目录下的 /fw/bin/ 和 /fw/sf/ 
目录；注意要先对bin/的父目录使用!规则，使其不被排除。

## 如果中途.gitignore文件要变更

```
git rm -r --cached .  //清除之前的文件
git add .
git commit -m 'update .gitignore'
```

## 回滚梳理一下git命令

```
git init //初始化仓库
git add .//添加暂存
git commit -m 'msg'//备注
git remote add xieyuhuagithub git@github.com:xieyuhua/git-test.git
git push -u xieyuhuagithub master //推送

//修改
git checkout .
git add .
git commit -m 'msg'//备注
git push -u xieyuhuagithub master //推送

//版本回滚，，，，暂存、本地
git reset --hard HEAD^ //回滚本地和暂存区，上一次提交
git reset --hard  554cce8514c59dd7552de3fea82 指定提交commit_id 

//将本地的状态回退到和远程的一样
git reset --hard xieyuhuagithub/master

//强推到远程：
git add .
git commit -m '回滚之后强制提交'
git push xieyuhuagithub HEAD --force  //强制执行回滚之后、之前提交内容都没有了

//修改仓库名称
git remote rm xieyuhuagithub  //删除仓库
git remote rename xieyuhuagithub new_name //修改仓库名
```