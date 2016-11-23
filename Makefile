
ifeq "$(CIRCLECI)" "true"
	BUILDINFO=
else
	BUILDINFO=-ldflags "-X main.stratuxVersion=`git describe --tags --abbrev=0` -X main.stratuxBuild=`git log -n 1 --pretty=%H`"
$(if $(GOROOT),,$(error GOROOT is not set!))
endif

all:
	make xdump978
	make xdump1090
	make xlinux-mpu9150
	make xgen_gdl90

xgen_gdl90:
	go get -t -d -v ./main ./test ./linux-mpu9150/mpu ./godump978 ./mpu6050 ./uatparse
	go build -v $(BUILDINFO) -p 3 main/gen_gdl90.go main/traffic.go main/ry835ai.go main/network.go main/export_datalog.go main/managementinterface.go main/sdr.go main/ping.go main/uibroadcast.go main/monotonic.go main/datalog.go main/tracklog.go main/equations.go

xdump1090:
	git submodule update --init
	cd dump1090 && make -j4

xdump978:
	cd dump978 && make -j4 lib
	sudo cp -f ./libdump978.so /usr/lib/libdump978.so

xlinux-mpu9150:
	git submodule update --init
	cd linux-mpu9150 && make -f Makefile-native-shared

.PHONY: test
test:
	make -C test	

www:
	cd web && make -j4

install:
	install -m 0755 gen_gdl90 /usr/bin/gen_gdl90
	install image/10-stratux.rules /etc/udev/rules.d/10-stratux.rules
	install image/99-uavionix.rules /etc/udev/rules.d/99-uavionix.rules
	rm -f /etc/init.d/stratux
	install -m 0644 __lib__systemd__system__stratux.service /lib/systemd/system/stratux.service
	install -m 744 __root__stratux-pre-start.sh /root/stratux-pre-start.sh
	ln -fs /lib/systemd/system/stratux.service /etc/systemd/system/multi-user.target.wants/stratux.service
	make www
	install -m 0755 dump1090/dump1090 /usr/bin/dump1090

clean:
	rm -f gen_gdl90 libdump978.so
	cd dump1090 && make clean
	cd dump978 && make clean
	rm -f linux-mpu9150/*.o linux-mpu9150/*.so
