# Introduction

This is just a personal repo to write some Go helper packages.


## Collaborating

One of the most effective ways to collaborate on GitHub is by using a forking/branching model as described in the [GitHub Guides](https://guides.github.com/).

### Setup

* [Fork](https://guides.github.com/activities/forking/) the main repository to your personal GitHub space
* Clone this new fork locally to your computer. Make sure you use the SSH URL, not the HTTPS URL. This will be your `origin` remote.
* Add an `upstream` remote whose URL is the SSH URL of the main repository - `git remote add upstream {{url}}`, replacing `{{url}}` with the main repo's URL.

### Divvying Up Work

* Create [issues](https://guides.github.com/features/issues/) in the main repository to describe what work needs to be done
* Assign the issue to an individual to work on and add it to the `To Do` pipeline in the ZenHub boards
* When you have started working on an issue, move it to the `In Progress` pipeline.

### Doing Work

* Each time you begin doing work on a new issue, check out the master branch by doing `git checkout master`. You will only be able to do this if you don't have any changes in your local codebase.
* Pull in the latest changes from upstream's master branch - `git pull upstream master`
* Create a new [branch](https://guides.github.com/introduction/flow/), named something relevant to the issue being worked on - `git checkout -b {{branch-name}}`, replacing `{{branch-name}}` with the name of your branch.
* Push your new branch to your origin remote - `git push -u origin {{branch-name}}`
* Add your commits and push to that branch - `git push origin {{branch-name}}`
* Issue a Pull Request in to the upstream repository when the work is done. Make sure the Pull Request comment includes a [keyword for closing issues](https://help.github.com/articles/closing-issues-via-commit-messages/) for closing the issue the work is for - `Resolves #42` (with `42` being the issue number)
* Once the Pull Request is merged, delete the local and remote branch you worked on - `git branch -d {{branch-name}}` for local, `git push origin :{{branch-name}}` for remote. **Important: Never Reuse A Branch After It Has Been Merged**



## sitespeedAnalyzer

[Sitespeed.io](https://www.sitespeed.io/) is an open source tool that helps you analyze your website speed and performance based on performance best practices and timing metrics.

After running Sitespeed.io, there would generate a json file which records the socre of each performance rule.

`sitespeedAnalyzer` aims to analyse these results and gets some helpful statics information.



