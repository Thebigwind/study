package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type AccessLogDownloadParams struct {
	Type int
	Date string
	Hour string
	Path string
}

var LogTypeMap = map[int]string{
	1: "waiwang-lb",
	2: "neiwang-lb",
}

/*
curl -s 'http://10.4.1.20:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-00.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-0.log &
curl -s 'http://10.4.1.20:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-01.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-1.log &
curl -s 'http://10.4.1.25:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-02.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-2.log &
curl -s 'http://10.4.1.5:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-03.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-3.log &
curl -s 'http://10.4.1.15:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-04.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-4.log &
curl -s 'http://10.4.1.6:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-05.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-5.log &
curl -s 'http://10.4.1.3:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-06.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-6.log &
curl -s 'http://10.4.1.26:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-07.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-7.log &
curl -s 'http://10.4.1.27:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-08.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-8.log &
curl -s 'http://10.4.1.7:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-09.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-9.log &
curl -s 'http://10.4.1.20:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-10.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-10.log &
curl -s 'http://10.4.1.5:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-11.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-11.log &
curl -s 'http://10.4.1.0:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-12.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-12.log &
curl -s 'http://10.4.1.6:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-13.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-13.log &
curl -s 'http://10.4.1.7:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-14.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-14.log &
curl -s 'http://10.4.1.15:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-15.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-15.log &
curl -s 'http://10.4.1.22:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-16.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-16.log &
curl -s 'http://10.4.1.10:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-17.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-17.log &
curl -s 'http://10.4.1.20:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-18.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-18.log &
curl -s 'http://10.4.1.5:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-19.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-19.log &
curl -s 'http://10.4.1.23:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-20.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-20.log &
curl -s 'http://10.4.1.12:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-21.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-21.log &
curl -s 'http://10.4.1.6:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-22.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-22.log &
curl -s 'http://10.4.1.20:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-23.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-23.log &
*/

func DoDownloadLog(params AccessLogDownloadParams) error {

	if err := ParamFilter(params); err != nil {
		return err
	}

	//下载到目录目录
	err, path := GetPath(params.Path)
	if err != nil {
		return err
	}
	//如果目录不存在则创建目录
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if params.Hour == "" {
		for i := 0; i < 24; i++ {
			if err := DoDownloadHourLog(params, i, path); err != nil {
				return err
			}
		}
	} else {
		i, err := strconv.Atoi(params.Hour)
		if err != nil {
			return err
		}

		if err := DoDownloadHourLog(params, i, path); err != nil {
			return err
		}
	}

	return nil
}

func GetLogLocation(params AccessLogDownloadParams, i int) (Locations, error) {

	var loc Locations

	url := fmt.Sprintf("http://10.4.1.16:50070/webhdfs/v1/kafka-logs/%s/access_log/%s/access-%.2d.log.gz?op=GET_BLOCK_LOCATIONS", LogTypeMap[params.Type], params.Date, i)
	fmt.Printf("url:%s\n", url)

	var logLocation = LocatedBlocks{}
	if err := json.Unmarshal(GetUrl(url), &logLocation); err != nil {
		fmt.Printf("err:%s", err.Error())
		return loc, err
	}

	if len(logLocation.LastLocatedBlock.Locations) > 0 {
		loc = logLocation.LastLocatedBlock.Locations[0]
	} else {
		return loc, errors.New("没有取到location")
	}

	return loc, nil
}
func DownloadLog(params AccessLogDownloadParams, loc Locations, i int, path string) error {

	//command := "cd " + path + " && curl -s 'http://%s:%s/webhdfs/v1/kafka-logs/{}/access_log/{}/access-{:0>2d}.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o {}-{}-{}.log &"
	command := fmt.Sprintf("cd %s && curl -s 'http://%s:%s/webhdfs/v1/kafka-logs/%s/access_log/%s/access-%.2d.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o %s-%s-%d.log &",
		path, loc.IPAddr, loc.InfoPort, LogTypeMap[params.Type], params.Date, i, LogTypeMap[params.Type], params.Date, i)
	//fmt.Println(command)

	err, _ := ExecShell(command)
	if err != nil {
		return err
	}
	return nil
}

func DoDownloadHourLog(params AccessLogDownloadParams, i int, path string) error {

	loc, err := GetLogLocation(params, i)
	if err != nil {
		return err
	}

	if err := DownloadLog(params, loc, i, path); err != nil {
		return err
	}
	return nil
}

