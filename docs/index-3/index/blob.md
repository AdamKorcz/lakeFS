---
layout: default
title: Azure Blob Storage
description: >-
  This guide explains how to configure Azure Blob Storage as the underlying
  storage layer.
parent: Prepare Your Storage
grand_parent: Setup lakeFS
nav_order: 30
has_children: false
---

# Azure Blob Storage

[Create a container in Azure portal](https://docs.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-portal#create-a-container):

1. From the Azure portal, Storage Accounts, choose your account, then in the container tab click `+ Container`.
2. Make sure you block public access

## Authenticate with Secret Key

{: .no\_toc }

In case you want to use the secret key for authentication you will need to use the account key in the configuration Go to the `Access Keys` tab and click on `Show Keys` save the values under `Storage account name` and `Key` we will need them in the [installing lakeFS](https://github.com/treeverse/lakeFS/tree/edb733eaf01e4b78c3896e621a40cc0756aa3e41/docs/setup/storage/install.md) step

## Authenticate with Active Directory

{: .no\_toc }

In case you want your lakeFS Installation \(we will install in the next step\) to access this Container using Active Directory authentication, First go to the container you created in step 1.

* Go to `Access Control (IAM)`
* Go to the `Role assignments` tab
* Add the `Storage Blob Data Contributor` role to the Installation running lakeFS.

You are now ready to [create your first lakeFS repository](../create-repo.md).

