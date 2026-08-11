# 001-file-statistics

#### Project overview

This is a GO project for learning and practicing go functionalities for build project. It is used to get the statistics of a directory, which includes
1. No. of files
2. No. of directories without root directory
3. Largest file name and size
4. Get count by the file extensions, if a file has no extension mark it as no-extension

#### How to run it

```
run go ./cmd/main.go "root path"
```

#### Pending works

##### writing test cases

```
Test 1
normal directory
→ files, directories, size, largest, extensions

Test 2
empty directory
→ zero files

Test 3
file without extension
→ "no-extension"

Test 4
invalid directory
→ error
```