
package frontweb

import (
	"fmt"
	"net/http"
	_ "html/template"
	"log"

	"github.com/gin-gonic/gin"

	"google.golang.org/appengine"
    "google.golang.org/appengine/datastore"

	"github.com/vira0223/zippass/entity"

	_"golang.org/x/net/context"
	_"time"
	"strconv"
)

// initialize
func init() {

	http.Handle("/", GetMainEngine())

}

func GetMainEngine() *gin.Engine {
	log.Printf("debug: start %s", "GetMainEngine")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", RootGet)
	router.GET("/test/:id", TestGet)
	log.Printf("debug: end %s", "GetMainEngine")
	return router
}

func RootGet(g *gin.Context) {
	log.Printf("debug: start %s", "RootGet")
	g.HTML(http.StatusOK, "index.html", nil)
}

func TestGet(g *gin.Context) {
	log.Printf("debug: start %s", "TestGet")
	ids, _ := strconv.ParseInt(g.Param("id"), 10, 64)

	log.Printf("debug: ids-> %s", ids)

	// Datastoreからデータ取得。ids=KeyNameとして検索
	ctx := appengine.NewContext(g.Request)

	var passHeader entity.PassHeader
	//k := datastore.NewKey(ctx, "PassInfo", "", ids, nil)
	headKey := datastore.NewKey(ctx, "PassHeader", "", 5629499534213120, nil)
	log.Printf("debug: headkey-> %s", fmt.Sprint(headKey))
	// _ = datastore.Get(ctx, k, &passHeader)
	q := datastore.NewQuery("PassHeader").Ancestor(headKey)
	if _, err := q.GetAll(ctx, &passHeader); err != nil {
		log.Fatalf("とってもやばす：%v", err)
		return
	}

	g.HTML(http.StatusOK, "t.html", gin.H{
		"description": passHeader.PassDescription,
		/*
		"formatversion": passHeader.FormatVersion,
		"organizationname": passHeader.OrganizationName,
		"passtypeid": passHeader.PassTypeId,
		"serialnumber": passHeader.SerialNumber,
		"teamid": passHeader.TeamId,
		"applaunchurl": passHeader.AppLaunchURL,
		"associatedstoreid": passHeader.AssociatedStoreId,
		"userinfo": passHeader.Userinfo,
		"expirationdate": passHeader.ExpirationDate,
		"voided": passHeader.Voided,
		*/

//		"relevance": []entity.Beacon{ {passHeader.Beacons[0].Major, passHeader.Beacons[0].Minor, passHeader.Beacons[0].ProximityUUID, passHeader.Beacons[0].RelevantText},
//									  {passHeader.Beacons[1].Major, passHeader.Beacons[1].Minor, passHeader.Beacons[1].ProximityUUID, passHeader.Beacons[1].RelevantText} },
	})
}
