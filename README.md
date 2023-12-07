# kchange
KChange gets the change cause of the provided deployments or cronjobs. 

### Install
Clone this repo and then `go install` from it. 

### Usage
Deployments can be provided as a list of strings using the `-d` param. 
Cronjobs can be provided as a list of strings using the `-j` param.

### Example
```
╰─ kchange -d deployment1 -d deployment2 -j cronjob1
kubectl set image deployment deployment1 pod=repo/imagename:123456 --record=true
kubectl set image deployment deployment2 pod=repo/imagename:123456 --record=true
kubectl set image cronjob cronjob1 pod=repo/imagename:123456 --record=true
```