/*

	//response = requests.request("GET", url, headers={}, data={})
	//location = json.loads(response.text)['LocatedBlocks']['lastLocatedBlock']['locations'][0]
	//print("curl -s 'http://{}:{}/webhdfs/v1/kafka-logs/{}/access_log/{}/access-{:0>2d}.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o {}-{}-{}.log & ".format(
	//	location['ipAddr'], location['infoPort'], name, date, i, name, date, i))

{"LocatedBlocks":{"fileLength":393656398,"isLastBlockComplete":true,"isUnderConstruction":false,"lastLocatedBlock":{"block":{"blockId":1122203888,"blockPoolId":"BP-2028409268-10.4.1.16-1583828325109","generationStamp":2377258060,"numBytes":125220942},"blockToken":{"urlString":"AAAAAA"},"cachedLocations":[],"isCorrupt":false,"locations":[{"adminState":"NORMAL","blockPoolUsed":7844040235051,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":7844040235051,"hostName":"pg-kvm-hdfs-dn1015","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.15","ipcPort":50020,"lastBlockReportMonotonic":13750628739,"lastBlockReportTime":1640440500340,"lastUpdate":1640444445299,"lastUpdateMonotonic":13754573698,"name":"10.4.1.15:50010","networkLocation":"/default-rack","remaining":8141550999091,"storageID":"a3a145a9-88c6-4849-a2e2-216d8eef8143","xceiverCount":1,"xferPort":50010},{"adminState":"NORMAL","blockPoolUsed":7798472576268,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":7798472576268,"hostName":"pg-kvm-hdfs-dn1013","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.13","ipcPort":50020,"lastBlockReportMonotonic":13747610408,"lastBlockReportTime":1640437482009,"lastUpdate":1640444444963,"lastUpdateMonotonic":13754573363,"name":"10.4.1.13:50010","networkLocation":"/default-rack","remaining":8197323507063,"storageID":"fc4a02f7-4a83-44af-8438-a3afea9c691d","xceiverCount":2,"xferPort":50010}],"startOffset":268435456,"storageTypes":["DISK","DISK"]},"locatedBlocks":[{"block":{"blockId":1122169514,"blockPoolId":"BP-2028409268-10.4.1.16-1583828325109","generationStamp":2376930914,"numBytes":134217728},"blockToken":{"urlString":"AAAAAA"},"cachedLocations":[],"isCorrupt":false,"locations":[{"adminState":"NORMAL","blockPoolUsed":9015495742346,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":9015495742346,"hostName":"pg-kvm-hdfs-dn1018","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.18","ipcPort":50020,"lastBlockReportMonotonic":13749704685,"lastBlockReportTime":1640439576285,"lastUpdate":1640444445715,"lastUpdateMonotonic":13754574115,"name":"10.4.1.18:50010","networkLocation":"/default-rack","remaining":6980191326675,"storageID":"9a69bff0-dc92-4bf4-99dd-8a425165bb55","xceiverCount":1,"xferPort":50010},{"adminState":"NORMAL","blockPoolUsed":2830936782362,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":2830936782362,"hostName":"pg-kvm-hdfs-dn1005","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.5","ipcPort":50020,"lastBlockReportMonotonic":13739951773,"lastBlockReportTime":1640429823373,"lastUpdate":1640444443755,"lastUpdateMonotonic":13754572154,"name":"10.4.1.5:50010","networkLocation":"/default-rack","remaining":12368772591616,"storageID":"22ee4f6f-5c5d-4225-9972-d881401829e0","xceiverCount":1,"xferPort":50010}],"startOffset":0,"storageTypes":["DISK","DISK"]},{"block":{"blockId":1122185706,"blockPoolId":"BP-2028409268-10.4.1.16-1583828325109","generationStamp":2377095530,"numBytes":134217728},"blockToken":{"urlString":"AAAAAA"},"cachedLocations":[],"isCorrupt":false,"locations":[{"adminState":"NORMAL","blockPoolUsed":7798472576268,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":7798472576268,"hostName":"pg-kvm-hdfs-dn1013","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.13","ipcPort":50020,"lastBlockReportMonotonic":13747610408,"lastBlockReportTime":1640437482009,"lastUpdate":1640444444963,"lastUpdateMonotonic":13754573363,"name":"10.4.1.13:50010","networkLocation":"/default-rack","remaining":8197323507063,"storageID":"fc4a02f7-4a83-44af-8438-a3afea9c691d","xceiverCount":2,"xferPort":50010},{"adminState":"NORMAL","blockPoolUsed":9015495742346,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":9015495742346,"hostName":"pg-kvm-hdfs-dn1018","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.18","ipcPort":50020,"lastBlockReportMonotonic":13749704685,"lastBlockReportTime":1640439576285,"lastUpdate":1640444445715,"lastUpdateMonotonic":13754574115,"name":"10.4.1.18:50010","networkLocation":"/default-rack","remaining":6980191326675,"storageID":"9a69bff0-dc92-4bf4-99dd-8a425165bb55","xceiverCount":1,"xferPort":50010}],"startOffset":134217728,"storageTypes":["DISK","DISK"]},{"block":{"blockId":1122203888,"blockPoolId":"BP-2028409268-10.4.1.16-1583828325109","generationStamp":2377258060,"numBytes":125220942},"blockToken":{"urlString":"AAAAAA"},"cachedLocations":[],"isCorrupt":false,"locations":[{"adminState":"NORMAL","blockPoolUsed":7798472576268,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":7798472576268,"hostName":"pg-kvm-hdfs-dn1013","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.13","ipcPort":50020,"lastBlockReportMonotonic":13747610408,"lastBlockReportTime":1640437482009,"lastUpdate":1640444444963,"lastUpdateMonotonic":13754573363,"name":"10.4.1.13:50010","networkLocation":"/default-rack","remaining":8197323507063,"storageID":"fc4a02f7-4a83-44af-8438-a3afea9c691d","xceiverCount":2,"xferPort":50010},{"adminState":"NORMAL","blockPoolUsed":7844040235051,"cacheCapacity":0,"cacheUsed":0,"capacity":15997700407296,"dfsUsed":7844040235051,"hostName":"pg-kvm-hdfs-dn1015","infoPort":50075,"infoSecurePort":0,"ipAddr":"10.4.1.15","ipcPort":50020,"lastBlockReportMonotonic":13750628739,"lastBlockReportTime":1640440500340,"lastUpdate":1640444445299,"lastUpdateMonotonic":13754573698,"name":"10.4.1.15:50010","networkLocation":"/default-rack","remaining":8141550999091,"storageID":"a3a145a9-88c6-4849-a2e2-216d8eef8143","xceiverCount":1,"xferPort":50010}],"startOffset":268435456,"storageTypes":["DISK","DISK"]}]}}
---------

{'LocatedBlocks': {'fileLength': 393656398, 'isLastBlockComplete': True, 'isUnderConstruction': False, 'lastLocatedBlock': {'block': {'blockId': 1122203888, 'blockPoolId': 'BP-2028409268-10.4.1.16-1583828325109', 'generationStamp': 2377258060, 'numBytes': 125220942}, 'blockToken': {'urlString': 'AAAAAA'}, 'cachedLocations': [], 'isCorrupt': False, 'locations': [{'adminState': 'NORMAL', 'blockPoolUsed': 7844040235051, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 7844040235051, 'hostName': 'pg-kvm-hdfs-dn1015', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.15', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13750628739, 'lastBlockReportTime': 1640440500340, 'lastUpdate': 1640444445299, 'lastUpdateMonotonic': 13754573698, 'name': '10.4.1.15:50010', 'networkLocation': '/default-rack', 'remaining': 8141550999091, 'storageID': 'a3a145a9-88c6-4849-a2e2-216d8eef8143', 'xceiverCount': 1, 'xferPort': 50010}, {'adminState': 'NORMAL', 'blockPoolUsed': 7798472576268, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 7798472576268, 'hostName': 'pg-kvm-hdfs-dn1013', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.13', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13747610408, 'lastBlockReportTime': 1640437482009, 'lastUpdate': 1640444444963, 'lastUpdateMonotonic': 13754573363, 'name': '10.4.1.13:50010', 'networkLocation': '/default-rack', 'remaining': 8197323507063, 'storageID': 'fc4a02f7-4a83-44af-8438-a3afea9c691d', 'xceiverCount': 2, 'xferPort': 50010}], 'startOffset': 268435456, 'storageTypes': ['DISK', 'DISK']}, 'locatedBlocks': [{'block': {'blockId': 1122169514, 'blockPoolId': 'BP-2028409268-10.4.1.16-1583828325109', 'generationStamp': 2376930914, 'numBytes': 134217728}, 'blockToken': {'urlString': 'AAAAAA'}, 'cachedLocations': [], 'isCorrupt': False, 'locations': [{'adminState': 'NORMAL', 'blockPoolUsed': 9015495742346,'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 9015495742346, 'hostName': 'pg-kvm-hdfs-dn1018', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.18', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13749704685, 'lastBlockReportTime': 1640439576285, 'lastUpdate': 1640444445715, 'lastUpdateMonotonic': 13754574115, 'name': '10.4.1.18:50010', 'networkLocation': '/default-rack', 'remaining': 6980191326675, 'storageID': '9a69bff0-dc92-4bf4-99dd-8a425165bb55', 'xceiverCount': 1, 'xferPort': 50010}, {'adminState': 'NORMAL', 'blockPoolUsed': 2830936782362, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 2830936782362, 'hostName': 'pg-kvm-hdfs-dn1005', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.5', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13739951773, 'lastBlockReportTime': 1640429823373, 'lastUpdate': 1640444443755, 'lastUpdateMonotonic': 13754572154, 'name': '10.4.1.5:50010', 'networkLocation': '/default-rack', 'remaining': 12368772591616, 'storageID': '22ee4f6f-5c5d-4225-9972-d881401829e0', 'xceiverCount': 1, 'xferPort': 50010}], 'startOffset': 0, 'storageTypes': ['DISK', 'DISK']}, {'block': {'blockId': 1122185706, 'blockPoolId': 'BP-2028409268-10.4.1.16-1583828325109', 'generationStamp': 2377095530, 'numBytes': 134217728}, 'blockToken': {'urlString': 'AAAAAA'}, 'cachedLocations': [], 'isCorrupt': False, 'locations': [{'adminState': 'NORMAL','blockPoolUsed': 7798472576268, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 7798472576268, 'hostName': 'pg-kvm-hdfs-dn1013', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.13', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13747610408, 'lastBlockReportTime': 1640437482009, 'lastUpdate': 1640444444963, 'lastUpdateMonotonic': 13754573363, 'name': '10.4.1.13:50010', 'networkLocation': '/default-rack', 'remaining': 8197323507063, 'storageID': 'fc4a02f7-4a83-44af-8438-a3afea9c691d', 'xceiverCount': 2, 'xferPort': 50010}, {'adminState': 'NORMAL', 'blockPoolUsed': 9015495742346, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 9015495742346, 'hostName': 'pg-kvm-hdfs-dn1018', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.18', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13749704685, 'lastBlockReportTime': 1640439576285, 'lastUpdate': 1640444445715, 'lastUpdateMonotonic': 13754574115, 'name': '10.4.1.18:50010', 'networkLocation': '/default-rack', 'remaining': 6980191326675, 'storageID': '9a69bff0-dc92-4bf4-99dd-8a425165bb55', 'xceiverCount': 1, 'xferPort': 50010}], 'startOffset': 134217728, 'storageTypes': ['DISK', 'DISK']}, {'block': {'blockId': 1122203888, 'blockPoolId': 'BP-2028409268-10.4.1.16-1583828325109', 'generationStamp': 2377258060, 'numBytes': 125220942}, 'blockToken': {'urlString': 'AAAAAA'}, 'cachedLocations': [], 'isCorrupt': False, 'locations': [{'adminState': 'NORMAL', 'blockPoolUsed': 7798472576268, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 7798472576268, 'hostName': 'pg-kvm-hdfs-dn1013', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.13', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13747610408, 'lastBlockReportTime': 1640437482009, 'lastUpdate': 1640444444963, 'lastUpdateMonotonic': 13754573363, 'name': '10.4.1.13:50010', 'networkLocation': '/default-rack', 'remaining': 8197323507063, 'storageID': 'fc4a02f7-4a83-44af-8438-a3afea9c691d', 'xceiverCount': 2, 'xferPort': 50010}, {'adminState': 'NORMAL', 'blockPoolUsed': 7844040235051, 'cacheCapacity': 0, 'cacheUsed': 0, 'capacity': 15997700407296, 'dfsUsed': 7844040235051, 'hostName': 'pg-kvm-hdfs-dn1015', 'infoPort': 50075, 'infoSecurePort': 0, 'ipAddr': '10.4.1.15', 'ipcPort': 50020, 'lastBlockReportMonotonic': 13750628739, 'lastBlockReportTime': 1640440500340, 'lastUpdate':1640444445299, 'lastUpdateMonotonic': 13754573698, 'name': '10.4.1.15:50010', 'networkLocation': '/default-rack', 'remaining': 8141550999091, 'storageID': 'a3a145a9-88c6-4849-a2e2-216d8eef8143', 'xceiverCount': 1, 'xferPort': 50010}], 'startOffset': 268435456, 'storageTypes': ['DISK', 'DISK']}]}}
curl -s 'http://10.4.1.15:50075/webhdfs/v1/kafka-logs/neiwang-lb/access_log/20211223/access-15.log.gz?op=OPEN&namenoderpcaddress=mycluster&offset=0' -o neiwang-lb-20211223-15.log &
*/

func ParamFilter(params AccessLogDownloadParams) error {
	//参数校验
	if params.Hour != "" {
		hourInt, err := strconv.Atoi(params.Hour)
		if err != nil {
			return errors.New("hour 参数只能为0到23")
		}
		if hourInt < 0 || hourInt > 23 {
			return errors.New("hour 参数只能为0到23")
		}
	}

	if params.Type != 1 && params.Type != 2 {
		return errors.New("type 参数只能为1或2")
	}

	dateInt, err := strconv.Atoi(params.Date)
	if err != nil {
		return errors.New("date 参数格式错误")
	}

	start := time.Now().AddDate(0, 0, -30).Format("20060102")
	end := time.Now().Format("20060102")

	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)
	if dateInt < startInt || dateInt > endInt {
		return errors.New("date范围只能是最近30天")
	}
}
