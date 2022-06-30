# Rundeck Parser

[![Test Code](https://github.com/graytonio/mr-parser/actions/workflows/test_code.yml/badge.svg)](https://github.com/graytonio/mr-parser/actions/workflows/test_code.yml)

An easier way to extract errors from rundeck jobs.

### **Before**

```
15:32:45 [rundeck@runner 2][NORMAL] TASK [Fail if server version does not match Veritas version information] *******
15:32:45 [rundeck@runner 2][NORMAL] Wednesday 29 June 2022  10:32:45 -0500 (0:00:00.857)       0:00:13.338 ********
15:32:46 [rundeck@runner 2][NORMAL] skipping: [server-0]
15:32:46 [rundeck@runner 2][NORMAL] skipping: [server-1]
15:32:46 [rundeck@runner 2][NORMAL] skipping: [server-2]
15:32:46 [rundeck@runner 2][NORMAL] skipping: [server-3]
15:32:46 [rundeck@runner 2][NORMAL] skipping: [server-4]
15:32:46 [rundeck@runner 2][NORMAL]
15:32:46 [rundeck@runner 2][NORMAL] TASK [Fail if server hostname does not match Veritas hostname] *****************
15:32:46 [rundeck@runner 2]
.
.
.
*******************
15:38:53 [rundeck@runner 2][NORMAL] Wednesday 29 June 2022  10:38:53 -0500 (0:00:00.325)       0:06:21.271 ********
15:38:54 [rundeck@runner 2][NORMAL] ok: [server-0]
15:38:54 [rundeck@runner 2][NORMAL] ok: [server-1]
15:38:54 [rundeck@runner 2][NORMAL] ok: [server-2]
15:38:54 [rundeck@runner 2][NORMAL] skipping: [server-3]
15:38:54 [rundeck@runner 2][NORMAL] ok: [server-4]
15:38:54 [rundeck@runner 2][NORMAL]
```

### **After**

```
TASK [APPLY SQL - FAIL if SQL failed] ******************************************
    fatal: [server-2]: FAILED! => {"changed": false, "msg": "Traceback (most recent cal
l last):...

TASK [run script at the agencies] *********************************
    fatal: [server-4]: FAILED! => {"changed": true, "rc": 4}
```

## Usage

`mr-parser -i {input_file} -o {output_target}`

## Installation

The binary can be downloaded from the releases page for the latest stable version or you can compile the code yourself for the latest version.
