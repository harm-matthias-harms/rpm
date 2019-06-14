# How to deploy rpm?

1. Create a folder to place the application
2. Create a docker-compose.yml like the one provided in this folder.
3. Run `docker-compose up -d`
4. Create a nginx or apache config file, dependent on your Server. Examples can be found in this folder.
5. Create a backup script. It can make use of the mongo-backup.yml as well as the mongo-restore.yml. It could be favourable to compress the backup folder and to transfer it to an other server.
6. Enjoy!

## Troubleshooting
* If nginx or apache can't connect to the docker container add `127.0.0.1:` before the ports. It should fix the problem.
* Any other problems or something not working? Pleas open an issue. Be aware, that we can't provide support for your deployments.