# Contributing

Want to do some thing fun with the project? To contribute, you may 

## Make an issue
go to [issue area](https://github.com/qzyse2017/bleBird/issues)

## Pull Request
1. fork the repository
2. clone fork to local storage
3. branch

``` 
cd $working_dir/bleBird
git fetch upstream
git checkout master
git rebase upstream/master
git checkout -b myfeature
```

```myfeature``` here had better describe the feature you want to add.

4. add something new on this branch

5. keep your branch in sync

```
# While on your myfeature branch
git fetch upstream
git rebase upstream/master
```

6. commit your changes 

```
git commit
```

7. push commits 

```
git push -f ${your_remote_name} myfeature
```

8. create a new pull request

Visit your fork at https://github.com/$user/bleBird
Click the Compare & Pull Request button next to your myfeature branch.

