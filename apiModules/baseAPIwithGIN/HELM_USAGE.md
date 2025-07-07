# Helm Usage Guide

## Install/Deploy the Chart

```sh
helm install <release-name> ./helm
```
- Replace `<release-name>` with your desired Helm release name.
- Run this command from the `baseAPIwithGIN` directory or provide the full path to the `helm` folder.

## Upgrade/Update the Release

```sh
helm upgrade <release-name> ./helm
```
- Use this command to apply changes after modifying the chart or values.

## Uninstall/Delete the Release

```sh
helm uninstall <release-name>
```
- This will remove all Kubernetes resources associated with the release.

---

## Additional Tips

- To override values, use `-f myvalues.yaml` or `--set key=value`.
- To see all releases:  
  ```sh
  helm list
  ```
- To check the status of a release:  
  ```sh
  helm status <release-name>
  ```
