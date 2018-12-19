package cispkg

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"log"
)

func PostDataToApi(apiurl,uuid,data string)  {

	pdata := make(url.Values)

	pdata["uuid"] = []string{uuid}
	pdata["data"] = [] string{data}


	res,err:=http.PostForm(apiurl,pdata)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("post send success")
	fmt.Println(string(body))


}

func GetUUID(cfgpath string) string  {

	CreateUUIDFile(cfgpath)

	//添加UUID，便于进行CIS检测
	file, err := os.Open(cfgpath)
	if err != nil {
		log.Println("err:", err)
	}
	defer file.Close()

	uuid, _ := ioutil.ReadAll(file)

	uuidstr := string(uuid)

	return uuidstr
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateUUIDFile(cfgpath string) {

	exists, _ := PathExists(cfgpath)

	if exists {

	} else {

		//生成uuid
		muuid, _ := uuid.NewV4()
		uuidfile, _ := os.OpenFile(cfgpath, os.O_RDWR|os.O_CREATE, 0766)
		uuidfile.WriteString(muuid.String())
		uuidfile.Close()
	}
}