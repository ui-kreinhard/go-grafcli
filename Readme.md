What's this?
======

I would like to setup a grafana instance fully automatic including dashboards and datasources etc with the use of puppet or ansible.

There are possible solutions like wizzy(https://github.com/grafana-wizzy/wizzy) which could do that. But I don't like to install a complete node environment for just a "simple" command line task.
So I've got the idea to implement it in go. Just one binary :)

What it currently can do
===
It can currently:
* Import dashboards to a json file
````
./go-grafcli -mode=import-dashboards -filename=/tmp/dashboards.json -username=admin -password=admin -baseUrl="http://localhost:3000"
````
* Import single dashboard to a json file
````
./go-grafcli -mode=import-single-dashboard -filename=/tmp/dashboards.json -username=admin -password=admin -baseUrl="http://localhost:3000" -dashboardUri=myAwesomedashboard
````
* Export dashboards from a json file
````
./go-grafcli -mode=export-dashboards -filename=/tmp/dashboards.json -username=admin -password=admin -baseUrl="http://localhost:3000"
````
* Export datasources from a json file
````
./go-grafcli -mode=export-datasources -filename=/tmp/datasources.json -username=admin -password=admin -baseUrl="http://localhost:3000"
````
* Import datasources from a json file
```
./go-grafcli -mode=import-datasources -filename=/tmp/datasources.json -username=admin -password=admin -baseUrl="http://localhost:3000"
```
* Change password of current user
````
./go-grafcli -mode=changePassword -username=admin -password=admin -baseUrl="http://localhost:3000" -newPassword=blablabla
````
Status
====
Most probably really buggy. If it works for you - I'm really suprised :)