go build;
service api stop;
cp api.service /lib/systemd/system/api.service;
chmod 755 /lib/systemd/system/api.service;
service api start;