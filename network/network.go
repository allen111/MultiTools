package network

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func Fetch() {
	for _,url:=range os.Args[1:]{
		resp,err:= http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}

		head:=resp.StatusCode
		if head<300 {
			body,_:=ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			fmt.Printf("status code: %d \n body: %s \n",head,body)
		}else{
			fmt.Printf("status code: %d \n",head)
		}



	}

}
