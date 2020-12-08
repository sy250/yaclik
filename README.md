# yaclik
Yet Another cli-kintone

## Version

0.2.0

## Downloads

Not yet.

## Usage

```text
Usage of ./yaclik:
  -a string
    	App ID (default "0")
  -d string
    	Domain name
  -n string
    	User login name
  -o string
    	Output format (default "csv")
  -p string
    	User login password
  -g string
    	Guest Space ID
```

## Examples

### Export fields from an app as CSV format
```
yaclik -a <APP_ID> -d <SUB_DOMAIN> -n <USER_ID> -p <USER_PASSWORD>
```

### Export fields from an app as JSON format
```
yaclik -a <APP_ID> -d <SUB_DOMAIN> -n <USER_ID> -p <USER_PASSWORD> -o json
```
