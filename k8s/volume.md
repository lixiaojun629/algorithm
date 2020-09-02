## Volume

创建和销毁的生命周期和Pod保持一致

Pod的 .spec.volumes字段说明了拥有哪些volume
.spec.containers[*].volumeMounts字段指定了某volume被挂载到容器镜像的某个目录


来源
https://kubernetes.io/docs/concepts/storage/volumes/


## PV(Perisit Volume)

删除保护，正在使用中的PVC和PV的删除行为会延迟到不再被使用的时候再删除
