## Requirements

To check that your system can run the backend onsite exercise in Python, run the
`setup.py` script included in this directory.

```bash
python ./setup.py
```

We encourage you to use whatever setup is most comfortable for you, whether
that's running the script locally, in Docker, or in a VM. As long as the `setup.py`
script runs successfully, you should be good to go.

## Dependencies

The exercise has been tested against all version of Python under `bugfix` and `security` maintenance status. See https://www.python.org/downloads/

In addition, it uses:

- [SQLite3](https://www.sqlite.org/index.html) via [`sqlite3`](https://docs.python.org/3/library/sqlite3.html)
  - `sqlite3` comes preinstalled as part of the Python Standard Library
- [`socketserver`](https://docs.python.org/3/library/socketserver.html)
  - `socketserver` comes preinstalled as part of the Python Standard Library

## Troubleshooting

If for some reason the `setup.py` script won't run successfully, please file an issue against this repository containing the details of your operating system and any relevant stacktraces or error messages.
