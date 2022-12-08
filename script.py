import subprocess

s = subprocess.check_output('docker-compose --env-file=.env up', shell=True)