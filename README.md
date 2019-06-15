[![Build Status](https://travis-ci.com/harm-matthias-harms/rpm.svg?token=VdHPqtvZnsqSz7z9NJXz&branch=master)](https://travis-ci.com/harm-matthias-harms/rpm)
[![codecov](https://codecov.io/gh/harm-matthias-harms/rpm/branch/master/graph/badge.svg?token=pqYDv80hOb)](https://codecov.io/gh/harm-matthias-harms/rpm)
# Role play management
The role play management system is designed to support large scale medical exercises. The application connects high-quality medical cases with roleplayers and teams to play injects. The application exceeds the management part in terms of simple performance analysis for the injects played in an exercise.

## Install
### Requirements:
* docker-compose >= 3.0
* nginx or apache

### Installation
1. Create a folder to place the application
```bash
mkdir rpm
cd rpm/
```

2. Download docker-compose file
```bash
wget https://raw.githubusercontent.com/harm-matthias-harms/rpm/master/deploy/docker-compose.yml
```

3. Add your env variables to the docker-compose file

4. Start the server
```bash
docker-compose up -d
```

5. Add nginx or apache config to redirect to docker container. Examples are provided under `deploy/`.

6. Create a backup script. It can make use of the `mongo-backup.yml` as well as the `mongo-restore.yml` from `deploy/`. It could be favourable to compress the backup folder and to transfer it to an other server.

7. Enjoy

All examples can be found under `deploy/`

### Troubleshooting
* If nginx or apache can't connect to the docker container add 127.0.0.1: before the ports. It should fix the problem.
* Any other problems or something not working? Pleas open an issue. Be aware, that we can't provide support for your deployments.

## Options
We recommend using our docker-compose.yml. To set your own parameter just replace the values in the environment variables.

## Contribute
A contribution could be:
* Bugs
* Suggested enhancements
* Code contribution


### Bugs
Please open an issue for a bug: [New Issue](https://github.com/harm-matthias-harms/rpm/issues/new)
Tell us what kind of bug this is and how to trigger it.

### Suggested enhancements
If you seeing the need for a new feature please write an [issue](https://github.com/harm-matthias-harms/rpm/issues/new) and tell us about the feature.

### Code contribution
Help is warmly welcome. If you want to implement a feature or fix a bug please let us know. Please comment on the issue you are working on to minify the risk of doubled work. Please assign the PR to an issue. If all pipelines will pass, we will take a look and finish it together with you.