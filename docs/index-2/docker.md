---
layout: default
title: With Docker
parent: Deploy lakeFS
description: >-
  This guide will help you deploy your production lakeFS environment with
  Docker.
nav_order: 50
---

# With Docker

{: .no\_toc }

## Database

{: .no\_toc }

lakeFS requires a PostgreSQL database to synchronize actions on your repositories. This section assumes you already have a PostgreSQL database accessible from where you intend to install lakeFS. Instructions for creating the database can be found on the deployment instructions for [AWS](aws.md#creating-the-database-on-aws-rds), [Azure](azure.md#creating-the-database-on-azure-database) and [GCP](gcp.md#creating-the-database-on-gcp-sql).

## Table of contents

{: .no\_toc .text-delta }

1. TOC 

   {:toc}

## Installing on Docker

To deploy using Docker, create a yaml configuration file. Here is a minimal example, but you can see the [reference](https://github.com/treeverse/lakeFS/tree/edb733eaf01e4b78c3896e621a40cc0756aa3e41/docs/reference/configuration.md#example-aws-deployment) for the full list of configurations.

* [AWS](docker.md#docker-tabs-1)
* [Google Cloud](docker.md#docker-tabs-2)
* [Microsoft Azure](docker.md#docker-tabs-3)

 {% include\_relative includes/aws-docker-config.md %} {% include\_relative includes/gcp-docker-config.md %} {% include\_relative includes/azure-docker-config.md %}

Save the configuration file locally as `lakefs-config.yaml` and run the following command:

```bash
docker run \
  --name lakefs \
  -p 8000:8000 \
  -v $(pwd)/lakefs-config.yaml:/home/lakefs/.lakefs.yaml \
  treeverse/lakefs:latest run
```

## Load balancing

You should have a load balancer direct requests to the lakeFS server. By default, lakeFS operates on port 8000, and exposes a `/_health` endpoint which you can use for health checks.

## DNS

As mentioned above, you should create 3 DNS records for lakeFS: 1. One record for the lakeFS API: `lakefs.example.com` 1. Two records for the S3-compatible API: `s3.lakefs.example.com` and `*.s3.lakefs.example.com`.

All records should point to your Load Balancer, preferably with a short TTL value.

## Next Steps

Your next step is to [prepare your storage](../index-3/index/). If you already have a storage bucket/container, you are ready to [create your first lakeFS repository](../index-3/create-repo.md).

