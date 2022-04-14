# Requirements

To check that your system can run the backend onsite exercise in Ruby, run the
`setup` script included in this directory.

We encourage you to use whatever setup is most comfortable for you, whether
that's running the script locally, in Docker, or in a VM. As long as the `setup`
script runs successfully, you should be good to go.


## Dependencies

The exercise has been tested against all versions of Ruby under normal
maintenance or security maintenance. See
https://www.ruby-lang.org/en/downloads/branches/

In addition, it uses:
* [SQLite3](https://www.sqlite.org/index.html) via [sqlite3-ruby](https://rubygems.org/gems/sqlite3/versions/1.3.11)
    * SQLite comes preinstalled on most operating systems
* [ActiveRecord](https://guides.rubyonrails.org/active_record_basics.html)
    * You are welcome to write database queries manually if you prefer
* [RSpec](https://rspec.info/)


## Troubleshooting

If for some reason the `setup` script won't run successfully, it will attempt to
provide advice. If all else fails, please file an issue against this repository
containing the details of your operating system and any relevant stacktraces or
error messages.
