# Prism

A tool for deploying apps to multiple service prividers. This will work by looking at file types, build files, and configurations to make a graph of services. This graph will be stored in a database that can be visualized and build/deployments will be automatically kicked off.


## Example Usage

```bash
$ prism deploy

Prisma will deploy:
    Terraform       (2)
    AWS Lambda      (3)

Deploying services ...

```
