### aah Config loading test

```
$ aah rc cfgtest
---------------------------------------------------------------
             aah framework v0.12.2 (cli v0.13.0)
---------------------------------------------------------------
# Report improvements/bugs at https://aahframework.org/issues #

Loaded aah project file: /www/aahcfgtest/aah.project
Compile starts for 'aahcfgtest' [aahcfgtest]
Compile successful for 'aahcfgtest' [aahcfgtest]
Running application 'aahcfgtest' console command: 'cfgtest'
2019-01-22 17:08:35.877 INFO  from aah.conf
2019-01-22 17:08:35.877 INFO  main value
2019-01-22 17:08:35.877 INFO  from dev profile
2019-01-22 17:08:35.877 INFO  dev value
2019-01-22 17:08:35.877 INFO  final value
2019-01-22 17:08:35.877 INFO  dev value
```

```
$ aah rc cfgtest -e prod
---------------------------------------------------------------
             aah framework v0.12.2 (cli v0.13.0)
---------------------------------------------------------------
# Report improvements/bugs at https://aahframework.org/issues #

Loaded aah project file: /www/aahcfgtest/aah.project
Compile starts for 'aahcfgtest' [aahcfgtest]
Compile successful for 'aahcfgtest' [aahcfgtest]
Running application 'aahcfgtest' console command: 'cfgtest -e prod'
2019-01-22 17:08:42.309 INFO  from aah.conf
2019-01-22 17:08:42.309 INFO  main value
2019-01-22 17:08:42.309 INFO  from prod profile
2019-01-22 17:08:42.309 INFO  prod value
2019-01-22 17:08:42.309 INFO  final value
2019-01-22 17:08:42.309 INFO  prod value
```

```
$ aah rc cfgtest -e prod -c external.conf
---------------------------------------------------------------
             aah framework v0.12.2 (cli v0.13.0)
---------------------------------------------------------------
# Report improvements/bugs at https://aahframework.org/issues #

Loaded aah project file: /www/aahcfgtest/aah.project
Compile starts for 'aahcfgtest' [aahcfgtest]
Compile successful for 'aahcfgtest' [aahcfgtest]
Running application 'aahcfgtest' console command: 'cfgtest -e prod -c external.conf'
2019-01-22 17:08:54.023 INFO  from aah.conf
2019-01-22 17:08:54.023 INFO  main value
2019-01-22 17:08:54.023 INFO  from prod profile
2019-01-22 17:08:54.023 INFO  prod value
2019-01-22 17:08:54.023 INFO  final value
2019-01-22 17:08:54.023 INFO  external value
```
