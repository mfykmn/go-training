# Infrastructure as Code

事前にすること
tfstate管理用のGCSバケットをGUIで作成

```bash
$ gcloud auth application-default login
$ terraform init
$ terraform plan
$ terraform apply
```

GAE defaultサービスはimportが必要

```bash
$ terraform import google_app_engine_application.default {project}
$ terraform import google_app_engine_application.worker worker.{project}
```
