# mR Parser

[![Test Code](https://github.com/graytonio/mr-parser/actions/workflows/test_code.yml/badge.svg)](https://github.com/graytonio/mr-parser/actions/workflows/test_code.yml)

An easier way to extract errors from rundeck jobs.

### **Before**

```
15:32:45 [rundeck@zt-ops-0 2][NORMAL] TASK [Fail if server version does not match Veritas version information] *******
15:32:45 [rundeck@zt-ops-0 2][NORMAL] Wednesday 29 June 2022  10:32:45 -0500 (0:00:00.857)       0:00:13.338 ********
15:32:46 [rundeck@zt-ops-0 2][NORMAL] skipping: [sd-yanktonsiouxtribal-pd-prod-0]
15:32:46 [rundeck@zt-ops-0 2][NORMAL] skipping: [mo-webster-911-prod-0]
15:32:46 [rundeck@zt-ops-0 2][NORMAL] skipping: [sd-rosebud-pd-prod-1]
15:32:46 [rundeck@zt-ops-0 2][NORMAL] skipping: [la-grant-so-standby-0]
15:32:46 [rundeck@zt-ops-0 2][NORMAL] skipping: [la-grant-so-prod-1]
15:32:46 [rundeck@zt-ops-0 2][NORMAL]
15:32:46 [rundeck@zt-ops-0 2][NORMAL] TASK [Fail if server hostname does not match Veritas hostname] *****************
15:32:46 [rundeck@zt-ops-0 2]
.
.
.
*******************
15:38:53 [rundeck@zt-ops-0 2][NORMAL] Wednesday 29 June 2022  10:38:53 -0500 (0:00:00.325)       0:06:21.271 ********
15:38:54 [rundeck@zt-ops-0 2][NORMAL] ok: [sd-yanktonsiouxtribal-pd-prod-0]
15:38:54 [rundeck@zt-ops-0 2][NORMAL] ok: [mo-webster-911-prod-0]
15:38:54 [rundeck@zt-ops-0 2][NORMAL] ok: [sd-rosebud-pd-prod-1]
15:38:54 [rundeck@zt-ops-0 2][NORMAL] skipping: [la-grant-so-standby-0]
15:38:54 [rundeck@zt-ops-0 2][NORMAL] ok: [la-grant-so-prod-1]
15:38:54 [rundeck@zt-ops-0 2][NORMAL]
```

### **After**

```
TASK [APPLY SQL - FAIL if SQL failed] ******************************************
    fatal: [sd-rosebud-pd-prod-1]: FAILED! => {"changed": false, "msg": "Traceback (most recent cal
l last):\n  File \"/leds/bin/apply_sql\", line 403, in apply_patch\n    con.execute(\"SELECT util.a
pply_patch(:patchesid);\")\n  File \"/usr/local/lib/python2.7/dist-packages/zt/db/connections.py\",
 line 1108, in execute\n    self._execute(sql, vars=vars, dn=dn)\n  File \"/usr/local/lib/python2.7
/dist-packages/zt/db/connections.py\", line 1072, in _execute\n    cur.execute(sql)\n  File \"/usr/
local/lib/python2.7/dist-packages/zt/db/drivers/_psycopg.py\", line 304, in execute\n    return _cu
rsor.execute(self, query, vars)\nProgrammingError: relation \"hcp_address_responses\" already exist
s\nCONTEXT:  SQL statement \"CREATE TABLE hcp_address_responses (\n  cad_incidentsid BIGINT NOT NUL
L,\n  addressesid BIGINT,\n  hcp_address_requestsid BIGINT,\n  response_json JSONB\n ); \n\nALTER T
ABLE hcp_address_responses OWNER TO leds_cjis;\"\nPL/pgSQL function util.apply_patch(bigint) line 1
7 at EXECUTE\n\n\nError applying 2696: relation \"hcp_address_responses\" already exists\nCONTEXT:
 SQL statement \"CREATE TABLE hcp_address_responses (\n  cad_incidentsid BIGINT NOT NULL,\n  addres
sesid BIGINT,\n  hcp_address_requestsid BIGINT,\n  response_json JSONB\n ); \n\nALTER TABLE hcp_add
ress_responses OWNER TO leds_cjis;\"\nPL/pgSQL function util.apply_patch(bigint) line 17 at EXECUTE
"}

TASK [run set_permissions.yml at the agencies] *********************************
    fatal: [la-grant-so-prod-1]: FAILED! => {"changed": true, "rc": 4}
```

## Usage

`mr-parser -i {input_file} -o {output_target}`

## Installation

The binary can be downloaded from the releases page for the latest stable version or you can compile the code yourself for the latest version.
