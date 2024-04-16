# Prettyprint Log

Format json logs from Zerolog nicely

Install with `go install github.com/richard87/prettyprintlog@latest`

## Example:

```
prettyprintlog --help
# Usage of prettyprintlog:
#   -l, --level string       Level field name (default "level")
#   -m, --message string     Message field name (default "message")
#   -t, --timestamp string   Timestamp field name (default "time")

kubectl logs somepodwithjsonlogging -f | prettyprintlog
```
