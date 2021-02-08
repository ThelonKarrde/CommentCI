# CommentCI

![CircleCI](https://img.shields.io/circleci/build/github/ThelonKarrde/CommentCI/master?style=plastic) ![Docker Image Version (latest by date)](https://img.shields.io/docker/v/rivshiell/commentci) ![GitHub](https://img.shields.io/github/license/thelonkarrde/commentci)


A tool to sent comments to Issues or Pull Requests in Github from CI tools.

---

### Usage
  
Required environment variables:
* `GITHUB_COMMENT_USER` - User from which name will be  sent comments to Github
* `GITHUB_COMMENT_TOKEN` - Personal Token for comment user. [Github doc](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token)

Available parameters:
```
usage: CommentCI [-h|--help] -o|--github-owner "<value>" -r|--github-repository
                 "<value>" [-s|--single-comment "<value>"] [-c|--codify]
                 [-f|--file "<value>" [-f|--file "<value>" ...]]
                 [-l|--file-comment "<value>" [-l|--file-comment "<value>"
                 ...]] -i|--issue-number <integer> [-m|--multi-comment]

                 Sent a comment to GitHub PR or Issue from your CI

Arguments:

  -h  --help               Print help information
  -o  --github-owner       Owner of the repository. User/Organisations.
  -r  --github-repository  Name of the github repository.
  -s  --single-comment     Single comment string to sent to GitHub.
  -c  --codify             Put comments to the Markdown code block.
  -f  --file               By repeating this flag you can specify multiple
                           files which content will be sent to comment.
  -l  --file-comment       By repeating this flag you can specify comments for
                           provided files in according order.
  -i  --issue-number       Number(id) of the Issue/PR to sent a comment.
  -m  --multi-comment      Put each file into a separate comment in GitHub..
                           Default: false
```

Usage examples:  
<br></br>
Single comment:  
```
GITHUB_COMMENT_USER=user GITHUB_COMMENT_TOKEN=xxx commentci -o repo_owner -r repo_name -i 2 -s "Single comment"
```  
<br></br>
Single file with a comment:  
```
GITHUB_COMMENT_USER=user GITHUB_COMMENT_TOKEN=xxx commentci -o repo_owner -r repo_name -i 2 -c -l "Comment to example file" -f ./example.txt
```  
<br></br>
Multiple files with comments:  
```
GITHUB_COMMENT_USER=user GITHUB_COMMENT_TOKEN=xxx commentci -o repo_owner -r repo_name -i 2 -c -l "Comment to example file 1" -f ./example_1.txt -l "Comment to example file 2" -f ./example_2.txt
```  