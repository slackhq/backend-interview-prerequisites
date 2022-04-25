# Requirements

To check that your system can run the backend onsite exercise in PHP, run the `setup` script included in this directory (if you're on Windows, run `setup.ps1` instead).

We encourage you to use whatever setup is most comfortable for you, whether that's running the script locally, in Docker, or in a VM. As long as the `setup` (or `setup.ps1`) script runs successfully, you should be good to go.

## Dependencies

The PHP version of the backend onsite exercise doesn't have dependencies on any third-party PHP libraries. Instead, we simply require the following:

* A PHP version >= 7.
* [SQLite3](https://www.sqlite.org/index.html) (this comes pre-installed on most operating systems).
* If you're on a Windows machine, you'll need [PowerShell](https://docs.microsoft.com/en-us/powershell/). Once PowerShell is installed, you can run `pwsh setup.ps1` to test.

## Troubleshooting

If for some reason the `setup` script won't run successfully, try the following:

1. Ensure that your PHP version is >= 7.
1. Ensure that SQLite3 is installed.
1. Manually run `php src/server.php` in one tab and run tests against it (`php tests/test.php`) in another. This is a good approach if your machine can run PHP but can't run bash or PowerShell.

If you try the above and are still having trouble, please reach out to your recruiting coordinator and we'll do our best to assist you. If possible, please include relevant stacktraces and system details (OS and language versions, etc) in your message.