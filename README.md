# zabbix.ecc-memory

Utility for monitoring ECC memory errors.

Features
--------

- discover system memory controllers and Chip-Select Rows (csrow)
- get stats for selected mc or csrow

Usage
-----

Discovery memory controllers in system:
```code
./ecc-memory discovery -type mc

[
    {
        "{#DEVICE.NAME}":"mc0"
    },
    {
        "{#DEVICE.NAME}":"mc1"
    }
]
```


Discovery csrow in system:
```code
./ecc-memory discovery -type csrow

[
    {
        "{#DEVICE.NAME}":"mc0.csrow0"
    },
    {
        "{#DEVICE.NAME}":"mc0.csrow1"
    },
    ...
]
```

Get some memory controller statistic:
```code
./ecc-memory stats -type mc -name mc1
```

Get some csrow statistic:
```shell
./ecc-memory stats -type csrow -name mc0.csrow1 
```