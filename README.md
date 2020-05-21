## MyLabDiscoveries Client
# Introduction
In the first phase of development, we are going to analyse 96 samples (6 dyes
each) and show the results to the end user. This results should be shown in tabular format
and graphical format. This machine will be industrial machine with Touch screen.


## Golang Boilerplate
We have used Golang boilerplate to kickstart any go api project.


### Testing

Run test locally
```
$ make test
```

### DB Support

For SQLite use following commands:
./mylabdiscoveries start dev
./mylabdiscoveries migrate dev
./mylabdiscoveries create_migrattion filename dev

For PostgreSQL use following commands
./mylabdiscoveries start prod
./mylabdiscoveries migrate prod
./mylabdiscoveries create_migrattion filename prod
