# World Generator

This is a tool and library to procedurally generate fantasy worlds.

## Configuration

The best way to configure the executable is to put a `world.yaml` file in your
working directory and pass it to the binary with the `-c` flag. For example:

`world -c world.yaml culture`

The file is structured like this:

```yaml
data_dir: <data directory>
save_dir: <save directory>
save_target: <either "DO" or "filesystem">
web_domain: <final web domain>
```

The following environment variables configure the executable and are required if
not specified in the config file:

- `WORLD_DATA_DIRECTORY`: The absolute directory that contains data files for the application.
- `WORLD_SAVE_DIRECTORY`: The absolute directory where files will be written to.
- `WORLD_SAVE_TARGET`: Either `DO` to save to Digital Ocean or `filesystem` to save to the local file system.
- `WORLD_WEB_DOMAIN`: The domain name (without protocol) remote images will serve from. For example, `www.world.com`.

## Building

Run `./build.sh`. The binary will be `dist/world`.

## Running in Development

Run `./run.sh <command>`. It will set sane defaults and use the local filesystem for saving files.
