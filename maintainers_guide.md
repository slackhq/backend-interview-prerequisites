# Maintainers Guide

This document describes tools, tasks and workflow that one needs to be familiar with in order to effectively maintain this project. If you do not plan to make changes to this project, this guide is not for you.

## Tasks

### Testing

Tests for each language are run using the `setup` script found in each language-level directory. Tests are also run by CI when pushing to a branch.

### Releasing

Candidates should use the latest version of `backend-interview-prerequisites`, found [here](https://github.com/slackhq/backend-interview-prerequisites/releases/latest). To release a new version of `backend-interview-prerequisites`, simply cut a new release on GitHub.

## Workflow

### Versioning and Tags

`backend-interview-prerequisites` uses [semver](https://semver.org/). Each new release should increment the PATCH version of the project. We are not yet sure what constitutes minor and major releases for `backend-interview-prerequisites`.

### Issue Management

Labels are used to run issues through an organized workflow. Here are the basic definitions:

*  `bug`: A confirmed bug report. A bug is considered confirmed when reproduction steps have been
   documented and the issue has been reproduced.
*  `enhancement`: A feature request for something this package might not already do.
*  `docs`: An issue that is purely about documentation work.
*  `tests`: An issue that is purely about testing work.
*  `needs feedback`: An issue that may have claimed to be a bug but was not reproducible, or was otherwise missing some information.
*  `discussion`: An issue that is purely meant to hold a discussion. Typically the maintainers are looking for feedback in this issues.
*  `question`: An issue that is like a support request because the user's usage was not correct.
*  `semver:major|minor|patch`: Metadata about how resolving this issue would affect the version number.
*  `security`: An issue that has special consideration for security reasons.
*  `good first contribution`: An issue that has a well-defined relatively-small scope, with clear expectations. It helps when the testing approach is also known.
*  `duplicate`: An issue that is functionally the same as another issue. Apply this only if you've linked the other issue by number.

Issues are closed when a resolution has been reached. If for any reason a closed issue seems
relevant once again, reopening is great and better than creating a duplicate issue.

## Everything else

When in doubt, file an issue.